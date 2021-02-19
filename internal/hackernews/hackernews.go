package hackernews

import (
	"encoding/json"
	"fmt"
	repo "github.com/gymshark/software-onboarding/internal/repository"
	"github.com/gymshark/software-onboarding/internal/request"
	"github.com/sirupsen/logrus"
)

const (
	TopStoriesAPI = "https://hacker-news.firebaseio.com/v0/topstories.json"
	ItemAPI       = "https://hacker-news.firebaseio.com/v0/item/"
)

type Service interface {
	GetItems() ([]repo.Item, error)
}

type service struct {
	request request.Service
}

func NewHackersService() (Service, error) {
	h := &service{
		request: request.Service{},
	}

	return h, nil
}

func (h *service) GetItems() ([]repo.Item, error) {
	var itemIds []int
	var items []repo.Item

	data, err := h.request.Get(TopStoriesAPI)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &itemIds); err != nil {
		return nil, err
	}

	// Make it a goroutine
	for _, id := range itemIds {
		var item repo.Item

		url := fmt.Sprintf("%s%d.json", ItemAPI, id)

		rawData, reqErr := h.request.Get(url)

		if reqErr != nil {
			logrus.WithError(err).Error("error fetching item")
		} else {
			if unmarshalErr := item.UnmarshalJSON(rawData); unmarshalErr != nil {
				logrus.WithError(unmarshalErr).Error("error unmarshalling item")
			} else {
				items = append(items, item)
			}
		}
	}

	return items, nil
}