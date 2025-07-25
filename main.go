package main

import (
	"fmt"
	"time"
)

type (
	Info struct {
		temp        int32
		chipVersion uint8
	}
)

func main() {
	controller := NewLEDController(GPIO16)

	fmt.Printf("starting up\r\n")
	for {
		info := getMachineInfo()
		fmt.Printf("chip_version: %v\r\n", info.chipVersion)
		fmt.Printf("temperature: %v\r\n", info.temp)
		controller.RunPattern(1, 2*time.Second)
	}
}
