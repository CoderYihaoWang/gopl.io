// Rewrite dup2 to print the names of all files in which
// the duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fnames := make(map[string]map[string]bool)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fnames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fnames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			var sep string
			fmt.Print("In ")
			for name := range fnames[line] {
				fmt.Print(sep)
				fmt.Print(name)
				sep = ", "
			}
			fmt.Println(":")
			fmt.Printf("%d\t%s\n\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fnames map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if fnames[text] == nil {
			fnames[text] = make(map[string]bool)
		}
		fnames[text][f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
