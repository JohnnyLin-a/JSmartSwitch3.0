package lights

import (
	"log"
	"net"
	"time"

	"github.com/JohnnyLin-a/JSmartSwitch3.0/internal/config"
	"github.com/johnnylin-a/go-hs100/pkg/configuration"
	"github.com/johnnylin-a/go-hs100/pkg/hs100"
	"github.com/johnnylin-a/go-magic-home/pkg/magichome"
)

var (
	magicHomeDevices map[string]struct{}
	hs100Devices     map[string]struct{}
)

func init() {
	hs100Devices = make(map[string]struct{})
	magicHomeDevices = make(map[string]struct{})
	go func() {
		for {
			Scan()
			time.Sleep(time.Minute / 2)
		}
	}()
}

func Scan() {
	ScanMagicHome()
	ScanHS100()
	log.Println("Devices", OnlineDevices())
	for d := range hs100Devices {
		log.Println(d)
	}
	for d := range magicHomeDevices {
		log.Println(d)
	}
}

func ScanMagicHome() {
	log.Println("Scan MagicHome")
	devices, err := magichome.Discover(magichome.DiscoverOptions{
		BroadcastAddr: config.Get.MagicHome.BroadcastIP,
		Timeout:       1,
	})
	if err == nil {
		for _, device := range *devices {
			magicHomeDevices[device.IP.String()] = struct{}{}
		}
	}
}

func ScanHS100() {
	log.Println("Scan HS100")
	devices, err := hs100.Discover(config.Get.HS100.BroadcastIP, configuration.Default().WithTimeout(time.Second))
	if err == nil {
		for _, device := range devices {
			hs100Devices[device.Address] = struct{}{}
		}
	}
}

func PowerMagicHome(powerState bool) error {
	if len(magicHomeDevices) == 0 {
		return nil
	}

	for deviceIP := range magicHomeDevices {
		ip, _, _ := net.ParseCIDR(deviceIP + "/24")
		controller, err := magichome.New(ip, 5577)
		if err != nil {
			return err
		}
		if powerState {
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

	for deviceIP := range hs100Devices {
		device := hs100.NewHs100(deviceIP, configuration.Default().WithTimeout(time.Second))
		if powerState {
			device.TurnOn()
		} else {
			device.TurnOff()
		}
	}
	return nil
}

func OnlineDevices() int {
	count := len(hs100Devices) + len(magicHomeDevices)
	return count
}
