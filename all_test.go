package lockkey

import (
	"testing"
)

func BenchmarkUser(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		User(15)
	}
}

func BenchmarkAttackee(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Attackee(1, 1, 1)
	}
}
