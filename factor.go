package main

func isPrime(n uint) bool {
	for i := uint(2); i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func largestPrimeFactor(n uint) uint {
	fact := uint(2)
	if isPrime(n) {
		return n
	}
	for i, max := fact, n/fact; i <= max; i++ {
		if n%i == 0 && isPrime(i) {
			fact = i
			max = n / fact
		}
	}
	return fact
}
