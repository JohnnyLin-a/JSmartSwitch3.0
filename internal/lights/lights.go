package lights

import (
	"time"

	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/config"
	"github.com/johnnylin-a/go-hs100/pkg/configuration"
	"github.com/johnnylin-a/go-hs100/pkg/hs100"
	"github.com/johnnylin-a/go-magic-home/pkg/magichome"
)

var (
	magicHomeDevices *[]magichome.Device
	hs100Devices     []*hs100.Hs100
)

func init() {
	hs100Devices = make([]*hs100.Hs100, 0)
	go Scan()
}

func Scan() {
	ScanMagicHome()
	ScanHS100()
}

func ScanMagicHome() {
	devices, err := magichome.Discover(magichome.DiscoverOptions{
		BroadcastAddr: config.Get.MagicHome.BroadcastIP,
		Timeout:       1,
	})
	if err != nil {
		magicHomeDevices = devices
	}
}

func ScanHS100() {
	devices, err := hs100.Discover(config.Get.HS100.BroadcastIP, configuration.Default().WithTimeout(time.Second))
	if err != nil {
		hs100Devices = devices
	}
}

func PowerMagicHome(powerState bool) error {
	if magicHomeDevices == nil || len(*magicHomeDevices) == 0 {
		return nil
	}

	for _, device := range *magicHomeDevices {
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
	if len(hs100Devices) == 0 {
		return nil
	}

	for _, device := range hs100Devices {
		if powerState {
			device.TurnOn()
		} else {
			device.TurnOff()
		}
	}
	return nil
}
