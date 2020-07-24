package main

import (
	"fmt"
	"strconv"
)

type tostr interface {
	Convert() string
}
type Emp struct {
	Empno int
	Ename string
}

func (e Emp) Convert() string {
	return "Empno = " + strconv.Itoa(e.Empno) + ", Ename = " + e.Ename
}
func main() {
	var a tostr
	e := Emp{10, "aaa"}
	a = e
	fmt.Println(a.Convert())

}
