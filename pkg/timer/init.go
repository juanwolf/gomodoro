package timer

import (
	"time"
)

func Start(pomodoroDuration time.Duration, refreshRate time.Duration) (<-chan time.Time, <-chan bool) {
	ticker := time.NewTicker(refreshRate)
	done := make(chan bool)
	go func() {
		time.Sleep(pomodoroDuration)
		ticker.Stop()
		done <- true
	}()
	return ticker.C, done
}
