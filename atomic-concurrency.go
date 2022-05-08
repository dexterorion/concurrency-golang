package main

import (
	"sync/atomic"
	"time"
)

var counter int32 = 0

func runSum() {
	for {
		atomic.AddInt32(&counter, 1)
	}
}

func runSub() {
	for {
		atomic.AddInt32(&counter, -1)
	}
}

func main() {
	go runSum()
	go runSub()

	time.Sleep(10 * time.Second)
}
