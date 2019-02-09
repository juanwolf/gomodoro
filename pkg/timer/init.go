package timer

import (
	"time"
)

func Start(pomodoroDuration time.Duration, refreshRate time.Duration) (<-chan time.Duration, <-chan bool) {
	startingTime := time.Now()
	ticker := time.NewTicker(refreshRate)
	done := make(chan bool)
	tickChannel := make(chan time.Duration)
	go func() {
		time.Sleep(pomodoroDuration)
		ticker.Stop()
		done <- true
	}()
	go func() {
		for range ticker.C {
			tickChannel <- time.Since(startingTime)
		}
	}()

	return tickChannel, done
}
