package main

import (
	"github.com/gymshark/software-onboarding/internal/hackernews"
	"github.com/gymshark/software-onboarding/internal/repository/dynamodb"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	repo, err := dynamodb.NewDynamo(dynamodb.DynamoConfig{
		Table:    os.Getenv("DYNAMO_TABLE"),
		Region:   os.Getenv("DYNAMO_REGION"),
		Endpoint: os.Getenv("DYNAMO_ENDPOINT"),
	}, nil)

	if err != nil {
		logrus.WithError(err).Error("cannot initialise new dynamo")
		return
	}

	sync := service{
		api:  hackernews.HackerNewService{},
		repo: repo,
	}

	if syncErr := sync.syncItems(); syncErr != nil {
		logrus.WithError(err).Error("error while syncing items")
		return
	}
}
