package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	repo "github.com/gymshark/software-onboarding/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDynamo(t *testing.T) {
	type args struct {
		cfg    DynamoConfig
		client dynamodbiface.DynamoDBAPI
	}
	tests := []struct {
		name    string
		args    args
		want    repo.Repository
		wantErr error
	}{
		{
			name: "should work as expected",
			args: args{
				cfg: DynamoConfig{
					Region:   "localhost",
					Endpoint: "http://127.0.0.1:8000",
				},
				client: nil,
			},
			wantErr: nil,
		},
		{
			name: "should return error for missing region",
			args: args{
				cfg: DynamoConfig{
					Region:   "",
					Endpoint: "http://127.0.0.1:8000",
				},
				client: nil,
			},
			wantErr: ErrRegion,
		},
		{
			name: "should return error for missing endpoint",
			args: args{
				cfg: DynamoConfig{
					Region:   "localhost",
					Endpoint: "",
				},
				client: nil,
			},
			wantErr: ErrEndpoint,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDynamo(tt.args.cfg, tt.args.client)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

type mockDynamoDBClient struct {
	getItemsFunc  func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	saveItemError error
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamoDBClient) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return m.getItemsFunc(input)
}

func (m *mockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, m.saveItemError
}

func Test_repository_GetItems(t *testing.T) {
	type args struct {
		table string
		index string
		items interface{}
	}
	tests := []struct {
		name    string
		mockClient  *mockDynamoDBClient
		args    args
		want    []repo.Story
		wantErr bool
	}{
		{
			name: "can get items",
			mockClient: &mockDynamoDBClient{
				getItemsFunc: func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
					return &dynamodb.QueryOutput{
						Items: []map[string]*dynamodb.AttributeValue{
							{
								"id": {
									S: aws.String("blah-blah"),
								},
								"title": {
									S: aws.String("test"),
								},
								"url": {
									S: aws.String("https://test.com"),
								},
								"by": {
									S: aws.String("john doe"),
								},
							},
						},
					}, nil
				},
			},
			args: args{
				table: "stories",
				index: "index",
				items: nil,
			},
			want: []repo.Story{
				{
					Id:    "blah-blah",
					Title: "test",
					Url:   "https://test.com",
					By:    "john doe",
				},
			},
			wantErr: false,
		},
		{
			name: "error unmarshalling",
			mockClient: &mockDynamoDBClient{
				getItemsFunc: func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
					return &dynamodb.QueryOutput{
						Items: []map[string]*dynamodb.AttributeValue{
							{
								"id": {
									BOOL: aws.Bool(false),
								},
							},
						},
					}, nil
				},
			},
			args: args{
				table: "stories",
				index: "index",
				items: nil,
			},
			want: []repo.Story{{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				client: tt.mockClient,
			}

			var items []repo.Story

			err := r.GetItems(tt.args.table, tt.args.index, &items)

			assert.True(t, (err != nil) == tt.wantErr)
			assert.Equal(t, tt.want, items)
		})
	}
}
