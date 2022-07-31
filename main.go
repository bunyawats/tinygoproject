package main

import (
	"machine"
	"math/rand"
	"sync"
	"time"
)

const (
	ledOnB = machine.LED
	led3   = machine.GP3
	led8   = machine.GP8
	led9   = machine.GP9
	led10  = machine.GP10
	bnt    = machine.GP20

	max = 700
	min = 300
)

var (
	state = true
	done  = make(chan bool, 1)
	wg    sync.WaitGroup
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

	led10.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	bnt.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})

	bnt.SetInterrupt(machine.PinFalling, isr)
}

func blink(led machine.Pin) {

	time.Sleep(time.Second * 2)

	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(max-min) + min

	for i := 0; i < 10; i++ {
		led.High()
		time.Sleep(time.Millisecond * time.Duration(delay))
		led.Low()
		time.Sleep(time.Millisecond * time.Duration(delay))
	}

}

func waitBlinkCh(led machine.Pin, d chan bool) {
	<-d
	blink(led)
	d <- true
}

func waitBlinkWg(led machine.Pin, w sync.WaitGroup) {

	blink(led)
	w.Done()
}

func blinkInRoutine() {

	go func() {
		blink(led8)
		blink(led9)
		blink(led10)
	}()
}

func waitBlingChan() {

	go waitBlinkCh(led8, done)
	go waitBlinkCh(led9, done)
	go waitBlinkCh(led10, done)
}

func waitBlinkGroup() {

	wg.Add(3)

	go waitBlinkWg(led8, wg)
	go waitBlinkWg(led9, wg)
	go waitBlinkWg(led10, wg)

	wg.Done()
}

func isr(p machine.Pin) {

	state = !state
	led3.Set(state)

	//blinkInRoutine()
	//waitBlingChan()
	waitBlinkGroup()

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
