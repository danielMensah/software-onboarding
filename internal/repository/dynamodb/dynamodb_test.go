package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	repo "github.com/gymshark/software-onboarding/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
					Table: "table",
					Region:   "localhost",
					Endpoint: "endpoint",
				},
				client: nil,
			},
			wantErr: nil,
		},
		{
			name: "should return error for missing region",
			args: args{
				cfg: DynamoConfig{
					Table: "table",
					Region:   "",
					Endpoint: "endpoint",
				},
				client: nil,
			},
			wantErr: ErrRegion,
		},
		{
			name: "should return error for missing endpoint",
			args: args{
				cfg: DynamoConfig{
					Table: "table",
					Region:   "localhost",
					Endpoint: "",
				},
				client: nil,
			},
			wantErr: ErrEndpoint,
		},
		{
			name: "should return error for missing table",
			args: args{
				cfg: DynamoConfig{
					Table: "",
					Region:   "localhost",
					Endpoint: "endpoint",
				},
				client: nil,
			},
			wantErr: ErrTable,
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
		index string
		items interface{}
	}
	tests := []struct {
		name       string
		mockClient *mockDynamoDBClient
		args       args
		want       []repo.Item
		wantErr    bool
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
								"text": {
									S: aws.String("sample text"),
								},
								"type": {
									S: aws.String("story"),
								},
								"time": {
									S: aws.String("2021-02-13T00:00:00.00000Z"),
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
				index: "index",
				items: nil,
			},
			want: []repo.Item{
				{
					ID:    "blah-blah",
					Title: "test",
					Text:  "sample text",
					Type:  repo.ItemTypeStory,
					Time:  time.Date(2021, 02, 13, 0, 0, 0, 0, time.UTC),
					URL:   "https://test.com",
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
				index: "index",
				items: nil,
			},
			want:    []repo.Item{{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				client: tt.mockClient,
			}

			var items []repo.Item

			err := r.GetItems(tt.args.index, &items)

			assert.True(t, (err != nil) == tt.wantErr)
			assert.Equal(t, tt.want, items)
		})
	}
}
