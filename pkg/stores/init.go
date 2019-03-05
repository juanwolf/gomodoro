package stores

import (
	"time"
)

type Pomodoro struct {
	Id         int
	Message    string
	Stopped    bool
	Created_at time.Time
}

func NewPomodoro(message string) Pomodoro {
	return Pomodoro{
		Message:    message,
		Stopped:    false,
		Created_at: time.Now(),
	}
}

type Store interface {
	// Init will initialize the Store
	Init() error

	// AddPomodoro will add a pomodoro to the store or throw an error
	AddPomodoro(p Pomodoro) error

	// GetPomodoros will return all the pomodoros from the store
	GetPomodoros() ([]Pomodoro, error)

	// GetPomodorosGroupByDay will return a list of pomodoros from the store group by creation date
	GetPomodorosPerDay() (map[string][]Pomodoro, error)

	// GetLastPomodoro will return the last pomodoro entry from the store
	GetLastPomodoro() (Pomodoro, error)

	// GetNextBreakDuration will calculate the duration of the next break for the day.
	// Aka: pomodoroOfTheDay % 4 == 0 ? 15min : 5min
	GetNextBreakDuration() (time.Duration, error)
}
