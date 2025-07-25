package machine

import "sync"

type PinCollection struct {
	pins []PinState
	lock sync.Mutex
}

func (pc *PinCollection) registerPin(p Pin) {
	if len(pc.pins) < int(p)+1 {
		newPins := make([]PinState, int(p)+1)
		copy(newPins, pc.pins)
		pc.pins = newPins
	}
}

func (pc *PinCollection) withLock(fn func()) {
	pc.Lock()
	defer pc.Unlock()
	fn()
}

func (pc *PinCollection) GetValue(p Pin) bool {
	var result bool
	pc.withLock(func() {
		pc.registerPin(p)
		result = pc.pins[p].value
	})

	return result
}

func (pc *PinCollection) SetValue(p Pin, val bool) {
	pc.withLock(func() {
		pc.registerPin(p)
		pc.pins[p].value = val
	})
}

func (pc *PinCollection) GetMode(p Pin) PinMode {
	var result PinMode
	pc.withLock(func() {
		pc.registerPin(p)
		result = pc.pins[p].mode
	})
	return result
}

func (pc *PinCollection) Configure(p Pin, config PinConfig) {
	pc.withLock(func() {
		pc.registerPin(p)
		pc.pins[p].mode = config.Mode
	})
}

func (pc *PinCollection) Lock() {
	pc.lock.Lock()
}

func (pc *PinCollection) Unlock() {
	pc.lock.Unlock()
}
