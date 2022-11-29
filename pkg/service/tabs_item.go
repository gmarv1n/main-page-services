package service

import (
	mps "portal/main-page-services"
	"portal/main-page-services/pkg/repository"
)

type TabsItemsService struct {
	repo     repository.TabsItem
	listRepo repository.TabsList
}

func NewTabsItemsService(repo repository.TabsItem, listRepo repository.TabsList) *TabsItemsService {
	return &TabsItemsService{repo: repo, listRepo: listRepo}
}

func (s *TabsItemsService) Create(regionId int, item mps.PictureOfTheDayTabsListsItem) (int, error) {
	_, err := s.listRepo.GetByRegionId(regionId)
	if err != nil {
		// list does not exists or does not belongs to region
		return 0, err
	}

	return s.repo.Create(regionId, item)
}

func (s *TabsItemsService) GetAllByRegion(regionId int) ([]mps.PictureOfTheDayTabsListsItem, error) {

	return s.repo.GetAllByRegion(regionId)
}

func (s *TabsItemsService) Delete(tabId int) error {
	return s.repo.Delete(tabId)
}

func (s *TabsItemsService) Update(tabId int, input mps.UpdateTabItemInput) error {
	return s.repo.Update(tabId, input)
}
