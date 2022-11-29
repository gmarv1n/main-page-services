package service

import (
	mps "portal/main-page-services"
	"portal/main-page-services/pkg/repository"
)

type QuoteService struct {
	repo repository.Quote
}

func NewQuoteService(repo repository.Quote) *QuoteService {
	return &QuoteService{repo: repo}
}

func (s *QuoteService) Create(quote mps.Quote) (int, error) {
	return s.repo.Create(quote)
}

func (s *QuoteService) GetAll() ([]mps.Quote, error) {
	return s.repo.GetAll()
}

func (s *QuoteService) GetByRegionId(regionId int) (mps.Quote, error) {
	return s.repo.GetByRegionId(regionId)
}

func (s *QuoteService) Delete(quoteId int) error {
	return s.repo.Delete(quoteId)
}

func (s *QuoteService) Update(regionId int, input mps.UpdateQuoteInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(regionId, input)
}
