package web

import (
	"net/http"

	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/web/apiv1"
	"github.com/gin-gonic/gin"
)

func BuildRouter() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")

	// Healthcheck
	apiGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	apiV1Group := apiGroup.Group("/v1")
	apiv1.BuildRouter(apiV1Group)

	return r
}
