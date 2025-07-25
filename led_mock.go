//go:build !tinygo
// +build !tinygo

package main

import (
	"fmt"
	"image/color"
)

// MockLEDStrip implements LEDStrip for testing
type MockLEDStrip struct {
	pin Pin
}

func (m *MockLEDStrip) WriteColors(colors []color.RGBA) {
	fmt.Printf("[MOCK WS2812] Writing %d colors to pin %d: %v\n", len(colors), m.pin, colors)
}

// newLEDStrip creates the appropriate LED strip implementation
func newLEDStrip(pin Pin) LEDStrip {
	return &MockLEDStrip{pin: pin}
}
