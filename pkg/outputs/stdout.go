package outputs

import (
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"time"
)

type Stdout struct {
	ShowPercent bool
	Size        int
	bar         *pb.ProgressBar
}

func NewStdout(showPercent bool) *Stdout {
	stdout := Stdout{
		ShowPercent: showPercent,
	}
	return &stdout
}

func (s *Stdout) Start(pomodoroDuration time.Duration, refreshRate time.Duration) {
	fmt.Println("Starting pomodoro for", pomodoroDuration)
	// create bar
	s.bar = pb.New((int(pomodoroDuration.Seconds())))
	// refresh info every second (default 200ms)
	s.bar.SetRefreshRate(refreshRate)
	// show percents (by default already true)
	s.bar.ShowPercent = s.ShowPercent
	// show bar (by default already true)
	s.bar.ShowBar = true
	// no counters
	s.bar.ShowCounters = false
	// show "time left"
	s.bar.ShowTimeLeft = true
	// show average speed
	s.bar.ShowSpeed = false
	// sets the width of the progress bar
	s.bar.SetWidth(s.Size)
	// sets the width of the progress bar, but if terminal size smaller will be ignored
	s.bar.SetMaxWidth(s.Size)
	s.bar.Start()
}

func (s *Stdout) Refresh() {
	s.bar.Increment()
}

func (s *Stdout) End() {
	s.bar.Finish()
}
