package pizzameeting

import "log"

type Meeting struct {
	Pizzeria        Pizzeria
	Solver          Solver
	PizzasPerPerson float64
	attendees       []Attendee
}

func (m *Meeting) Invite(attendees ...Attendee) {
	for _, a := range attendees {
		m.attendees = append(m.attendees, a)
	}
}

func (m *Meeting) Menu() []Pizza {
	if m.Solver == nil {
		return nil
	}

	menus := m.Solver.Solve(m.attendees, m.Pizzeria.Menu())

	log.Printf("Generated %d acceptable menus and %d optimal menus\n", len(menus.Acceptable), len(menus.Optimal))

	if len(menus.Optimal) > 0 {
		return menus.Optimal[0]
	} else {
		return nil
	}
}
