package main

import (
	"fmt"
)

func main() {
	fmt.Println("Number please")
	empno := 0
	ename := ""
	fmt.Println("Current Value of empno ", empno)
	no, err := fmt.Scanf("%d", &empno)
	fmt.Println("Current Value of empno ", empno, " , "+ename)
	fmt.Println(no, "  ", err)
}
