package main

import (
	"sync"
	"sync/atomic"
)

var x atomic.Int64
var y int64

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		print(x.Load())
		print(y)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		x.Store(1)
		y = 1
		wg.Done()
	}()
	wg.Wait()
}
