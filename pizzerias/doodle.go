package pizzerias

import (
	"roob.re/pizzameeting"
	"roob.re/pizzameeting/doodle"
)

type DoodlePizzeria struct {
	doodle *doodle.Doodle
}

func (dp *DoodlePizzeria) Menu() []pizzameeting.Pizza {
	var menu []pizzameeting.Pizza
	for _, opt := range dp.doodle.Options() {
		menu = append(menu, pizzameeting.Pizza(opt))
	}

	return menu
}

func (dp *DoodlePizzeria) StringMenu() []string {
	var menu []string
	for _, opt := range dp.doodle.Options() {
		menu = append(menu, string(opt))
	}

	return menu
}

func FromDoodle(d *doodle.Doodle) pizzameeting.Pizzeria {
	return &DoodlePizzeria{doodle: d}
}
