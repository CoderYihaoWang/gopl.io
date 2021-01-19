// Prints all command line arguments to the console,
// including the program name. Separated by a space
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
