package pizzerias

import "roob.re/pizzameeting"

var AlTaglioBCN = pizzameeting.Pizzeria(pizzameeting.PizzeriaFunc(taglioMenu))

var alTaglioMenu = []pizzameeting.Pizza{
	"Buffala",
	"Margherita",
	"Diavola",
	"4Fromatges",
	"Prosciutto",
	"Parmesana",
	"Boscaiola",
	"Coppata",
	"Vegetariana",
	"Rúcula",
	"BBQ",
	"Tedesca",
	"Tonno",
	"Carbonara",
	"Rianella",
	"Huevos",
}

func taglioMenu() []pizzameeting.Pizza {
	return alTaglioMenu
}
