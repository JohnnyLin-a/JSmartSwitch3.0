package config

import "github.com/ilyakaznacheev/cleanenv"

type configDB struct {
	Process struct {
		Port string `env:"PORT" env-required:"true"`
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
