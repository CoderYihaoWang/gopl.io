package anagram

import "testing"

var data = [...]struct{
	s1, s2 string
	want bool
}{
	{"abcde", "eabcd", true},
	{"aabbcc", "aabcbc", true},
	{"", "", false},
	{"abc", "abcd", false},
	{"abc", "abcc", false},
	{"abc", "def", false},
	{"abcde", "abcde", false},
	{"aabbcc", "abcccc", false},
	{"世界", "界世", true},
	{"世界","世界杯", false},
	{"世界","世界", false},
}

func TestIsAnagram(t *testing.T) {
	for _, d := range data {
		if result := IsAnagram(d.s1, d.s2); result != d.want {
			t.Errorf("IsAnagram(%s, %s) = %t, want %t",
				d.s1, d.s2, result, d.want)
		}
	}
}