package main

import (
	"time"
)

var counter int32 = 0

func runSum() {
	for {
		counter = counter + 1
	}
}

func runSub() {
	for {
		counter = counter - 1
	}
}

func main() {
	go runSum()
	go runSub()
	time.Sleep(10 * time.Second)
}
