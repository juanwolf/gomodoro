package timer

import (
	"context"
	"time"
)

func Start(pomodoroDuration time.Duration, refreshRate time.Duration, ctx context.Context) (<-chan time.Duration, <-chan bool) {
	startingTime := time.Now()
	ticker := time.NewTicker(refreshRate)
	done := make(chan bool)
	tickChannel := make(chan time.Duration)
	go func() {
		select {
		case <-ctx.Done():
			ticker.Stop()
			done <- true
		}
	}()
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
