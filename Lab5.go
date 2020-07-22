package main

import (
	"fmt"
	"time"
)


func add(i int) time.Time {
	fmt.Println(i)
	x:= time.Now()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("in add", i ," time = " , x)
	return  x
}
func helper1(){
	for i :=1;i<5;i++{
		defer add(i)
		fmt.Println("cccc",i)
	}
	
}


func main() {

	helper1()
}

