package main

import (
	"fmt"
	"sync"
)

var x, y int64

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		x = 1
		fmt.Println("y =", y)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		y = 1
		fmt.Println("x =", x)
		wg.Done()
	}()
	wg.Wait()
}
