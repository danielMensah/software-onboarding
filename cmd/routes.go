package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gymshark/software-onboarding/cmd/job_service"
	"github.com/gymshark/software-onboarding/cmd/story_service"
	"net/http"
)

func initializeRoutes() {
	router.GET("/all", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello %s", "Daniel")
	})

	router.GET("/stories", story_service.GetStories)
	router.GET("/jobs", job_service.GetJobs)
}
