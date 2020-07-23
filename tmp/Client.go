package main

import (
	"empmgr"
	"fmt"
)

func main() {
	emp := empmgr.Emp{}
	fmt.Println("Enter Employee Details")
	fmt.Scan(&emp.Empno)
	fmt.Scan(&emp.Ename)
	fmt.Scan(&emp.Salary)
	empmgr.Create(emp)
	fmt.Println("Employee Created Successfully ", emp)

}
