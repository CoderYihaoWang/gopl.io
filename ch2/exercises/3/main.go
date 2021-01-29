package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popcount(x uint64) int {
	var count byte
	for i := 0; i < 8; i++ {
		count += pc[byte(x)]
		x >>= 8
	}
	return int(count)
}

func main() {
	nums := [64]uint64{}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = 0xffffffffffffffff << i >> i
	}
	for _, n := range nums {
		fmt.Printf("%064b: %d\n", n, popcount(n))
	}
}
