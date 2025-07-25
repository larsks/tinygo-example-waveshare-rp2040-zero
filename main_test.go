package main

import (
	"testing"
)

func TestLEDController(t *testing.T) {
	// Test LED controller with mock
	controller := NewLEDController(GPIO16)
	controller.RunPattern(3)
}

func TestMachinePackage(t *testing.T) {
	// Test that we can use the machine package in tests
	pin := GPIO16
	pin.Configure(PinConfig{Mode: PinOutput})

	// These should work without panicking
	pin.High()
	pin.Low()
	pin.Set(true)
	pin.Set(false)
}
