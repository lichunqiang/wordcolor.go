package wordcolor

import (
	"strconv"
	"testing"
)

func BenchmarkWordColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordColor(strconv.Itoa(i))
	}
}

func BenchmarkGetRGB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRGB(strconv.Itoa(i))
	}
}
