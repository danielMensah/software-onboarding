package item_service

import (
	"github.com/gin-gonic/gin"
	"github.com/gymshark/software-onboarding/internal/repository"
	"github.com/gymshark/software-onboarding/internal/repository/dynamodb"
	"net/http"
	"os"
)

func GetItems(c *gin.Context) {
	var items []repository.Item

	repo, repoErr := dynamodb.NewDynamo(dynamodb.DynamoConfig{
		Table:    os.Getenv("DYNAMO_TABLE"),
		Region:   os.Getenv("DYNAMO_REGION"),
		Endpoint: os.Getenv("DYNAMO_ENDPOINT"),
	}, nil)

	if repoErr != nil {
		_ = c.Error(repoErr)
	}

	itemType := c.Query("type")

	err := repo.GetItems("id", itemType, &items)

	if err != nil {
		_ = c.Error(err)
	} else {
		c.JSON(http.StatusOK, items)
	}
}
