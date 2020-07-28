package main

import "fmt"

func add(i int, j int) int {
	return i + j
}
func main() {
	x := add(20, 30)
	fmt.Println("x = ", x)
}
