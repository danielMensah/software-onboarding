package main

import (
	"fmt"
	"github.com/gymshark/software-onboarding/internal/hackernews"
	"github.com/gymshark/software-onboarding/internal/repository"
)

type service struct {
	api hackernews.Service
	repo repository.Repository
}

func (s *service) syncItems() error {
	items, err := s.api.GetItems()

	if err != nil {
		return err
	}

	if len(items) < 1 {
		fmt.Println("nothing to sync")
		return nil
	}

	err = s.repo.SaveItems(items)

	if err != nil {
		return err
	}

	return nil
}
