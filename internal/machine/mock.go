//go:build !tinygo
// +build !tinygo

// Package machine provides a mock implementation for testing
package machine

import "fmt"

// Pin represents a single GPIO pin
type Pin uint8

// PinMode sets the pin mode
type PinMode uint8

type PinState struct {
	mode  PinMode
	value bool
}

const (
	PinInput PinMode = iota
	PinOutput
)

// PinConfig holds the configuration for a pin
type PinConfig struct {
	Mode PinMode
}

var pins []PinState

func registerPin(p Pin) {
	if len(pins) < int(p)+1 {
		newPins := make([]PinState, int(p)+1)
		copy(newPins, pins)
		pins = newPins
	}
}

// Configure sets up the pin with the given configuration
func (p Pin) Configure(config PinConfig) {
	fmt.Printf("[MOCK] Configuring pin %d with mode %d\n", p, config.Mode)
	registerPin(p)
	pins[p].mode = config.Mode
}

// High sets the pin to high
func (p Pin) High() {
	fmt.Printf("[MOCK] Setting pin %d HIGH\n", p)
	registerPin(p)
	pins[p].value = true
}

// Low sets the pin to low
func (p Pin) Low() {
	fmt.Printf("[MOCK] Setting pin %d LOW\n", p)
	registerPin(p)
	pins[p].value = false
}

// Set sets the pin to the given value
func (p Pin) Set(value bool) {
	if value {
		p.High()
	} else {
		p.Low()
	}
}

func (p Pin) Get() bool {
	fmt.Printf("[MOCK] Get value of pin %d", p)
	registerPin(p)
	return pins[p].value
}

// GPIO pin definitions
const (
	GPIO16 Pin = 16
	// Add other pins as needed
)

func ReadTemperature() int32 {
	fmt.Printf("[MOCK] Read temperature")
	return 0
}

func ChipVersion() uint8 {
	fmt.Printf("[MOCK] Get chip version")
	return 0
}
