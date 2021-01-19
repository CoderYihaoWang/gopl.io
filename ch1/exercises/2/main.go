// Prints the indexes and values of all command line arguments
// including the program name. One line at a time
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
