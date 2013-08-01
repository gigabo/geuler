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
	solve := solutions[n]
	solution := "No solution yet!"
	if solve != nil {
		solution = solve()
	}
	return fmt.Sprintf("%v:\t%s", n, solution)
}

func main() {
	todo := os.Args[1:]
	if len(todo) == 0 {
		for k, _ := range solutions {
			todo = append(todo, k)
		}
	}
	for _, n := range todo {
		fmt.Println(run(n))
	}
}
