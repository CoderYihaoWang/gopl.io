package main

import (
	"strings"
	"testing"
)

var data = strings.Split("The quick brown fox jumps over the lazy dog", " ")

func concat(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func join(args []string) string {
	return strings.Join(args, " ")
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(data)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join(data)
	}
}