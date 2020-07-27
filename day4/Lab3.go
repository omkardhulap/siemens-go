package main

import (
	"fmt"
	"time"
)

var total = 0

func deposit() {

	for i := 1; i <= 500; i++ {
		x := total
		x++
		total = x
		time.Sleep(time.Millisecond * 1)
	}
	fmt.Println("Deposit Current Total =", total)
}
func widraw() {
	for i := 1; i <= 500; i++ {
		x := total
		x--
		time.Sleep(time.Millisecond * 1)
		total = x
	}
	fmt.Println("Widraw Current Total =", total)
}
func main() {
	go deposit()
	go widraw()
	for {
	}
}
