package main

import (
	"strings"
	"testing"
)

var benchmarks = []struct {
	name string
	n    int64
}{
	{"4B", 4},
	{"16B", 16},
	{"100B", 100},
	{"4KB", 4e3},
	{"10MB", 10e6},
}

func BenchmarkSum64StringMurmur3(b *testing.B) {
	b.ReportAllocs()
	for _, bb := range benchmarks {
		s := strings.Repeat("a", int(bb.n))
		b.Run(bb.name, func(b *testing.B) {
			b.SetBytes(bb.n)
			for i := 0; i < b.N; i++ {
				_ = Sum64StringMurmur3(s)
			}
		})
	}
}

func BenchmarkSum64StringXxhash(b *testing.B) {
	b.ReportAllocs()
	for _, bb := range benchmarks {
		s := strings.Repeat("a", int(bb.n))
		b.Run(bb.name, func(b *testing.B) {
			b.SetBytes(bb.n)
			for i := 0; i < b.N; i++ {
				_ = Sum64StringXxhash(s)
			}
		})
	}
}

func BenchmarkSum64StringSHA1(b *testing.B) {
	b.ReportAllocs()
	for _, bb := range benchmarks {
		s := strings.Repeat("a", int(bb.n))
		b.Run(bb.name, func(b *testing.B) {
			b.SetBytes(bb.n)
			for i := 0; i < b.N; i++ {
				_ = Sum64StringSHA1(s)
			}
		})
	}
}

func BenchmarkSum64StringSHA256(b *testing.B) {
	b.ReportAllocs()
	for _, bb := range benchmarks {
		s := strings.Repeat("a", int(bb.n))
		b.Run(bb.name, func(b *testing.B) {
			b.SetBytes(bb.n)
			for i := 0; i < b.N; i++ {
				_ = Sum64StringSHA256(s)
			}
		})
	}
}

func BenchmarkSum64StringMD5(b *testing.B) {
	b.ReportAllocs()
	for _, bb := range benchmarks {
		s := strings.Repeat("a", int(bb.n))
		b.Run(bb.name, func(b *testing.B) {
			b.SetBytes(bb.n)
			for i := 0; i < b.N; i++ {
				_ = Sum64StringMD5(s)
			}
		})
	}
}

func BenchmarkSum64StringXXH3(b *testing.B) {
	b.ReportAllocs()
	for _, bb := range benchmarks {
		s := strings.Repeat("a", int(bb.n))
		b.Run(bb.name, func(b *testing.B) {
			b.SetBytes(bb.n)
			for i := 0; i < b.N; i++ {
				_ = Sum64StringXXH3(s)
			}
		})
	}
}

//func BenchmarkSum64String(b *testing.B) {
//	b.ReportAllocs()
//	for _, bb := range benchmarks {
//		s := strings.Repeat("a", int(bb.n))
//		b.Run(bb.name, func(b *testing.B) {
//			b.SetBytes(bb.n)
//			for i := 0; i < b.N; i++ {
//				_ = Sum64StringXXH3(s)
//			}
//		})
//	}
//}
