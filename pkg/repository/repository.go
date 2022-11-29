package repository

import (
	mps "portal/main-page-services"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user mps.User) (int, error)
	GetUser(username, password string) (mps.User, error)
}
type Quote interface {
	Create(quote mps.Quote) (int, error)
	GetAll() ([]mps.Quote, error)
	GetByRegionId(RegionId int) (mps.Quote, error)
	Delete(quoteId int) error
	Update(regionId int, input mps.UpdateQuoteInput) error
}

type TabsList interface {
	Create(list mps.PictureOfTheDayTabsList) (int, error)
	GetAll() ([]mps.PictureOfTheDayTabsList, error)
	GetByRegionId(regionId int) (mps.PictureOfTheDayTabsList, error)
	Delete(regionId int) error
}

type TabsItem interface {
	Create(regionId int, item mps.PictureOfTheDayTabsListsItem) (int, error)
	GetAllByRegion(regionId int) ([]mps.PictureOfTheDayTabsListsItem, error)
	Delete(tabId int) error
	Update(tabId int, input mps.UpdateTabItemInput) error
}

type Repository struct {
	Authorization
	Quote
	TabsList
	TabsItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Quote:         NewQuotePostgres(db),
		TabsList:      NewTabsListPostgres(db),
		TabsItem:      NewTabsItemPostgres(db),
	}
}
