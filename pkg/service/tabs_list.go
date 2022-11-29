package service

import (
	mps "portal/main-page-services"
	"portal/main-page-services/pkg/repository"
)

type TabsListService struct {
	repo repository.TabsList
}

func NewTabsListService(repo repository.TabsList) *TabsListService {
	return &TabsListService{repo: repo}
}

func (s *TabsListService) Create(list mps.PictureOfTheDayTabsList) (int, error) {
	return s.repo.Create(list)
}

func (s *TabsListService) GetAll() ([]mps.PictureOfTheDayTabsList, error) {
	return s.repo.GetAll()
}

func (s *TabsListService) Delete(regionId int) error {
	return s.repo.Delete(regionId)
}
