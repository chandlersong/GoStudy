package concurrency

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	atomicCount int64
)

func printChar(wg *sync.WaitGroup, t *testing.T, character byte) {
	defer wg.Done()
	for count := 0; count < 3; count++ {
		for char := character; char < character+26; char++ {
			t.Logf("%c", char)
		}
	}
}

func incAtomicCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	for count := 0; count < 3; count++ {
		atomic.AddInt64(&atomicCount, 1)
		runtime.Gosched()
	}
}

func TestGoRoutine(t *testing.T) {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	t.Log("start run go")

	go printChar(&wg, t, 'a')
	go printChar(&wg, t, 'A')

	t.Log("waiting finished")
	wg.Wait()
	t.Log("finish")
}

func TestAtomicCounter(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	t.Log("start Atomic Counter")

	go incAtomicCounter(&wg)
	go incAtomicCounter(&wg)

	t.Log("waiting finished")
	wg.Wait()
	t.Logf("finish: and atomicCount is %d", atomicCount)
}
