package anagram

import "unicode/utf8"

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	letters := make(map[int32]int)
	for _, l := range s1 {
		letters[l]++
	}
	same := true
	for i, l := range s2 {
		if r, _ := utf8.DecodeRune([]byte(s1[i:])); r != l {
			same = false
		}
		letters[l]--
	}
	if same {
		return false
	}
	for _, v := range letters {
		if v != 0 {
			return false
		}
	}
	return true
}