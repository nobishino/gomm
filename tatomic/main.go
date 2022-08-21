package main

import "sync/atomic"

var x atomic.Int64
var y int64
var ch = make(chan struct{}) // 待ち合わせ

func main() {
	go func() {
		print(y)
		println(x.Load())
		ch <- struct{}{}
	}()
	x.Store(1)
	y = 1
	<-ch
}
