package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	cur := 0
	for cur < len(s) {
		buf.WriteByte(s[cur])
		if (len(s)-cur)%3 == 1 && cur != len(s)-1 {
			buf.WriteByte(',')
		}
		cur++
	}
	return buf.String()
}