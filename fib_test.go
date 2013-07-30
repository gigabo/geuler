package main

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		in  uint
		out uint
	}{{0, 1}, {1, 1}, {2, 2}, {3, 3}, {4, 5}, {5, 8}, {6, 13}}
	for _, test := range tests {
		f := fib(test.in)
		if f != test.out {
			t.Errorf("fib(%v) != %v [got: %v]", test.in, test.out, f)
		}
	}
}
