package grammar

import (
	"fmt"
	"testing"
)

func TestUtExample(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Print("ut")
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Print("benchmark")
	}
}
