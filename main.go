package main

import (
	"fmt"
	"time"
)

func main() {
	controller := NewLEDController(GPIO16)

	fmt.Printf("starting up\r\n")
	for {
		controller.RunPattern(1, 2*time.Second)
	}
}
