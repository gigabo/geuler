package main

import (
	"fmt"
	"os"
)

var solutions = map[string]func() string{
	"1": func() string {
		n := 0
		for i := 0; i < 1000; i++ {
			if (i%3) == 0 || (i%5) == 0 {
				n += i
			}
		}
		return fmt.Sprintf("%d", n)
	},
}

func run(n string) string {
	return fmt.Sprintf("%v:\t%s", n, solutions[n]())
}

func main() {
	fmt.Println(run(os.Args[1]))
}
