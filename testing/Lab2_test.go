package main

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	t.Run("testdivide1", func(t *testing.T) {
		fmt.Println("testdivide1")
		x := divide("40", "20")
		if x != 2 {
			t.Error("Invalid Divide with 40,20")
		}
	})
	t.Run("testdivide2", func(t *testing.T) {
		fmt.Println("testdivide2")
		x := divide("40a", "20")
		if x != 0 {
			t.Error("Invalid Divide for 40a and 20")
		}
	})

}

func TestDivide3(t *testing.T) {
	x := divide("", "20")
	if x != 0 {
		t.Error("Invalid Divide '' and 20 ")
	}
}

func TestDivide4(t *testing.T) {
	defer func() {
		fmt.Println("end of main")
		r := recover()
		if r != nil {

			fmt.Println("Recover status ", r)
		} else {
			t.Error("Invalid Divide - divide by zero exception not raiased")
		}
	}()
	x := divide("40", "0")
	if x != 0 {
		t.Error("Invalid Divide divide by zero")
	}
}
