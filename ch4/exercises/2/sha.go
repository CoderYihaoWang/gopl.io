package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var algo = flag.Int("algo", 256, "SHA algorithm: one of  1 | 224 | 256 | 384 | 512")

func main() {
	flag.Parse()
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var output []byte
	switch *algo {
	case 1:
		sum := sha1.Sum(input)
		output = sum[:]
	case 224:
		sum := sha256.Sum224(input)
		output = sum[:]
	case 256:
		sum := sha256.Sum256(input)
		output = sum[:]
	case 384:
		sum := sha512.Sum384(input)
		output = sum[:]
	case 512:
		sum := sha512.Sum512(input)
		output = sum[:]
	default:
		log.Fatalf("sha: the algo parameter %d is invalid\n", *algo)
	}
	fmt.Printf("%x\n", output)
}