package app

import (
	"github.com/gin-gonic/gin"

	httpRequestMetadata "imperial-fleet-inventory/common/request_metadata/http"
)

func (a *Application) initGinRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(httpRequestMetadata.NewGINMiddleware())

	return router
}
