package main

import (
	"github.com/gymshark/software-onboarding/internal/hackernews"
	"github.com/gymshark/software-onboarding/internal/repository"
)

type service struct {
	api hackernews.HackerNewService
	repo repository.Repository
}

func (s *service) syncItems() error {
	items, err := s.api.GetItems()

	if err != nil {
		return err
	}

	err = s.repo.SaveItems(items)

	if err != nil {
		return err
	}

	return nil
}
