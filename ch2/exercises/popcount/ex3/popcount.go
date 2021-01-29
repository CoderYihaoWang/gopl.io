package ex3

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount calculates the number of set bits in a uint64
func PopCount(x uint64) int {
	var count byte
	for i := 0; i < 8; i++ {
		count += pc[byte(x)]
		x >>= 8
	}
	return int(count)
}
