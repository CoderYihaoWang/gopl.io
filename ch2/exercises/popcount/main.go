package main

import (
	"fmt"

	p3 "popcount/ex3"
)

func main() {
	nums := [64]uint64{}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = 0xffffffffffffffff << i >> i
	}

	fmt.Println("Ex 2.3")
	for _, n := range nums {
		fmt.Printf("%064b: %d\n", n, p3.PopCount(n))
	}
}
