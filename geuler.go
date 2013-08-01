package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
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
	"2": func() string {
		n := uint(0)
		for i := uint(0); fib(i) < 4000000; i++ {
			if (fib(i) % 2) == 0 {
				n += fib(i)
			}
		}
		return fmt.Sprintf("%d", n)
	},
}

type result struct {
	id, solution string
	time         time.Duration
}

func run(n string) result {
	t0 := time.Now()
	solve := solutions[n]
	solution := "No solution yet!"
	if solve != nil {
		solution = solve()
	}
	return result{n, solution, time.Since(t0)}
}

func main() {
	t0 := time.Now()
	todo := os.Args[1:]
	if len(todo) == 0 {
		for k, _ := range solutions {
			todo = append(todo, k)
		}
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	for _, n := range todo {
		result := run(n)
		fmt.Fprintf(
			w, "%v:\t%s\t%v\n",
			result.id, result.solution, result.time,
		)
	}
	fmt.Fprintf(w, "\t\t\nTotal:\t%d problem(s)\t%s\n", len(todo), time.Since(t0))
	w.Flush()
}
