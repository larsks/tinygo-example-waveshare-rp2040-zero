//go:build !tinygo
// +build !tinygo

// Package machine provides a mock implementation for testing
package machine

import "fmt"

// Pin represents a single GPIO pin
type Pin uint8

// PinMode sets the pin mode
type PinMode uint8

const (
	PinInput PinMode = iota
	PinOutput
)

// PinConfig holds the configuration for a pin
type PinConfig struct {
	Mode PinMode
}

// Configure sets up the pin with the given configuration
func (p Pin) Configure(config PinConfig) {
	fmt.Printf("[MOCK] Configuring pin %d with mode %d\n", p, config.Mode)
}

// High sets the pin to high
func (p Pin) High() {
	fmt.Printf("[MOCK] Setting pin %d HIGH\n", p)
}

// Low sets the pin to low
func (p Pin) Low() {
	fmt.Printf("[MOCK] Setting pin %d LOW\n", p)
}

// Set sets the pin to the given value
func (p Pin) Set(value bool) {
	if value {
		p.High()
	} else {
		p.Low()
	}
}

// GPIO pin definitions
const (
	GPIO16 Pin = 16
	// Add other pins as needed
)
