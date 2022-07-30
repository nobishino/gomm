package main

import "sync"

var x, y int

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { // Goroutine 1
		defer wg.Done()
		x = 1
		y = 1
	}()
	wg.Add(1)
	go func() { // Goroutine 2
		defer wg.Done()
		y := y
		x := x
		if x == 0 && y == 1 {
			panic("Message Passing Failed!")
		}
	}()
	wg.Wait()
}
