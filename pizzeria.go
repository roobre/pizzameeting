package pizzameeting


type Pizzeria interface {
	Menu() []Pizza
}

type PizzeriaFunc func() []Pizza
func (pf PizzeriaFunc) Menu() []Pizza {
	return pf()
}
