package lights

import (
	"log"
	"net/http"

	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/lights"
	"github.com/gin-gonic/gin"
)

func BuildRouter(r *gin.RouterGroup) {
	r.POST("/on", powerLights(true))
	r.POST("/off", powerLights(false))
	r.GET("/status", status())
}

func status() gin.HandlerFunc {
	return func(c *gin.Context) {
		count := lights.OnlineDevices()
		c.JSON(http.StatusOK, gin.H{
			"success":        true,
			"online_devices": count,
		})
	}
}

func powerLights(powerState bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := lights.PowerHS100(powerState)
		if err != nil {
			log.Println("Error when Power HS100", powerState, err)
		}
		err = lights.PowerMagicHome(powerState)
		if err != nil {
			log.Println("Error when Power MagicHome", powerState, err)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
