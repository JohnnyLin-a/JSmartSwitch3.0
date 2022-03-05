package lights

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildRouter(r *gin.RouterGroup) {
	r.POST("/on", on())
	r.POST("/off", off())
}

func on() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}

func off() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}