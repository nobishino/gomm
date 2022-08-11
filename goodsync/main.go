package main

var a, b int = 0, 0
var ch = make(chan struct{})

func f() {
	a = 1 // write a
	ch <- struct{}{}
	b = 2 // write b
}

func g() {
	print(b) // read b
	<-ch
	print(a) // read a
}

func main() {
	go f()
	g()
}
