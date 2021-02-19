package main

import (
	"github.com/gymshark/software-onboarding/cmd/item_service"
)

func initializeRoutes() {
	router.GET("/item", item_service.GetItems)
}