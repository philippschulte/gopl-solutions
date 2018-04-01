// Exercise 1.4:
// Modify dup2 to print the names of all files
// in which each duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines("stdin", os.Stdin, counts, names)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dub2: %v\n", err)
				continue
			}

			countLines(arg, f, counts, names)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s:\t%d\t%s\n", names[line], n, line)
		}
	}
}

func countLines(name string, f *os.File, counts map[string]int, names map[string]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
		names[input.Text()] = name
	}

	// NOTE: ignoring potential errors from input.Err()
}
