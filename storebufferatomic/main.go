package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x, y int64

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		atomic.StoreInt64(&x, 1)
		y := atomic.LoadInt64(&y)
		fmt.Println("y =", y)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		atomic.StoreInt64(&y, 1)
		x := atomic.LoadInt64(&x)
		fmt.Println("x =", x)
		wg.Done()
	}()
	wg.Wait()
}
