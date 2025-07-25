package main

import (
	"fmt"
	"image/color"
	"machine"
	"math/rand"
	"time"
	"tinygo.org/x/drivers/ws2812"
)

func main() {
	ledPin := machine.GPIO16
	ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws := ws2812.New(ledPin)

	fmt.Println("starting up\r\n")
	for {
		// We mask the random values with 0x3f to try to keep the brightness
		// to a reasonable intensity.
		leds := []color.RGBA{
			{
				R: uint8(rand.Int() & 0x3f),
				G: uint8(rand.Int() & 0x3f),
				B: uint8(rand.Int() & 0x3f),
			},
		}
		ws.WriteColors(leds)
		//fmt.Printf("hue: %v\r\nleds: %v\r\n", hue, leds)

		time.Sleep(2 * time.Second)
	}
}
