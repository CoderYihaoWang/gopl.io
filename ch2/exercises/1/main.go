package main

import (
	"fmt"
)

func main() {
	k := Kalvin(0)
	f := Fahrenheit(0)
	c := Celsius(0)

	fmt.Printf("%v is %v\n", k, KToC(k))
	fmt.Printf("%v is %v\n", k, KToF(k))
	fmt.Printf("%v is %v\n", f, FToK(f))
	fmt.Printf("%v is %v\n", f, FToC(f))
	fmt.Printf("%v is %v\n", c, CToK(c))
	fmt.Printf("%v is %v\n", c, CToF(c))
}
