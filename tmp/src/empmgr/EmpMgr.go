package empmgr

import (
	"fmt"
)

type Emplist struct {
	EmpArr [10]Emp
	cp     int
}

func (emplist *Emplist) print(i int) {
	fmt.Println(emplist.EmpArr[i])
}

func (emplist *Emplist) Create(e Emp) {
	fmt.Println("in create ", e.Empno)
	emplist.EmpArr[emplist.cp] = e
	emplist.cp++
}
