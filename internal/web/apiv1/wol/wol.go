package wol

import (
	"net/http"

	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/johnnylin-a/go-wol-lib/pkg/wol"
)

func BuildRouter(r *gin.RouterGroup) {
	r.POST("/wake", wake())
}

func wake() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := wol.Wake(config.Get.WOL.MacAddr, config.Get.WOL.BroadcastIP+":"+config.Get.WOL.Port)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
