package main

import (
	"sync"
)

var x, y int

func main() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { // Goroutine 1
		x = 1
		ch <- struct{}{}
		y = 1
		wg.Done()
	}()
	wg.Add(1)
	go func() { // Goroutine 2
		y := y
		<-ch
		x := x
		if x == 0 && y == 1 {
			panic("Message Passing Failed!")
		}
		wg.Done()
	}()
	wg.Wait()
}
