package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x, y int64

// Litmus Test: ES2017 racy reads on ARMv8
// Can this program see r1 = 0, r2 = 1?
//
// Translation to Go program
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		atomic.StoreInt64(&x, 1)
		r1 := atomic.LoadInt64(&y)
		fmt.Println("r1 =", r1)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		atomic.StoreInt64(&y, 1)
		r2 := x
		fmt.Println("r2 =", r2)
		wg.Done()
	}()

	wg.Wait()
}
