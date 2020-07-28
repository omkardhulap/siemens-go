package main

import "testing"

func TestAdd1(t *testing.T) {
	x := add(10, 30)
	if x != 40 {
		t.Error("Invalid Sum")
	}
}
