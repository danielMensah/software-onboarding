package dynamodb

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	repo "github.com/gymshark/software-onboarding/internal/repository"
)

var (
	ErrEndpoint  = errors.New("endpoint is required")
	ErrRegion  = errors.New("region is required")
)

type DynamoConfig struct {
	Table string
	Region string
	Endpoint string
}

type repository struct {
	client dynamodbiface.DynamoDBAPI
}

func NewDynamo(cfg DynamoConfig, client dynamodbiface.DynamoDBAPI) (repo.Repository, error) {
	d := &repository{}

	if cfg.Endpoint == "" {
		return nil, ErrEndpoint
	}

	if cfg.Region == "" {
		return nil, ErrRegion
	}

	if client != nil {
		d.client = client
	} else {
		d.client = dynamodb.New(session.Must(session.NewSession()), &aws.Config{
			Region:   aws.String(cfg.Region),
			Endpoint: aws.String(cfg.Endpoint),
			Credentials: credentials.NewChainCredentials([]credentials.Provider{
				&credentials.SharedCredentialsProvider{},
				&credentials.EnvProvider{},
			}),
		})
	}

	return d, nil
}

func (r repository) GetItems(table string, items interface{}) error {
	panic("")
}

func (r repository) SaveItem(table string, item interface{}) error {
	panic("")
}