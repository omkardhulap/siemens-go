package main

import (
	"fmt"
)

func print(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println("for ", str, " Current i = ", i)
	}
}

func main() {
	fmt.Println("in main.....")
	go print("str1")
	go print("str2")
	go print("str3")
	go print("str4")

	x := 0
	fmt.Println("Enter 1 number ")
	fmt.Scan(&x)
	fmt.Println("in end of  main.....")
}
