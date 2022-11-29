package repository

import (
	"fmt"

	mps "portal/main-page-services"

	"github.com/jmoiron/sqlx"
)

type TabsListPostgres struct {
	db *sqlx.DB
}

func NewTabsListPostgres(db *sqlx.DB) *TabsListPostgres {
	return &TabsListPostgres{db: db}
}

func (r *TabsListPostgres) Create(list mps.PictureOfTheDayTabsList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1) RETURNING id", tabsListsTable)
	row := tx.QueryRow(createListQuery, list.RegionId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TabsListPostgres) GetAll() ([]mps.PictureOfTheDayTabsList, error) {
	var lists []mps.PictureOfTheDayTabsList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", tabsListsTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *TabsListPostgres) GetByRegionId(regionId int) (mps.PictureOfTheDayTabsList, error) {
	var list mps.PictureOfTheDayTabsList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", tabsListsTable)
	err := r.db.Select(&list, query)

	return list, err
}

func (r *TabsListPostgres) Delete(regionId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", tabsListsTable)
	_, err := r.db.Exec(query, regionId)

	return err
}
