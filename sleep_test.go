package sleep

import "time"
import "testing"

func TestSleep(t *testing.T) {
	start := time.Now()
	sleep(10)
	end := time.Now()

	if end.Sub(start) < 10 {
		t.Error("Sleep exited before 10s")
	}
}
