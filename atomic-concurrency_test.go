// from: https://gist.github.com/quasilyte/009edaf14aad08f6d1997b026c63c0a0

package main

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var counter int64

func addAtomic(x int64) {

	atomic.AddInt64(&counter, x)
}

func runBenchmark(b *testing.B, fn func()) {

	counter = 0
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {

		go func() {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					fn()
				}
			}
		}()
	}

	for int(counter) <= b.N {

		time.Sleep(100)
	}
	cancel()
	wg.Wait()
}

func BenchmarkAtomic(b *testing.B) {

	runBenchmark(b, func() {
		addAtomic(1)
	})
}
