package main

import (
	"flag"
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
		g := fibGenerator()
		n := uint(0)
		for f := <-g; f < 4000000; f = <-g {
			if (f % 2) == 0 {
				n += f
			}
		}
		return fmt.Sprintf("%d", n)
	},
	"3": func() string {
		n := uint(600851475143)
		return fmt.Sprintf("%d", largestPrimeFactor(n))
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
func runConcurrent(n string, c chan result) {
	c <- run(n)
}

func main() {
	concurrent := flag.Bool("concurrent", false, "Run concurrently")
	flag.Parse()
	t0 := time.Now()
	todo := flag.Args()
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	var channels []chan result
	if len(todo) == 0 {
		for k, _ := range solutions {
			todo = append(todo, k)
		}
	}
	if *concurrent {
		for _, n := range todo {
			c := make(chan result)
			channels = append(channels, c)
			go runConcurrent(n, c)

		}
	}
	for i, n := range todo {
		var res result
		if *concurrent {
			res = <-channels[i]
		} else {
			res = run(n)
		}
		fmt.Fprintf(
			w, "%v:\t%s\t%v\n",
			res.id, res.solution, res.time,
		)
	}
	fmt.Fprintf(w, "\t\t\nTotal:\t%d problem(s)\t%s\n", len(todo), time.Since(t0))
	w.Flush()
}
