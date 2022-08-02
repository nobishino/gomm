package main

import (
	"sync"
	"sync/atomic"
)

var x, y int64

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { // Goroutine 1
		defer wg.Done()
		atomic.StoreInt64(&x, 1)
		atomic.StoreInt64(&y, 1)
	}()
	wg.Add(1)
	go func() { // Goroutine 2
		defer wg.Done()
		y := atomic.LoadInt64(&y)
		x := atomic.LoadInt64(&x)
		if x == 0 && y == 1 {
			panic("Message Passing Failed!")
		}
	}()
	wg.Wait()
}
