package main

import (
	"fmt"
)

func calc(x,y int) (int,int) {
	return x+y, x-y
}
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func main() {
	
	ans := add(10, 66)
	fmt.Println("Sum = ", ans)
	fmt.Println("Sub = ", sub(20,10))
	sum, sub:=calc(30,20)
	fmt.Println("Calc" , sum, " , " , sub)
}
