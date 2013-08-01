package main

import "fmt"

func isPalindrome(n int) bool {
	s := fmt.Sprintf("%d", n)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
