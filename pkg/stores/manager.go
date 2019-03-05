package stores

type StoreManager struct {
	Stores []*Store
}

func NewStoreManager() StoreManager {
	return StoreManager{}
}

func (s *StoreManager) Add(store *Store) {
	s.Stores = append(s.Stores, store)
}

func (s *StoreManager) Init() error {
	for _, store := range s.Stores {
		err := (*store).Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StoreManager) CreatePomodoro(message string) error {
	pomodoro := NewPomodoro(message)
	for _, store := range s.Stores {
		err := (*store).AddPomodoro(pomodoro)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StoreManager) GetPomodoros() ([]Pomodoro, error) {
	var pomodoros []Pomodoro
	for _, store := range s.Stores {
		pomodorosStore, err := (*store).GetPomodoros()
		if err != nil {
			return []Pomodoro{}, err
		}
		pomodoros = append(pomodoros, pomodorosStore...)
	}
	return pomodoros, nil
}

func (s *StoreManager) Stop() {
}
