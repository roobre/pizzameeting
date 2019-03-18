package pizzameeting

import (
	"math"
)

const minScore = 10
const maxScore = 2*minScore - 1

const penaltyFactor = 0.5

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
		repeatPenalties[pizza] = -penaltyFactor
	}

	if len(p.scores) >= 2 {
		return int(math.Round(float64(score - minScore*(1+penaltyFactor))))
	} else {
		return int(math.Round(float64(score))) - (minScore - 1)
	}
}
