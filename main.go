package main

import (
	"fmt"

	"github.com/amimof/huego"
)

func main() {
	bridge := huego.New("192.168.1.59", "username")
	l, err := bridge.GetLights()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found %d lights", len(l))
}
