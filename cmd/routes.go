package main

import (
	"github.com/gymshark/software-onboarding/cmd/item_service"
	"github.com/sirupsen/logrus"
)

func initializeRoutes() {
	s, err := item_service.New()

	if err != nil {
		logrus.WithError(err).Error("error initializing item service")
	}

	router.GET("/item", s.GetItems)
}