//go:build tinygo
// +build tinygo

package main

import (
	"image/color"
	"machine"
	"tinygo.org/x/drivers/ws2812"
)

// TinyGoLEDStrip implements LEDStrip for TinyGo
type TinyGoLEDStrip struct {
	ws ws2812.Device
}

func (t *TinyGoLEDStrip) WriteColors(colors []color.RGBA) {
	t.ws.WriteColors(colors)
}

// newLEDStrip creates the appropriate LED strip implementation
func newLEDStrip(pin Pin) LEDStrip {
	return &TinyGoLEDStrip{ws: ws2812.New(machine.Pin(pin))}
}
