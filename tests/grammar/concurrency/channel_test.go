package concurrency

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func player(name string, court chan int, wg *sync.WaitGroup, t *testing.T) {

	defer wg.Done()

	for {
		ball, ok := <-court

		if !ok {
			t.Logf("Player %s Won \n", name)
			return
		}

		n := rand.Intn(100)

		if n%13 == 0 {
			t.Logf("Player %s Missed\n", name)
			close(court)
			return
		}

		t.Logf("Player %s Hit %d", name, ball)
		ball++
		court <- ball
	}
}

func TestChannelRelative(t *testing.T) {

	t.Run("TestPlayTennis", func(t *testing.T) {
		var wg sync.WaitGroup
		rand.Seed(time.Now().UnixNano())
		court := make(chan int)

		wg.Add(3)

		go player("chandler", court, &wg, t)
		go player("monica", court, &wg, t)
		go player("a", court, &wg, t)
		court <- 1

		wg.Wait()
	})
}
