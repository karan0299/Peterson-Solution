package main

import (
	"fmt"
	"time"
)

var turn int = 0

var flag [2]bool

var buffer [5]int

var bufferSize int = 5

var counter int = 0

func consumer() {
	var nextConsumed int
	out := 0
	for true {
		for counter == 0 {
		}
		nextConsumed = buffer[out]
		fmt.Println("Comsumed ", nextConsumed)
		out = (out + 1) % bufferSize
		flag[0] = true
		turn = 1
		for flag[1] && turn == 1 {
		}
		counter--
		fmt.Println("counter:", counter)
		flag[0] = false
	}
}

func producer() {
	var nextProduced int = 0
	in := 0
	for true {
		for counter == bufferSize {
		}
		buffer[in] = nextProduced
		fmt.Println("Produced :", nextProduced)
		nextProduced++
		in = (in + 1) % bufferSize
		flag[1] = true
		turn = 0
		for flag[0] && turn == 0 {
		}
		counter++
		fmt.Println("counter:", counter)
		flag[1] = false
	}
}
func main() {

	go consumer()

	go producer()

	time.Sleep(5 * time.Millisecond)

}
