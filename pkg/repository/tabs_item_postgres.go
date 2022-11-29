package repository

import (
	"fmt"
	"strings"

	mps "portal/main-page-services"

	"github.com/jmoiron/sqlx"
)

type TabstemPostgres struct {
	db *sqlx.DB
}

func NewTabsItemPostgres(db *sqlx.DB) *TabstemPostgres {
	return &TabstemPostgres{db: db}
}

func (r *TabstemPostgres) Create(regionId int, item mps.PictureOfTheDayTabsListsItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", tabsListsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.EntityId, item.Type, regionId)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TabstemPostgres) GetAllByRegion(regionId int) ([]mps.PictureOfTheDayTabsListsItem, error) {
	var items []mps.PictureOfTheDayTabsListsItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id 
						INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`, tabsListsTable)

	if err := r.db.Select(&items, query, regionId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TabstemPostgres) Delete(tabId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2", tabsListsTable)

	_, err := r.db.Exec(query, tabId)

	return err
}

func (r *TabstemPostgres) Update(tabId int, input mps.UpdateTabItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.EntityId != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.EntityId)
		argId++
	}

	if input.Type != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Type)
		argId++
	}

	if input.OrderPlace != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.OrderPlace)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul 
							WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		tabsListsTable, setQuery, argId, argId+1)
	args = append(args, tabId)

	_, err := r.db.Exec(query, args...)
	return err
}
