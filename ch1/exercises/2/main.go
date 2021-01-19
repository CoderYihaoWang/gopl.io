package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Index\tValue")
	fmt.Println("-----\t-----")
	for i, arg := range os.Args {
		fmt.Printf("%d\t%s\n", i, arg)
	}
}
