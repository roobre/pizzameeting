package pizzameeting

import (
	"encoding/json"
	"log"
	"time"
)

type Meeting struct {
	Identifier
	Date      time.Time
	Topic     string
	Pizzeria  Pizzeria
	Solver    Solver
	Attendees []Attendee
}

func (m *Meeting) Invite(attendees ...Attendee) {
	for _, a := range attendees {
		m.Attendees = append(m.Attendees, a)
	}
}

func (m *Meeting) Menu() []Pizza {
	if m.Solver == nil {
		return nil
	}

	menus := m.Solver.Solve(m.Attendees, m.Pizzeria.Menu())

	log.Printf("Generated %d acceptable menus and %d optimal menus\n", len(menus.Acceptable), len(menus.Optimal))

	if len(menus.Optimal) > 0 {
		return menus.Optimal[0]
	} else {
		return nil
	}
}

func (m *Meeting) MarshalJSON() ([]byte, error) {
	var jsonmeeting struct {
		Identifier
		Date      time.Time  `json:"date"`
		Topic     string     `json:"topic"`
		Attendees []Attendee `json:"attendees"`
		Menu      []Pizza    `json:"menu"`
	}
	jsonmeeting.Identifier = m.Identifier
	jsonmeeting.Attendees = m.Attendees
	jsonmeeting.Date = m.Date
	jsonmeeting.Menu = m.Menu()

	return json.Marshal(jsonmeeting)
}

func (m *Meeting) UnmarshalJSON(raw []byte) error {
	var jsonmeeting struct {
		Date  time.Time `json:"date"`
		Topic string    `json:"topic"`
	}
	err := json.Unmarshal(raw, &jsonmeeting)
	if err != nil {
		return err
	}

	m.Date = jsonmeeting.Date
	m.Topic = jsonmeeting.Topic

	return nil
}
