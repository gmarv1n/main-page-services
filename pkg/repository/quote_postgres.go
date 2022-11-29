package repository

import (
	"fmt"
	"strings"

	mps "portal/main-page-services"

	"github.com/jmoiron/sqlx"
)

type QuotePostgres struct {
	db *sqlx.DB
}

func NewQuotePostgres(db *sqlx.DB) *QuotePostgres {
	return &QuotePostgres{db: db}
}

func (r *QuotePostgres) Create(quote mps.Quote) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var quoteId int
	createQuoteQuery := fmt.Sprintf("INSERT INTO %s (quote, author, author_description, record_id, region_id) values ($1, $2, $3, $4, $5) RETURNING id", quotesTable)
	row := tx.QueryRow(createQuoteQuery, quote.Quote, quote.Author, quote.AuthorDescription, quote.RecordId, quote.RegionId)
	err = row.Scan(&quoteId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return quoteId, tx.Commit()
}

func (r *QuotePostgres) GetAll() ([]mps.Quote, error) {
	var items []mps.Quote
	query := fmt.Sprintf(`SELECT q.id, q.quote, q.author, q.author_description, q.record_id, q.region_id 
						  FROM %s q`, quotesTable)

	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *QuotePostgres) GetByRegionId(regionId int) (mps.Quote, error) {
	var item mps.Quote
	query := fmt.Sprintf(`
		SELECT q.id, q.quote, q.author, q.author_description, q.record_id, q.region_id  
		FROM %s q
		WHERE q.region_id = $1`, quotesTable)

	if err := r.db.Get(&item, query, regionId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *QuotePostgres) Delete(QuoteId int) error {
	query := fmt.Sprintf("DELETE FROM %s q WHERE q.id = $1", quotesTable)

	_, err := r.db.Exec(query, QuoteId)

	return err
}

func (r *QuotePostgres) Update(regionId int, input mps.UpdateQuoteInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Quote != nil {
		setValues = append(setValues, fmt.Sprintf("quote=$%d", argId))
		args = append(args, *&input.Quote)
		argId++
	}

	if input.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
		args = append(args, *&input.Author)
		argId++
	}

	if input.AuthorDescription != nil {
		setValues = append(setValues, fmt.Sprintf("author_description=$%d", argId))
		args = append(args, *input.AuthorDescription)
		argId++
	}

	if input.RecordId != nil {
		setValues = append(setValues, fmt.Sprintf("record_id=$%d", argId))
		args = append(args, *&input.RecordId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s q SET %s
						  WHERE q.region_id = $%d`,
		quotesTable, setQuery, argId+1)

	args = append(args, regionId)

	_, err := r.db.Exec(query, args...)
	return err
}
