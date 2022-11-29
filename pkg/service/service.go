package service

import (
	mps "portal/main-page-services"
	"portal/main-page-services/pkg/repository"
)

type Authorization interface {
	CreateUser(user mps.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
	Quote
	TabsLists
	TabsItems
}

type Quote interface {
	Create(quote mps.Quote) (int, error)
	GetAll() ([]mps.Quote, error)
	GetByRegionId(RegionId int) (mps.Quote, error)
	Delete(QuoteId int) error
	Update(regionId int, input mps.UpdateQuoteInput) error
}

type Themes interface {
}

type TabsLists interface {
	Create(tabList mps.PictureOfTheDayTabsList) (int, error)
	GetAll() ([]mps.PictureOfTheDayTabsList, error)
	Delete(regionId int) error
}

type TabsItems interface {
	Create(regionId int, tab mps.PictureOfTheDayTabsListsItem) (int, error)
	GetAllByRegion(regionId int) ([]mps.PictureOfTheDayTabsListsItem, error)
	Delete(tabId int) error
	Update(tabId int, input mps.UpdateTabItemInput) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Quote:         NewQuoteService(repos.Quote),
		TabsLists:     NewTabsListService(repos.TabsList),
		TabsItems:     NewTabsItemsService(repos.TabsItem, repos.TabsList),
	}
}
