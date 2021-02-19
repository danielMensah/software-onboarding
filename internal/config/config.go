package config

import (
	"github.com/gymshark/software-onboarding/internal/repository/dynamodb"
	"os"
)

func New() dynamodb.DynamoConfig {
	return dynamodb.DynamoConfig{
		Table:    os.Getenv("DYNAMO_TABLE"),
		Region:   os.Getenv("DYNAMO_REGION"),
		Endpoint: os.Getenv("DYNAMO_ENDPOINT"),
	}
}
