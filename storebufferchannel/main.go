package main

import (
	"fmt"
	"sync"
)

var x, y int64

func main() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		x = 1
		ch <- struct{}{}
		fmt.Println("y =", y)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		y = 1
		<-ch
		fmt.Println("x =", x)
		wg.Done()
	}()
	wg.Wait()
}
