package stores

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

const (
	FALSE      = 0
	TRUE       = 1
	DATEFORMAT = "2006-01-02 15:04:05"
)

type SQLite struct {
	*sql.DB
}

func NewSQLite(dbPath string) (*SQLite, error) {
	db, err := sql.Open("sqlite3", dbPath)
	return &SQLite{db}, err
}

func (s *SQLite) Init() error {
	sql := `
CREATE TABLE IF NOT EXISTS pomodoro (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  message TEXT,
  stopped INTEGER,
  created_at DATETIME
);
`
	err := s.exec(sql)
	return err
}

func (s *SQLite) exec(query string, args ...string) error {
	tx, err := s.Begin()
	if err != nil {
		return err
	}
	sqlStatement, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer sqlStatement.Close()
	argsInterfaces := make([]interface{}, len(args))
	for i, v := range args {
		argsInterfaces[i] = v
	}
	_, err = sqlStatement.Exec(argsInterfaces...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *SQLite) AddPomodoro(p Pomodoro) error {
	date := p.Created_at.Format(DATEFORMAT)
	fmt.Println(date)

	err := s.exec(`INSERT INTO pomodoro(stopped, created_at, message) VALUES(0, ?, ?);`, date, p.Message)
	return err
}

func (s *SQLite) GetPomodoros() ([]Pomodoro, error) {
	rows, err := s.Query("select id, message, stopped, created_at from pomodoro")
	var pomodoros []Pomodoro
	if err != nil {
		return []Pomodoro{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var message string
		var stopped bool
		var createdAt time.Time
		rows.Scan(&id, &message, &stopped, &createdAt)
		// createdAtTime, _ := time.Parse("19-03-05 21:58:33", createdAt)
		pomodoros = append(pomodoros, Pomodoro{id, message, stopped, createdAt})
	}
	return pomodoros, nil
}

func (s *SQLite) GetPomodorosPerDay() (map[string][]Pomodoro, error) {
	return nil, nil
}

func (s *SQLite) GetLastPomodoro() (Pomodoro, error) {
	rows, err := s.Query("select id, message, stopped, created_at from pomodoro ORDER BY id DESC LIMIT 1")
	var pomodoro Pomodoro
	if err != nil {
		return Pomodoro{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var message string
		var stopped bool
		var created_at time.Time
		rows.Scan(&id, &message, &stopped, &created_at)
		pomodoro = Pomodoro{id, message, stopped, created_at}
	}
	return pomodoro, nil
}

func (s *SQLite) GetNextBreakDuration() (time.Duration, error) {
	today := time.Now().Format("2019-09-26")
	pomodorosPerDay, _ := s.GetPomodorosPerDay()
	pomodorosNumber := len(pomodorosPerDay[today])

	if pomodorosNumber%4 == 0 {
		return 15 * time.Minute, nil
	}

	return 5 * time.Minute, nil
}
