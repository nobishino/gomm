package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x, y int64

// go run -race で実行するとraceが検出されたりされなかったりします
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { // Goroutine 1
		defer wg.Done()
		x = 1
		atomic.StoreInt64(&y, 1)
	}()
	wg.Add(1)
	go func() { // Goroutine 2
		defer wg.Done()
		r2 := atomic.LoadInt64(&y)
		r1 := x
		fmt.Println(r2)
		if r1 == 0 && r2 == 1 {
			panic("Message Passing Failed!")
		}
	}()
	wg.Wait()
}
