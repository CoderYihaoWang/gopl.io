package main

import (
	"fmt"
	"popcount/popcount"
)

func main() {
	nums := [64]uint64{}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = 0xffffffffffffffff << i >> i
	}
	for _, n := range nums {
		fmt.Printf("%064b: %d\n", n, popcount.PopCount(n))
	}
}
