package main

import (
	"machine"
	"time"
)

const (
	led  = machine.LED
	led3 = machine.GP3
	led8 = machine.GP8
	bnt  = machine.GP20
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

	go func() {
		for i := 0; i < 5; i++ {
			led8.High()
			time.Sleep(time.Millisecond * 500)
			led8.Low()
			time.Sleep(time.Millisecond * 500)
		}
	}()
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
