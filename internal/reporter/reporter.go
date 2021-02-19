package reporter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   error
}

func ErrorReporter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(gin.ErrorTypeAny)

		fmt.Println("Handle APP error")

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err

			var parsedError *apiError

			parsedError = &apiError{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
				Error:   err,
			}

			c.AbortWithStatusJSON(parsedError.Code, parsedError)
			return
		}

	}
}
