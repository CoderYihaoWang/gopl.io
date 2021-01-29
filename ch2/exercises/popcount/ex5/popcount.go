package ex5

// PopCount counts the number of set bits in a uint64
func PopCount(x uint64) int {
	var count int
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}
