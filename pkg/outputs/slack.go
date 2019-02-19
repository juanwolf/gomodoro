package outputs

import (
	"fmt"
	"github.com/juanwolf/slack"
	"time"
)

type Slack struct {
	DoNotDisturb bool
	Emoji        string
	client       *slack.Client
}

func NewSlack(token string, dnd bool, emoji string) *Slack {
	slackClient := slack.New(token)

	// Checking if the token is valid
	_, err := slackClient.AuthTest()
	if err != nil {
		fmt.Println(err)
	}
	slack := Slack{
		DoNotDisturb: dnd,
		client:       slackClient,
		Emoji:        emoji,
	}
	return &slack
}

func (s *Slack) GetName() string {
	return "slack"
}

func (s *Slack) Start(pomodoroDuration time.Duration, refreshRate time.Duration, message string) {
	pomodoroFinished := time.Now().Add(pomodoroDuration)
	err := s.client.SetUserCustomStatus(message, s.Emoji, pomodoroFinished.Unix())
	if err != nil {
		fmt.Println(err)
	}
	if s.DoNotDisturb {
		s.client.SetSnooze(int(pomodoroDuration.Minutes()))
	}

}

func (s *Slack) Refresh(timeLeft time.Duration) {
}

func (s *Slack) End() {
	s.client.SetUserCustomStatus("", "", 0)
	if s.DoNotDisturb {
		s.client.EndSnooze()
	}
}
