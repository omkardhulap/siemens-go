package main

import "testing"

func TestAdd(t *testing.T) {
	x := add(10, 30)

	if x != 40 {
		t.Error("Invalid Sum")
	}
}
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = add(i*10, i*30)
	}
}
