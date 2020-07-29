package main

import (
	"fmt"
	"strconv"
)

func divide(s string, s1 string) int {
	defer fmt.Println("end of divide")
	n1, err := strconv.Atoi(s)
	fmt.Println("err = ", err)
	n2, err1 := strconv.Atoi(s1)
	fmt.Println("err1 = ", err1)
	fmt.Println("n1  =", n1, "n2 = ", n2)
	ans := n1 / n2
	fmt.Println("ans = ", ans)
	return ans
}
func main() {
	defer func() {
		fmt.Println("end of main")
		r := recover()
		if r != nil {
			fmt.Println("Recover status ", r)
		}
	}()
	fmt.Println("div = ", divide("100", "0"))
}
