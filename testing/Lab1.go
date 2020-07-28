package main

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int {
	return i + j
}
func divide(s string, s1 string) int {
	n1, _ := strconv.Atoi(s)
	n2, _ := strconv.Atoi(s1)
	return n1 / n2
}
func main() {
	x := add(20, 30)
	fmt.Println("x = ", x)
	fmt.Println("div = ", divide("10", "5"))
}
