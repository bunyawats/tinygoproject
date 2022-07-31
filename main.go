package main

import (
	"machine"
	"time"
)

const (
	ledOnB = machine.LED
	led3   = machine.GP3
	led8   = machine.GP8
	led9   = machine.GP9
	bnt    = machine.GP20
)

var (
	state = true
	done  = make(chan bool, 1)
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

	led9.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	bnt.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})

	bnt.SetInterrupt(machine.PinFalling, isr)
}

func waitBlink(led machine.Pin, d chan bool) {
	<-d
	blink(led)
	d <- true
}

func blink(led machine.Pin) {

	time.Sleep(time.Second * 2)

	for i := 0; i < 10; i++ {
		led.High()
		time.Sleep(time.Millisecond * 500)
		led.Low()
		time.Sleep(time.Millisecond * 500)
	}

}

func isr(p machine.Pin) {

	state = !state
	led3.Set(state)

	go waitBlink(led8, done)
	go waitBlink(led9, done)

}

func main() {

	configuration()
	led3.Set(state)
	done <- true

	for {
		if state {
			blink(ledOnB)
		}
		time.Sleep(time.Second * 2)
	}
}
