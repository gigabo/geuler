package main

var memo []uint

func fib(n uint) uint {
	if len(memo) <= int(n) {
		newMemo := make([]uint, n+1)
		copy(newMemo, memo)
		memo = newMemo
	}
	res := memo[n]
	if res != 0 {
	} else if n < 2 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}
	memo[n] = res
	return res
}
