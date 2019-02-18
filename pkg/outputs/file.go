package outputs

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type File struct {
	Path string
}

func NewFile(path string) *File {
	file := File{
		Path: path,
	}
	return &file
}

func (s *File) GetName() string {
	return "file"
}

func (s *File) Start(pomodoroDuration time.Duration, refreshRate time.Duration, message string) {
	ioutil.WriteFile(s.Path, []byte(pomodoroDuration.String()), 0644)
}

func (s *File) Refresh(timeLeft time.Duration) {
	ioutil.WriteFile(s.Path, []byte(formatDuration(timeLeft)), 0644)
}

func (s *File) End() {
	os.Remove(s.Path)
}

func formatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d", m, s)
}
