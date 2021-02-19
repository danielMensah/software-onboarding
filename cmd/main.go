package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gymshark/software-onboarding/internal/reporter"
	log "github.com/sirupsen/logrus"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	// using middleware for custom errors
	router.Use(reporter.ErrorReporter())

	initializeRoutes()

	if err := router.Run(":8080"); err != nil {
		log.WithError(err).Error("running server")
	}
}
