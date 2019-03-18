package pizzerias

import (
	"roob.re/pizzameeting"
	"roob.re/pizzameeting/doodle"
)

func FromDoodle(d *doodle.Doodle) pizzameeting.Pizzeria {
	return pizzameeting.PizzeriaFunc(func() []pizzameeting.Pizza {
		var menu []pizzameeting.Pizza
		for _, opt := range d.Options() {
			menu = append(menu, pizzameeting.Pizza(opt))
		}

		return menu
	})
}
