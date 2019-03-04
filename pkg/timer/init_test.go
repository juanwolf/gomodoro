package timer

import (
	"context"
	"testing"
	"time"
)

var (
	pomodoroDuration time.Duration = 2 * time.Second
	refreshRate      time.Duration = 1 * time.Second
)

func TestStartShouldCancelTimerIfContextCanceled(t *testing.T) {
	expected := true
	ctx, cancel := context.WithCancel(context.Background())
	start := time.Now()
	_, done := Start(pomodoroDuration, refreshRate, ctx)

	cancel()

	found := <-done
	execDuration := time.Since(start)
	t.Log(execDuration)

	if found != expected {
		t.Errorf("Expected %v, found %v", expected, found)
	}

	if execDuration >= pomodoroDuration {
		t.Errorf("Start was not cancelled from context but from time exhaustion. Execution time expected < %v.\n Execution time found: %v", pomodoroDuration, execDuration)
	}
}

func TestStartShouldLastPomodoroDuration(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	start := time.Now()
	_, done := Start(pomodoroDuration, refreshRate, ctx)
	<-done

	execDuration := time.Duration.Round(time.Since(start), time.Second)

	if execDuration < pomodoroDuration {
		t.Errorf("Start channel was not cancelled because of time exhaustion... Execution time expected == %v.\n Execution time found: %v", pomodoroDuration, execDuration)
	}
}

func TestStartShouldReturnChannelUpdatedEveryRefreshRate(t *testing.T) {
	refreshEventsNumberExpected := int(pomodoroDuration.Seconds() / refreshRate.Seconds())
	refreshEventsNumberFound := 0
	ctx, _ := context.WithCancel(context.Background())
	ticker, done := Start(pomodoroDuration, refreshRate, ctx)

	for {
		select {
		case timeElapsed := <-ticker:
			refreshEventsNumberFound++
			secondsElapsed := time.Duration.Round(timeElapsed, time.Second).Seconds()
			secondsExpected := float64(refreshEventsNumberFound) * refreshRate.Seconds()
			if secondsElapsed != secondsExpected {
				t.Errorf("Time Elapsed expected: %v got %v", secondsExpected, secondsElapsed)
			}
		case <-done:
			if refreshEventsNumberFound != refreshEventsNumberExpected {
				t.Errorf("Expected %d refresh events, got %d", refreshEventsNumberExpected, refreshEventsNumberFound)
			}
			return
		}
	}

}
