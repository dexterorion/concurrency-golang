// https://www.reddit.com/r/golang/comments/apv6gj/why_is_a_channel_faster_than_a_mutex_in_this_test/

package main

import (
	"sync"
	"testing"
)

func BenchmarkStopByMutexReset(b *testing.B) {
	wg := new(sync.WaitGroup)
	wg.Add(b.N)
	m := make([]*sync.Mutex, b.N)
	for i := 0; i < b.N; i++ {
		m[i] = new(sync.Mutex)
		m[i].Lock()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func(m *sync.Mutex) {
			m.Lock()
			wg.Done()
		}(m[i])
	}
	for i := 0; i < b.N; i++ {
		m[i].Unlock()
	}
	wg.Wait()
}

func BenchmarkStopByChanReset(b *testing.B) {
	wg := new(sync.WaitGroup)
	wg.Add(b.N)
	done := make([]chan bool, b.N)
	for i := 0; i < b.N; i++ {
		done[i] = make(chan bool)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func(done chan bool) {
			<-done
			wg.Done()
		}(done[i])
	}
	for i := 0; i < b.N; i++ {
		done[i] <- true
	}
	wg.Wait()
}
