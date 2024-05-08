package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type config struct {
	Key string `env:"KEY"`
}

var Config config

func Init() {
	if err := env.Parse(&Config); err != nil {
		log.Fatalf("Unable to parse config: %v \n", err)
	}

}
