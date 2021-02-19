package main

import (
	"github.com/gymshark/software-onboarding/internal/hackernews"
	"github.com/gymshark/software-onboarding/internal/repository/dynamodb"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	repo, repoErr := dynamodb.NewDynamo(dynamodb.DynamoConfig{
		Table:    os.Getenv("DYNAMO_TABLE"),
		Region:   os.Getenv("DYNAMO_REGION"),
		Endpoint: os.Getenv("DYNAMO_ENDPOINT"),
	}, nil)

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
