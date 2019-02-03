package outputs

import (
	"time"
)

type Output interface {
	Refresh()
	Start(pomodoroDuration time.Duration, refreshRate time.Duration)
	End()
}
