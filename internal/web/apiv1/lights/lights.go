package lights

import (
	"net/http"

	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/lights"
	"github.com/gin-gonic/gin"
)

func BuildRouter(r *gin.RouterGroup) {
	r.POST("/on", powerLights(true))
	r.POST("/off", powerLights(false))
}

func powerLights(powerState bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		go lights.PowerHS100(powerState)
		go lights.PowerMagicHome(powerState)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
