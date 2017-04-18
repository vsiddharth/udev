// +build !linux

package udev

import "errors"

type UDevMonitor struct{}

type UEvent struct{}

// NewMonitor creates and connects a new monitor
func NewMonitor() (mon *UDevMonitor, err error) {
	return &UDevMonitor{}, errors.New("Unsupported platform")
}

func (mon *UDevMonitor) Monitor(notify chan *UEvent) (shutdown chan bool) {
	return
}

func (mon *UDevMonitor) Close() error {
	return errors.New("Unsupported platform")
}
