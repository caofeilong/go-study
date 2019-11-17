package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	args := os.Args[1:]

	if len(args) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, " dup: %v \n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for text, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, text)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
}
