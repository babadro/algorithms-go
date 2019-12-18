package _461_hamming_distance

import "testing"

func BenchmarkHammingDistance(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = hammingDistance(34123414, 987993)
	}
	_ = res
}

func BenchmarkHammingDistance2(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = hammingDistance2(34123414, 987993)
	}
	_ = res
}

func BenchmarkHammingDistance3(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = hammingDistance3(34123414, 987993)
	}
	_ = res
}
