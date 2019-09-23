package concurrency

import (
	"runtime"
	"sync"
	"testing"
)

func printChar(wg *sync.WaitGroup, t *testing.T, character byte) {
	defer wg.Done()
	for count := 0; count < 3; count++ {
		for char := character; char < character+26; char++ {
			t.Logf("%c", char)
		}
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
