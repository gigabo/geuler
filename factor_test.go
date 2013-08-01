package main

import "testing"

func TestIsPrime(t *testing.T) {
	for _, test := range []struct {
		in  uint
		out bool
	}{
		{0, false}, {1, false}, {2, true}, {3, true}, {4, false},
		{5, true}, {6, false}, {7, true}, {8, false}, {9, false},
	} {
		f := isPrime(test.in)
		if f != test.out {
			t.Errorf("factor(%v) != %v [got: %v]", test.in, test.out, f)
		}
	}
}

func TestLargestPrimeFactor(t *testing.T) {
	for _, test := range []struct {
		in  uint
		out uint
	}{
		{2, 2}, {3, 3}, {4, 2}, {5, 5}, {6, 3},
		{7, 7}, {8, 2}, {9, 3}, {10, 5}, {11, 11},
		{13195, 29}, // This is the PE example.
	} {
		f := largestPrimeFactor(test.in)
		if f != test.out {
			t.Errorf("factor(%v) != %v [got: %v]", test.in, test.out, f)
		}
	}
}
