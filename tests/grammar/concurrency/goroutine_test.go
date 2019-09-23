package concurrency

import (
	"runtime"
	"sync"
	"testing"
)

func printUpperCase(wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	for count := 0; count < 3; count++ {
		for char := 'A'; char < 'A'+26; char++ {
			t.Logf("%c", char)
		}
	}
}

func printLowerCase(wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	for count := 0; count < 3; count++ {
		for char := 'a'; char < 'a'+26; char++ {
			t.Logf("%c", char)
		}
	}
}

func TestGoRoutine(t *testing.T) {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	t.Log("start run go")

	go printLowerCase(&wg, t)
	go printUpperCase(&wg, t)

	t.Log("waiting finished")
	wg.Wait()
	t.Log("finish")
}
