package lights

import (
	"log"
	"time"

	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/config"
	"github.com/johnnylin-a/go-hs100/pkg/configuration"
	"github.com/johnnylin-a/go-hs100/pkg/hs100"
	"github.com/johnnylin-a/go-magic-home/pkg/magichome"
)

func PowerMagicHome(powerState bool) error {
	devices, err := magichome.Discover(magichome.DiscoverOptions{
		BroadcastAddr: config.Get.MagicHome.BroadcastIP,
		Timeout:       1,
	})
	if err != nil {
		return err
	}
	for _, device := range *devices {
		controller, err := magichome.New(device.IP, 5577)
		if powerState && err == nil {
			controller.SetState(magichome.On)
		} else {
			controller.SetState(magichome.Off)
		}
	}
	return nil
}

func PowerHS100(powerState bool) error {
	devices, err := hs100.Discover(config.Get.HS100.BroadcastIP, configuration.Default().WithTimeout(time.Second))
	if err != nil {
		return err
	}
	log.Println(devices)
	for _, device := range devices {
		if powerState {
			device.TurnOn()
		} else {
			device.TurnOff()
		}
	}
	return nil
}
