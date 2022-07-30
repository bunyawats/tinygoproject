package main

import (
	"machine"
	"time"
)

const (
	ledOnB  = machine.LED
	led3 = machine.GP3
	led8 = machine.GP8
	bnt  = machine.GP20
)

var (
	state = true
)

func configuration() {

	ledOnB.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	led3.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	led8.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	bnt.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})

	bnt.SetInterrupt(machine.PinFalling, isr)
}

func isr(p machine.Pin) {

	state = !state
	led3.Set(state)

	go blink(led8)
}

func blink(led machine.Pin) {
	for i := 0; i < 5; i++ {
		led.High()
		time.Sleep(time.Millisecond * 500)
		led.Low()
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	configuration()

	count := 0

	for {
		if state {
			for count < 5 {
				ledOnB.High()
				time.Sleep(time.Millisecond * 500)
				ledOnB.Low()
				time.Sleep(time.Millisecond * 500)
				count++
			}
		} else {
			count = 0
		}
		ledOnB.Low()
	}
}
