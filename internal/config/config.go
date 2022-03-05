package config

import "github.com/ilyakaznacheev/cleanenv"

type configDB struct {
	Process struct {
		Port string `env:"PORT" env-required:"true"`
	}
	WOL struct {
		Port        string `env:"WOL_PORT" env-required:"true"`
		BroadcastIP string `env:"WOL_BROADCAST_IP" env-required:"true"`
		MacAddr     string `env:"WOL_MAC_ADDR" env-required:"true"`
	}
	HS100 struct {
		BroadcastIP string `env:"HS100_BROADCAST_IP" env-required:"true"`
	}
	MagicHome struct {
		BroadcastIP string `env:"MAGICHOME_BROADCAST_IP" env-required:"true"`
	}
}

var Get configDB

func init() {
	err := cleanenv.ReadConfig(".env", &Get)
	if err != nil {
		err = cleanenv.ReadEnv(&Get)
		if err != nil {
			panic(err)
		}
	}
}
