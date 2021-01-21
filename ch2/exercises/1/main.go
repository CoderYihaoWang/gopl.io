package main

import (
	"fmt"

	t "tempconv/tempconv"
)

func main() {
	k := t.Kalvin(0)
	f := t.Fahrenheit(0)
	c := t.Celsius(0)

	fmt.Printf("%v is %v\n", k, t.KToC(k))
	fmt.Printf("%v is %v\n", k, t.KToF(k))
	fmt.Printf("%v is %v\n", f, t.FToK(f))
	fmt.Printf("%v is %v\n", f, t.FToC(f))
	fmt.Printf("%v is %v\n", c, t.CToK(c))
	fmt.Printf("%v is %v\n", c, t.CToF(c))
}
