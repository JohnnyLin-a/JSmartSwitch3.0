package apiv1

import (
	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/web/apiv1/lights"
	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/web/apiv1/wol"
	"github.com/gin-gonic/gin"
)

func BuildRouter(r *gin.RouterGroup) {
	lightsGroup := r.Group("/lights")
	lights.BuildRouter(lightsGroup)
	wolGroup := r.Group("/wol")
	wol.BuildRouter(wolGroup)
}
