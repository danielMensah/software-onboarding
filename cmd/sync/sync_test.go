package main

import (
	repo "github.com/gymshark/software-onboarding/internal/repository"
	"github.com/gymshark/software-onboarding/internal/request"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockRepository struct {
	queryFunc error
}

func (m mockRepository) GetItems(index string, itemType string, items *[]repo.Item) error {
	return m.queryFunc
}

func (m mockRepository) SaveItems(items []repo.Item) error {
	return nil
}

type mockHackerNewService struct {
	getItemsFunc func() ([]repo.Item, error)
	request      request.Service
}

func (h mockHackerNewService) GetItems() ([]repo.Item, error) {
	return h.getItemsFunc()
}

func Test_service_syncItems(t *testing.T) {
	type fields struct {
		api  *mockHackerNewService
		repo *mockRepository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "can sync items",
			fields: fields{
				api: &mockHackerNewService{
					getItemsFunc: func() ([]repo.Item, error) {
						return []repo.Item{
							{
								ID:    "10",
								Title: "just title",
								Text:  "just text",
								Type:  "story",
								Time:  time.Now(),
								URL:   "https://hacker.com",
								By:    "dan",
							},
						}, nil
					},
				},
				repo: &mockRepository{
					queryFunc: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				api:  tt.fields.api,
				repo: tt.fields.repo,
			}

			err := s.syncItems()

			assert.True(t, (err != nil) == tt.wantErr)
		})
	}
}
