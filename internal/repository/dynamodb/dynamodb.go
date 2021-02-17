package dynamodb

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	repo "github.com/gymshark/software-onboarding/internal/repository"
)

var (
	ErrTable    = errors.New("table is required")
	ErrEndpoint = errors.New("endpoint is required")
	ErrRegion   = errors.New("region is required")
)

type DynamoConfig struct {
	Region   string
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

func (r repository) GetItems(table string, index string, items interface{}) error {
	if table == "" {
		return ErrTable
	}

	query := &dynamodb.QueryInput{
		TableName: &table,
		IndexName: &index,
	}

	result, err := r.client.Query(query)

	if err != nil {
		return fmt.Errorf("making query: %w", err)
	}

	if err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return fmt.Errorf("unmarshaling list of maps: %w", err)
	}


	return nil
}

func (r repository) SaveItem(table string, item interface{}) error {
	panic("")
}
