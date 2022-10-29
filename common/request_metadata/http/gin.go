package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	requestMetadata "imperial-fleet-inventory/common/request_metadata"
)

func NewGINMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()

		md := requestMetadata.Metadata{
			RequestID: requestID,
		}

		c.Request = c.Request.WithContext(requestMetadata.AddToContext(c.Request.Context(), md))
		c.Next()
	}
}
