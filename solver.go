package pizzameeting

import (
	"math"
	"roob.re/pizzameeting/combinator"
	"runtime"
	"sort"
	"sync"
)

type Solver interface {
	Solve(attendees []Attendee, menu []Pizza) Solution
	Use(maker combinator.CombinatorMaker)
}

type Solution struct {
	Acceptable [][]Pizza
	Optimal    [][]Pizza
}

type validMenu struct {
	menu  []Pizza
	score int
}

type PPPSolver struct {
	attendees       []Attendee
	PizzasPerPerson float64
	CombinatorMaker combinator.CombinatorMaker
}

func (ps PPPSolver) Use(maker combinator.CombinatorMaker) {
	ps.CombinatorMaker = maker
}

func (ps PPPSolver) Solve(attendees []Attendee, restaurantMenu []Pizza) Solution {
	ps.attendees = attendees

	generatedMenuChan := make(chan []Pizza, 4)
	acceptedMenuChan := make(chan *validMenu, 4)

	accpetedMenus := make([]*validMenu, 0)

	workerWg := &sync.WaitGroup{}
	// Spin up workers
	for i := 0; i < runtime.NumCPU(); i++ {
		workerWg.Add(1)
		go ps.worker(generatedMenuChan, acceptedMenuChan, workerWg)
	}

	collectorWg := &sync.WaitGroup{}
	// Spin up accepted menus collector
	go func() {
		collectorWg.Add(1)
		for vm := range acceptedMenuChan {
			accpetedMenus = append(accpetedMenus, vm)
		}
		collectorWg.Done()
	}()

	ppp := ps.PizzasPerPerson
	if ppp == 0 {
		ppp = 0.5
	}
	nPizzas := int(math.Ceil(float64(len(ps.attendees)) * ppp))

	comb := ps.CombinatorMaker(nPizzas, len(restaurantMenu))

	for i := comb.Len(); i > 0; i-- {
		generatedMenu := make([]Pizza, nPizzas)
		for menuIndex, pizzaIndex := range comb.Next() {
			generatedMenu[menuIndex] = restaurantMenu[pizzaIndex]
		}
		generatedMenuChan <- generatedMenu
	}

	close(generatedMenuChan)

	workerWg.Wait()         // Wait until all workers have finished
	close(acceptedMenuChan) // Workers finished, it's safe to close to let collector stop
	collectorWg.Wait()      // Wait for the collector to finish

	solution := Solution{}

	if len(accpetedMenus) == 0 {
		return solution
	}

	sort.Slice(accpetedMenus, func(i, j int) bool {
		return accpetedMenus[i].score > accpetedMenus[j].score
	})

	max := 0
	for _, vm := range accpetedMenus {
		solution.Acceptable = append(solution.Acceptable, vm.menu)
		if vm.score >= max {
			solution.Optimal = append(solution.Optimal, vm.menu)
			max = vm.score
		}
	}

	return solution
}

func (ps PPPSolver) worker(menus chan []Pizza, validMenus chan *validMenu, wg *sync.WaitGroup) {
	for menu := range menus {
		if score := ps.evaluateMenu(menu); score > 0 {
			validMenus <- &validMenu{menu, score}
		}
	}
	wg.Done()
}

func (ps PPPSolver) evaluateMenu(menu []Pizza) int {
	score := 0
	attScore := 0
	for _, att := range ps.attendees {
		attScore = att.Evaluate(menu)
		if attScore > 0 {
			score += attScore
		} else {
			return 0
		}
	}

	return score
}
