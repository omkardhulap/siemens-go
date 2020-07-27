package main

import (
	"fmt"
	"time"
)

func reader() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Reader - in count ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func writer(c chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Writer - in count ", i)
		time.Sleep(time.Millisecond * 100)

	}
	fmt.Println("after for in Writer...")
	c <- "Writer finished"

}
func main() {
	c := make(chan string)
	go writer(c)
	go reader()
	x := <-c
	fmt.Println("x = ", x)

	for {
	}
}
