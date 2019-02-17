package outputs

import (
	"time"
)

type OutputManager struct {
	Outputs []*Output
}

func NewOutputManager() OutputManager {
	return OutputManager{}
}

func (o *OutputManager) Add(output *Output) {
	o.Outputs = append(o.Outputs, output)
}

func (o *OutputManager) Refresh(timeSpent time.Duration) {
	for _, output := range o.Outputs {
		(*output).Refresh(timeSpent)
	}
}

func (o *OutputManager) Start(pomodoroDuration time.Duration, refreshRate time.Duration) {
	for _, output := range o.Outputs {
		(*output).Start(pomodoroDuration, refreshRate)
	}
}

func (o *OutputManager) End() {
	for _, output := range o.Outputs {
		(*output).End()
	}
}
