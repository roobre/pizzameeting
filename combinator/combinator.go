package combinator

type Combinator interface {
	Next() []int
	Len() int
}

type CombinatorMaker func(nSamples, maxSample int) Combinator

type RecursiveCombinator struct {
	maxSample int
	imax      int
	seed      []int
	current   []int
	done      bool
}

func NewRecursiveCombinatorMaker() func(nSamples, maxSample int) Combinator {
	return func(nSamples, maxSample int) Combinator {
		if nSamples < 1 {
			panic("Cannot generate combinator for < 1 samples")
		}

		c := &RecursiveCombinator{
			maxSample: maxSample,
			imax:      maxSample - 1,
		}
		c.current = make([]int, nSamples)
		c.current[0] = -1

		return c
	}
}

func (rc *RecursiveCombinator) Len() int {
	total := 1
	for i := rc.maxSample + len(rc.current) - 1; i >= rc.maxSample; i-- {
		total *= i
	}
	for i := len(rc.current); i > 1; i-- {
		total /= i
	}

	return total
}

func (rc *RecursiveCombinator) Next() []int {
	rc.increment(rc.current)
	return rc.current
}

func (rc *RecursiveCombinator) increment(comb []int) {
	if comb[0] < rc.imax {
		comb[0]++
	} else {
		if len(comb) > 1 {
			rc.increment(comb[1:])
			comb[0] = comb[1]
		}
	}
}
