package grammar

import (
	"sync"
	"testing"
)

/**
show how a once work.
once is a variable. you define one and it run one
*/
func TestOnce(t *testing.T) {

	var outerOnce sync.Once
	onceBody := func() {
		t.Log("outerOnce out loop")
	}

	for i := 0; i < 10; i++ {

		outerOnce.Do(onceBody)
		var innerOnce sync.Once
		innerOnce.Do(func() {
			t.Logf("inner Once,Run %d ", i)
		})
	}

}
