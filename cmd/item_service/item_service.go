package item_service

import (
	"github.com/gin-gonic/gin"
	"github.com/gymshark/software-onboarding/internal/config"
	"github.com/gymshark/software-onboarding/internal/repository"
	"net/http"
)

type itemService struct {
	repo repository.Repository
}

func New() (*itemService, error) {
	cfg := config.New()
	repo, repoErr := repository.New(cfg)

	if repoErr != nil {
		return nil, repoErr
	}

	i := &itemService{repo: repo}

	return i, nil
}

func (s *itemService) GetItems(c *gin.Context) {
	var items []repository.Item

	itemType := c.Query("type")

	err := s.repo.GetItems("id", itemType, &items)

	if err != nil {
		_ = c.Error(err)
	} else {
		c.JSON(http.StatusOK, items)
	}
}
