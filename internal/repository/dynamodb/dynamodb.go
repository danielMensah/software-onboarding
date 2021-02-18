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
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	repo "github.com/gymshark/software-onboarding/internal/repository"
)

var (
	ErrTable    = errors.New("table is required")
	ErrEndpoint = errors.New("endpoint is required")
	ErrRegion   = errors.New("region is required")
)

type DynamoConfig struct {
	Table    string
	Region   string
	Endpoint string
}

type repository struct {
	table  string
	client dynamodbiface.DynamoDBAPI
}

func NewDynamo(cfg DynamoConfig, client dynamodbiface.DynamoDBAPI) (repo.Repository, error) {

	if cfg.Table == "" {
		return nil, ErrTable
	}

	if cfg.Endpoint == "" {
		return nil, ErrEndpoint
	}

	if cfg.Region == "" {
		return nil, ErrRegion
	}

	d := &repository{
		table: cfg.Table,
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

func (r repository) GetItems(index string, itemType string, items *[]repo.Item) error {
	query := &dynamodb.QueryInput{
		TableName: &r.table,
		IndexName: &index,
	}

	if itemType != "" {
		cond := expression.Key("type").Equal(expression.Value(itemType))

		expr, err := expression.NewBuilder().WithKeyCondition(cond).Build()

		if err != nil {
			return fmt.Errorf("building expression: %w", err)
		}

		query.ExpressionAttributeNames = expr.Names()
		query.ExpressionAttributeValues = expr.Values()
		query.KeyConditionExpression = expr.KeyCondition()
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

func (r repository) SaveItems(items []repo.Item) error {
	panic("")
}
