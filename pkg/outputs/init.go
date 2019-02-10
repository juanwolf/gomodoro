package outputs

import (
	"time"
)

// Output is a external representation of the current pomodoro.
type Output interface {
	GetName() string
	Refresh(timeLeft time.Duration)
	Start(pomodoroDuration time.Duration, refreshRate time.Duration)
	End()
}
