package pizzameeting

import (
	"math"
)

const minScore = 10
const maxScore = 2*minScore - 1

const repeatPenaltyFactor = 0.66 // Each repetition will multiply the score by this value. The closer to one, the smaller the penalty.

type Attendee interface {
	Evaluate([]Pizza) int
}

type Pizza string

type Score struct {
	pizza Pizza
	score int
}

type Person struct {
	Id          Identifier `json:"-"`
	Name        string
	scores      map[Pizza]int
	scoresSlice []Score
}

func (p *Person) Score(pizza Pizza, score int) {
	if p.scores == nil {
		p.scores = make(map[Pizza]int)
	}
	p.scores[pizza] = score

	p.scoresSlice = append(p.scoresSlice, Score{pizza, score})
}

func (p *Person) pizzaScore(pizza Pizza) int {
	if len(p.scoresSlice) < 16 {
		for _, score := range p.scoresSlice {
			if score.pizza == pizza {
				return score.score
			}
		}
		return 0
	} else {
		return p.scores[pizza]
	}
}

func (p *Person) Evaluate(menu []Pizza) int {
	score := 0.0
	//repeatPenalties := make(map[Pizza]float32, len(menu))
	penalized := make([]Pizza, 0, len(menu))

	for _, pizza := range menu {
		score += float64(p.scores[pizza]) * penalty(pizza, penalized)
		penalized = append(penalized, pizza)
	}

	if len(p.scores) >= 2 {
		return int(math.Round(score - minScore*(1+repeatPenaltyFactor)))
	} else {
		return int(math.Round(score)) - (minScore - 1)
	}
}

func penalty(pizza Pizza, penalties []Pizza) float64 {
	factor := 1.0
	for _, penalizedPizza := range penalties {
		if penalizedPizza == pizza {
			factor *= repeatPenaltyFactor
		}
	}

	return factor
}
