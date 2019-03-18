package pizzameeting

import (
	"math"
)

const minScore = 10
const maxScore = 2*minScore - 1

const repeatPenaltyFactor = 0.66 // Score will be multiplied by 1-repeatPenaltyFactor. The higher the factor, the more the penalty

type Attendee interface {
	Evaluate([]Pizza) int
}

type Pizza string

type Person struct {
	Id     Identifier `json:"-"`
	Name   string
	scores map[Pizza]int
}

func (p *Person) Score(pi Pizza, score int) {
	if p.scores == nil {
		p.scores = make(map[Pizza]int)
	}

	p.scores[pi] = score
}

func (p *Person) Evaluate(menu []Pizza) int {
	var score float32 = 0
	repeatPenalties := make(map[Pizza]float32, 4)
	for _, pizza := range menu {
		score += float32(p.scores[pizza]) * (repeatPenalties[pizza] + 1)
		repeatPenalties[pizza] = -repeatPenaltyFactor
	}

	if len(p.scores) >= 2 {
		return int(math.Round(float64(score - minScore*(1+repeatPenaltyFactor))))
	} else {
		return int(math.Round(float64(score))) - (minScore - 1)
	}
}
