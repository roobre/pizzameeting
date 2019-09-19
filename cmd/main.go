package main

import (
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"roob.re/pizzameeting"
	"roob.re/pizzameeting/combinator"
	"roob.re/pizzameeting/doodle"
	"roob.re/pizzameeting/pizzerias"
	"runtime/pprof"
)

func main() {

	doodleId := flag.String("doodle", "", "Doodle URL or ID")
	ppp := flag.Float64("pizzas-per-person", 0.5, "Pizzas per person")
	prof := flag.String("prof", "", "Write profile output to this file")
	flag.Parse()

	if *prof != "" {
		f, err := os.Create(*prof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if len(*doodleId) < 16 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	meeting := pizzameeting.Meeting{}

	log.Println("Fetching doodle info...")
	dood, err := doodle.ParseDoodle(*doodleId)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error fetching doodle information: %v", err)
		os.Exit(1)
	}

	meeting.Pizzeria = pizzerias.FromDoodle(dood)
	meeting.Solver = pizzameeting.PPPSolver{
		CombinatorMaker: combinator.RecursiveCombinatorMaker,
		PizzasPerPerson: *ppp,
	}

	for participant, votes := range dood.Results() {
		p := &pizzameeting.Person{Name: string(participant)}
		for pizza, score := range votes {
			if score != 0 {
				p.Score(pizzameeting.Pizza(pizza), 10+((score-1)*6))
			}
		}

		meeting.Invite(p)
	}

	out := make(map[pizzameeting.Pizza]int)
	log.Println("Computing menu...")
	for _, pizza := range meeting.Menu() {
		out[pizza]++
	}

	tbw := tablewriter.NewWriter(os.Stdout)
	tbw.SetHeader([]string{"Pizza", "#"})

	for pizza, number := range out {
		tbw.Append([]string{string(pizza), fmt.Sprint(number)})
	}

	tbw.Render()
}
