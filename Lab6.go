package main

import (
	"fmt"
)

func shift(v Vertex) Vertex {
	v.X += 10
	v.Y += 10
	return v
}

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{}
	fmt.Println("Current Value of v ", v)
	fmt.Println("Please enter 1 co-ord with x and y values")
	no, err := fmt.Scanf("%d%d", &v.X, &v.Y)
	fmt.Println("Current Value of v ", v)
	fmt.Println(no, "  ", err)
	v1 := shift(v)
	fmt.Println(v1)
}
