package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	return r
}
