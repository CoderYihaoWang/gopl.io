package main

import (
	"fmt"
	"os"
	"strconv"

	"lenconv/lenconv"
)

func main() {
	for _, n := range os.Args[1:] {
		l, err := strconv.ParseFloat(n, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "lenconv: %v\n", err)
			os.Exit(1)
		}
		i := lenconv.Inch(l)
		cm := lenconv.CM(l)
		fmt.Printf("%s = %s, %s = %s\n", i, lenconv.IToCM(i), cm, lenconv.CMToI(cm))
	}
}
