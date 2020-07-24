package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	fmt.Println("Slice = ", a, "Len = ", len(a), ", cap  = ", cap(a))

	b := make([]string, 2, 5)
	fmt.Println("Slice = ", b, "Len = ", len(b), ", cap  = ", cap(b))
	b[0] = "aa"
	b[1] = "bb"
	b = append(b, "cc")
	b = append(b, "dd")
	b = append(b, "ee")
	b = append(b, "6th")
	b = append(b, "7th")
	b = append(b, "8th")
	b = append(b, "9th")
	b = append(b, "10th")

	fmt.Println("Slice = ", b, "Len = ", len(b), ", cap  = ", cap(b))
	b = append(b, "11th")
	fmt.Println("Slice = ", b, "Len = ", len(b), ", cap  = ", cap(b))
}
