package main

import (
	"fmt"
)

func main() {
	x:= add(20,40)
	fmt.Println(x)
}

func add(x, y int) int {
	return x+y
}
