package main

import (
	"fmt"
)
func helper1(){
	for i :=1;i<5;i++{
		defer fmt.Println(" Current Value = " ,i)
	}
}


func main() {
	helper1()
}
