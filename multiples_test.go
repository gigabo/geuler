package main

import "testing"

func TestSmallestMultiple(t *testing.T) {
	for _, test := range []struct {
		in, out uint
	}{
		{2, 2}, {3, 6}, {4, 12}, {5, 60},
		{10, 2520}, // This is the example for problem 5.
	} {
		m := smallestMultiple(test.in)
		if m != test.out {
			t.Errorf("fib(%v) != %v [got: %v]", test.in, test.out, m)
		}
	}
}
