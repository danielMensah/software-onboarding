package main

import (
	"github.com/gymshark/software-onboarding/internal/config"
	"github.com/gymshark/software-onboarding/internal/hackernews"
	"github.com/gymshark/software-onboarding/internal/repository"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New()
	repo, repoErr := repository.New(cfg)

	if repoErr != nil {
		logrus.WithError(repoErr).Error("cannot initialise new dynamo")
		return
	}

	hackerService, hackerServiceErr := hackernews.NewHackersService()

	if hackerServiceErr != nil {
		logrus.WithError(hackerServiceErr).Error("cannot initialise new hacker news service")
		return
	}

	sync := service{api: hackerService, repo: repo}

	if syncErr := sync.syncItems(); syncErr != nil {
		logrus.WithError(syncErr).Error("error while syncing items")
		return
	}
}
