package main

import (
	"fmt"
	"sync"
	"time"
)

var total = 0
var mux = &sync.Mutex{}

func deposit() {

	for i := 1; i <= 500; i++ {
		mux.Lock()
		x := total
		x++
		time.Sleep(time.Millisecond * 2)
		total = x
		mux.Unlock()
	}
	fmt.Println("Deposit Current Total =", total)
}
func widraw() {
	for i := 1; i <= 500; i++ {
		mux.Lock()
		x := total
		x--
		time.Sleep(time.Millisecond * 1)
		total = x
		mux.Unlock()
	}
	fmt.Println("Widraw Current Total =", total)
}
func main() {
	go deposit()
	go widraw()
	for {
	}
}
