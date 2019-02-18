package pizzameeting

import (
	"fmt"
	"log"
	"math"
	"roob.re/pizzameeting/combinator"
	"runtime"
	"sync"
)

type Meeting struct {
	Pizzeria        Pizzeria
	CombinatorMaker combinator.CombinatorMaker
	PizzasPerPerson float64
	attendees       []Attendee
}

type validMenu struct {
	menu  []Pizza
	score int
}

func (m *Meeting) Invite(attendees ...Attendee) {
	for _, a := range attendees {
		m.attendees = append(m.attendees, a)
	}
}

func (m *Meeting) Menu() []Pizza {
	generatedMenuChan := make(chan []Pizza, 4)
	acceptedMenuChan := make(chan *validMenu, 4)

	accpetedMenus := make([]*validMenu, 0)

	wg := sync.WaitGroup{}

	// Spin up workers
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			m.worker(generatedMenuChan, acceptedMenuChan, &wg)
		}()
	}

	// Spin up accepted menus collector
	logback := make(chan string)
	go func() {
		accecptedScores := make([]int, 0)
		for vm := range acceptedMenuChan {
			accecptedScores = append(accecptedScores, vm.score)
			if len(accpetedMenus) == 0 {
				accpetedMenus = append(accpetedMenus, vm)
			} else {
				if accpetedMenus[0].score == vm.score {
					accpetedMenus = append(accpetedMenus, vm)
				} else if accpetedMenus[0].score < vm.score {
					accpetedMenus[0] = vm
				}
			}
		}
		logback <- fmt.Sprintf("Total of accepted menus: %d (%v)\n", len(accecptedScores), accecptedScores)
	}()

	pmenu := m.Pizzeria.Menu()
	ppp := m.PizzasPerPerson
	if ppp == 0 {
		ppp = 0.5
	}
	nPizzas := int(math.Ceil(float64(len(m.attendees)) * ppp))

	comb := m.CombinatorMaker(nPizzas, len(pmenu))

	for i := comb.Len(); i > 0; i-- {
		menu := make([]Pizza, nPizzas)
		for menuIndex, pizzaIndex := range comb.Next() {
			menu[menuIndex] = pmenu[pizzaIndex]
		}
		generatedMenuChan <- menu
	}

	close(generatedMenuChan)

	wg.Wait()               // Wait until all workers have finished
	close(acceptedMenuChan) // Workers finished, it's safe to close to let collector stop

	log.Print(<-logback)
	log.Printf("Total of optimal menus: %d\n", len(accpetedMenus))

	if len(accpetedMenus) == 0 {
		return nil
	} else {
		return accpetedMenus[0].menu
	}
}

func (m *Meeting) worker(menus chan []Pizza, validMenus chan *validMenu, wg *sync.WaitGroup) {
	for menu := range menus {
		if score := m.evaluateMenu(menu); score > 0 {
			validMenus <- &validMenu{menu, score}
		}
	}
	wg.Done()
}

func (m *Meeting) evaluateMenu(menu []Pizza) int {
	score := 0
	attScore := 0
	for _, att := range m.attendees {
		attScore = att.Evaluate(menu)
		if attScore > 0 {
			score += attScore
		} else {
			return 0
		}
	}

	return score
}

func (m *Meeting) validMenu(menu []Pizza, score int) {

}
