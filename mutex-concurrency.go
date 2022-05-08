package main

import (
	"sync"
	"time"
)

var counter int32 = 0

var mu = sync.Mutex{}

func runSum() {
	for {
		mu.Lock()
		counter = counter + 1
		mu.Unlock()
	}
}

func runSub() {
	for {
		mu.Lock()
		counter = counter - 1
		mu.Unlock()
	}
}

func main() {
	go runSum()
	go runSub()

	time.Sleep(10 * time.Second)
}
