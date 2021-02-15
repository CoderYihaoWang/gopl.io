package main

import (
	"crypto/sha256"
	"fmt"
)

var pc = func()[256]int {
	var pc [256]int
	for i := range pc {
		pc[i] = pc[i/2] + i&1
	}
	return pc
}()

var data = []struct{
	a, b string
} {
	{"x", "X"},
	{"a", "a"},
	{"hello", "world"},
}

func main() {
	for _, d := range data {
		shaA, shaB := sha256.Sum256([]byte(d.a)), sha256.Sum256([]byte(d.b))
		fmt.Printf("%s: %x\n", d.a, shaA)
		fmt.Printf("%s: %x\n", d.b, shaB)
		fmt.Printf("Different bit(s): %d\n", shaDiff(shaA, shaB))
	}
}

func shaDiff(a, b [sha256.Size]byte) int {
	var diff int
	for i := range a {
		diff += pc[a[i]^b[i]]
	}
	return diff
}