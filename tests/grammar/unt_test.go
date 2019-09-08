package grammar

import (
	"testing"
)

func TestUT(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log("ut")
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log("benchmark")
	}
}
