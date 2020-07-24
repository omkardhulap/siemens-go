package main

import (
	"fmt"
)

func main() {
	var marks map[string]int
	fmt.Println(marks)
	marks = make(map[string]int)
	fmt.Println(marks)
	marks["s1"] = 100
	marks["s2"] = 90
	marks["s3"] = 70
	fmt.Println("Map = ", marks, "Len = ", len(marks))
	x, avail := marks["s4"]
	fmt.Println("x = ", x, " avail= ", avail)
	x, avail = marks["s2"]
	fmt.Println("x = ", x, " avail = ", avail)
}
