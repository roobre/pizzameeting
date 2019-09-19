package pizzameeting

import (
	"math"
)

const minScore = 10
const maxScore = 2*minScore - 1

const repeatPenaltyFactor = 0.60 // Score will be multiplied by 1-repeatPenaltyFactor. The higher the factor, the more the penalty

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
	var score float32 = 0
	//repeatPenalties := make(map[Pizza]float32, len(menu))
	penalized := make([]Pizza, 0, len(menu))

	for _, pizza := range menu {
		// Accessing maps is cheap, appending is not
		score += float32(p.scores[pizza]) * penalty(pizza, penalized)
		//score += float32(p.scores[pizza]) * (repeatPenalties[pizza] + 1)
		//score += float32(p.pizzaScore(pizza)) * (repeatPenalties[pizza] + 1)
		penalized = append(penalized, pizza)
		//repeatPenalties[pizza] = -repeatPenaltyFactor
	}

	if len(p.scores) >= 2 {
		return int(math.Round(float64(score - minScore*(1+repeatPenaltyFactor))))
	} else {
		return int(math.Round(float64(score))) - (minScore - 1)
	}
}

func penalty(pizza Pizza, penalties []Pizza) float32 {
	for _, penalizedPizza := range penalties {
		if penalizedPizza == pizza {
			return 1 - repeatPenaltyFactor
		}
	}

	return 1
}
