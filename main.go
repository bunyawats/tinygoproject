package main

import (
	"machine"
	"time"
)

const (
	led  = machine.LED
	led3 = machine.GP3
	pin  = machine.GP20
)

var (
	state = true
)

func configuration() {

	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	led3.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	pin.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})

	pin.SetInterrupt(machine.PinFalling, isr)
}

func isr(p machine.Pin) {
	state = !state
	led3.Set(state)
}

func main() {

	configuration()

	count := 0

	for {
		if state {
			for count < 5 {
				led.High()
				time.Sleep(time.Millisecond * 500)
				led.Low()
				time.Sleep(time.Millisecond * 500)
				count++
			}
		} else {
			count = 0
		}
		led.Low()
	}
}
