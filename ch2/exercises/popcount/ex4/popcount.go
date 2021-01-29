package ex4

// PopCount count the number of set bits in a uint64
func PopCount(x uint64) int {
	var count uint64
	for i := 0; i < 64; i++ {
		count += x & 1
		x >>= 1
	}
	return int(count)
}
