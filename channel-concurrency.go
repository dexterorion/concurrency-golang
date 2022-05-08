package main

import (
	"time"
)

var counter int32 = 0

var sync chan bool = make(chan bool, 1)

func runSum() {
	for {
		<-sync
		counter = counter + 1
		sync <- true
	}
}

func runSub() {
	for {
		<-sync
		counter = counter - 1
		sync <- true
	}
}

func main() {
	go runSum()
	go runSub()

	sync <- true

	time.Sleep(10 * time.Second)
}
