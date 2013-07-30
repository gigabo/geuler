package main

func fib(n uint) uint {
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
