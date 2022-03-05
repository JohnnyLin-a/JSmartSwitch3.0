package main

import (
	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/config"
	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/web"
)

func main() {
	r := web.BuildRouter()

	r.Run(":" + config.Get.Process.Port)
}
