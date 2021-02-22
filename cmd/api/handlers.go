package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gymshark/software-onboarding/internal/config"
	"github.com/gymshark/software-onboarding/internal/repository"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetItemsHandler(c *gin.Context) {
	cfg := config.New()
	repo, err := repository.New(cfg)

	if err != nil {
		log.WithError(err).Error("error connecting to repo")
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var items []repository.Item

	queryErr := repo.GetItems("id", c.Query("type"), &items)

	if queryErr != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else {
		c.JSON(http.StatusOK, items)
	}
}

func GetItemByIdHandler(c *gin.Context) {

}