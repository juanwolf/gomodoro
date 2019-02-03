package outputs

import (
	"time"
)

type OutputManager struct {
	outputs []*Output
}

func NewOutputManager() OutputManager {

	return OutputManager{}
}

func (o *OutputManager) Add(output *Output) {
	o.outputs = append(o.outputs, output)
}

func (o *OutputManager) Refresh() {
	for _, output := range o.outputs {
		(*output).Refresh()
	}
}

func (o *OutputManager) Start(pomodoroDuration time.Duration, refreshRate time.Duration) {
	for _, output := range o.outputs {
		(*output).Start(pomodoroDuration, refreshRate)
	}
}

func (o *OutputManager) End() {
	for _, output := range o.outputs {
		(*output).End()
	}
}
