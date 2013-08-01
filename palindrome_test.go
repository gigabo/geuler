package main

import "testing"

func TestPalindrome(t *testing.T) {
	for _, test := range []struct {
		in  int
		out bool
	}{
		{0, true}, {1, true},
		{10, false}, {11, true},
		{101, true}, {111, true}, {110, false},
		{1001, true}, {1111, true}, {1101, false},
		{10001, true}, {11111, true}, {11011, true}, {11101, false},
		{100001, true}, {111111, true}, {110011, true},
		{101101, true}, {111101, false},
	} {
		f := isPalindrome(test.in)
		if f != test.out {
			t.Errorf("isPalindrome(%v) != %v [got: %v]", test.in, test.out, f)
		}
	}
}
