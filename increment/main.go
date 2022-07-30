package main

import "time"

var count int

func increment() {
	if count == 0 {
		count = count + 1
	}
}

func main() {
	go increment()
	go increment()
	time.Sleep(time.Second)
	print(count)
}
