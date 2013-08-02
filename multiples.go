package main

func smallestMultiple(n uint) uint {
OUTER:
	for i := n; ; i += n {
		for j := n - 1; j > 1; j-- {
			if i%j != 0 {
				continue OUTER
			}
		}
		return i

	}

}
