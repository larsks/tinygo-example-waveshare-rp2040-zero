//go:build !tinygo
// +build !tinygo

package main

import mock "example/internal/machine"

// Re-export mock machine types for consistent access
type Pin = mock.Pin
type PinMode = mock.PinMode
type PinConfig = mock.PinConfig

const (
	PinInput  = mock.PinInput
	PinOutput = mock.PinOutput
	GPIO16    = mock.GPIO16
)
