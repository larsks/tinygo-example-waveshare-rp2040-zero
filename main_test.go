package main

import (
	"testing"
	"time"
)

func TestLEDController(t *testing.T) {
	// Test LED controller with mock
	controller := NewLEDController(GPIO16)
	controller.RunPattern(3, 100*time.Millisecond)
}

func TestMachinePackage(t *testing.T) {
	// Test that we can use the machine package in tests
	pin := GPIO16
	pin.Configure(PinConfig{Mode: PinOutput})

	// Test pin state tracking
	pin.High()
	if !pin.Get() {
		t.Errorf("Expected pin to be HIGH after calling High(), got LOW")
	}

	pin.Low()
	if pin.Get() {
		t.Errorf("Expected pin to be LOW after calling Low(), got HIGH")
	}

	pin.Set(true)
	if !pin.Get() {
		t.Errorf("Expected pin to be HIGH after calling Set(true), got LOW")
	}

	pin.Set(false)
	if pin.Get() {
		t.Errorf("Expected pin to be LOW after calling Set(false), got HIGH")
	}
}
