package main

import (
	"fmt"
	"log"

	"github.com/amimof/huego"
	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigHue struct {
	Host  string `yaml:"host"  env:"HOST"`
	Token string `yaml:"token" env:"TOKEN"`
}

var cfg ConfigHue

func main() {
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	bridge := huego.New(cfg.Host, cfg.Token)
	light, err := bridge.GetLight(1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	// Check the current state of the light
	state := light.State
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	// If the light is on, turn it off; if it's off, turn it on
	if state.On {
		_ = light.Off()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		fmt.Println("The light was on, it's now off")
	} else {
		_ = light.On()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		fmt.Println("The light was off, it's now on")
	}

}
