package grammar

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	t.Run("epoch time", func(t *testing.T) {
		now := time.Now()

		nanoSeconds := now.UnixNano()
		t.Logf("now %v", now)
		t.Logf("nano seconds is %v", nanoSeconds)
		millSeconds := nanoSeconds / (1000 * 1000)
		t.Logf("mill seconds is %v", millSeconds)
		seconds := millSeconds / 1000
		nanoPart := (millSeconds % 1000) * 1000
		t.Logf("seconds is %v", seconds)
		t.Logf("naos is %v", nanoPart)

		t.Logf("back to time:%v", time.Unix(seconds, 0))
	})
}
