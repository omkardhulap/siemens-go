package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	primes := [10]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var s []int = primes[:]
	fmt.Println("s with :", s)
	s = primes[3:7]
	fmt.Println("s with 3:7", s)
	primes[3] = 99
	s[1] = 88
	fmt.Println("s with 3:7", s)
	fmt.Println("Prime:   ", primes)
	fmt.Println("slice len = ", len(s), ", cap = ", cap(s))
	fmt.Println("Array len = ", len(primes), ", cap = ", cap(primes))

}
