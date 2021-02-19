package reporter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	error
}

func ErrorReporter() gin.HandlerFunc {
	return errorReporterT(gin.ErrorTypeAny)
}

func errorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		fmt.Println("Handle APP error")

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err

			var parsedError *apiError

			switch err.(type) {
			case *apiError:
				parsedError = err.(*apiError)
			default:
				parsedError = &apiError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}

			c.AbortWithStatusJSON(parsedError.Code, parsedError)
			return
		}

	}
}
