package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	dataset := []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 3},
		{1, 4, 5},
		{2, 4, 6},
		{-2, 4, 2},
		{0, 0, 0},
	}

	for _, val := range dataset {
		re := add(val.a, val.b)
		if re != val.out {
			t.Errorf("add(%d, %d) = %d, want %d", val.a, val.b, re, val.out)
		}
	}
}

func BenchmarkAdd(bb *testing.B) {
	var a, b, c int
	a = 123
	b = 456
	c = 579
	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Println(a, b, c, actual)
		}
	}
}

const numbers = 10000

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = str + strconv.Itoa(j)
		}
	}
	b.StopTimer()
}

func BenchmarkBuildStr(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var build strings.Builder
		for j := 0; j < numbers; j++ {
			build.WriteString(strconv.Itoa(j))
		}
		_ = build.String()
	}
	b.StopTimer()
}
