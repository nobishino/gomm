package main

// https://go.dev/ref/mem#badsync

var a, b int = 0, 0

func f() {
	a = 1 // write a
	b = 2 // write b
}

func g() {
	print(b) // read b
	print(a) // read a
}

func main() {
	go f()
	g()
}
