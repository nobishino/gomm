package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	go store()
	fmt.Println(x)
}

var x int64

func store() {
	atomic.StoreInt64(&x, 1)
}
