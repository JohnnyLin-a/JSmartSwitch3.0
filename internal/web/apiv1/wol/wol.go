package wol

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildRouter(r *gin.RouterGroup) {
	r.POST("/wake", wake())
}

func wake() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
