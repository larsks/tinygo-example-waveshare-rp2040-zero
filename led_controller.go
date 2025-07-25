package main

import (
	"image/color"
	"math/rand"
	"time"
)

// LEDController handles LED operations
type LEDController struct {
	pin Pin
	ws  LEDStrip
}

// LEDStrip interface allows mocking
type LEDStrip interface {
	WriteColors(colors []color.RGBA)
}

// NewLEDController creates a new LED controller
func NewLEDController(pin Pin) *LEDController {
	pin.Configure(PinConfig{Mode: PinOutput})

	return &LEDController{
		pin: pin,
		ws:  newLEDStrip(pin),
	}
}

// RunPattern runs the LED pattern
func (c *LEDController) RunPattern(iterations int, duration time.Duration) {
	for range iterations {
		colors := []color.RGBA{
			{
				R: uint8(rand.Int() & 0x3f),
				G: uint8(rand.Int() & 0x3f),
				B: uint8(rand.Int() & 0x3f),
			},
		}
		c.ws.WriteColors(colors)
		time.Sleep(duration)
	}
}
