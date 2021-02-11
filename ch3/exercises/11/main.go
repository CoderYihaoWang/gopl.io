package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// None recursive version of comma using bytes.Buffer
// This version also allows floating point number
// and an optional sign
func comma(s string) string {
	var buf bytes.Buffer
	n := strings.Index(s, ".")
	if n < 0 {
		n = len(s)
	}
	cur := 0
	if n > 0 && (s[0] == '+' || s[0] == '-') {
		buf.WriteByte(s[0])
		cur++
	}
	for cur < n {
		buf.WriteByte(s[cur])
		if (n-cur)%3 == 1 && cur != n-1 {
			buf.WriteByte(',')
		}
		cur++
	}
	for cur < len(s) {
		buf.WriteByte(s[cur])
		cur++
	}
	return buf.String()
}