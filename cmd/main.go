package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	initializeRoutes()

 	if err := router.Run(":8080"); err != nil {
		log.WithError(err).Error("running server")
	}
}