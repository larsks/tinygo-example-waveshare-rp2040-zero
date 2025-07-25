//go:build tinygo
// +build tinygo

package main

import "machine"

// Re-export machine types for consistent access
type Pin = machine.Pin
type PinMode = machine.PinMode
type PinConfig = machine.PinConfig

const (
	PinInput  = machine.PinInput
	PinOutput = machine.PinOutput
	GPIO16    = machine.GPIO16
)

func getMachineInfo() *Info {
	return &Info{
		temp:        machine.ReadTemperature(),
		chipVersion: machine.ChipVersion(),
	}
}
