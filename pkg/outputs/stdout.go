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

func (s *Stdout) GetName() string {
	return "stdout"
}

func (s *Stdout) Start(pomodoroDuration time.Duration, refreshRate time.Duration, message string) {
	fmt.Println("Good luck on", message)
	// create bar
	s.bar = pb.New((int(pomodoroDuration.Seconds() / refreshRate.Seconds())))
	s.bar.SetRefreshRate(refreshRate)
	s.bar.ShowPercent = s.ShowPercent
	s.bar.ShowBar = true
	s.bar.ShowCounters = false
	s.bar.ShowTimeLeft = true
	s.bar.ShowSpeed = false
	s.bar.SetWidth(s.Size)
	s.bar.SetMaxWidth(s.Size)
	s.bar.Start()
}

func (s *Stdout) Refresh(_timeLeft time.Duration) {
	s.bar.Increment()
}

func (s *Stdout) End() {
	s.bar.Finish()
}
