package main

import (
	"fmt"
	"time"
)

var x int

func main() {
	go func() {
		x = 2
	}()
	go func() {
		x = 4
	}()

	time.Sleep(time.Second)
	fmt.Println(x % 2) // 0にならないことがありうるか？
}
