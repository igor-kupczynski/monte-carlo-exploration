// cointoss is a coin toss Monte Carlo experiment
//
// We toss a coin for n-rounds. Heads - we get a $1, tails - we lose a $1. If out capital reaches
// zero we are ruined. We can't play anymore.
package cointoss

import (
	"fmt"
	"sort"

	"monte-carlo-exploration/montecarlo"
)

// Args represents the parameters of the coin toss game
type Args struct {
	Histories      int
	Rounds         int
	InitialCapital int `toml:"initial_capital"`
}

func (a *Args) String() string {
	return fmt.Sprintf("Simulating %d executions of %d round coin toss with starting capital of $%d\n",
		a.Histories, a.Rounds, a.InitialCapital)
}

// experiment is coin toss Monte Carlo experiment
type experiment struct {
	states []*state
}

// Returns new experiment based on the args
func New(args *Args) montecarlo.Experiment {
	states := make([]*state, args.Histories)
	for i := range states {
		states[i] = &state{
			capital:     args.InitialCapital,
			ruined:      false,
			lastRoundNo: 0,

			wantRounds:     args.Rounds,
			initialCapital: args.InitialCapital,
		}
	}

	return &experiment{states: states}
}

// Samples returns collection of coin toss game states as montecarlo.Samples
func (e *experiment) Samples() []montecarlo.Sample {
	samples := make([]montecarlo.Sample, len(e.states))
	for i, state := range e.states {
		samples[i] = montecarlo.Sample(state)
	}
	return samples
}

// results returns the experiment summary
func (e *experiment) Results() fmt.Stringer {
	total := len(e.states)
	sort.Sort(stateSlice(e.states))

	firstNotRuined := sort.Search(total, func(i int) bool { return !e.states[i].ruined })
	firstSameCapital := sort.Search(total, func(i int) bool { return e.states[i].capital >= e.states[i].initialCapital })
	firstMoreCapital := sort.Search(total, func(i int) bool { return e.states[i].capital > e.states[i].initialCapital })

	percentiles := make(map[int]int)
	for _, p := range wantPercentiles {
		percentiles[p] = e.states[p*total/100].capital
	}

	return &results{
		total: total,

		firstNotRuined: firstNotRuined,
		procRuined:     100 * float64(firstNotRuined) / float64(total),

		firstSameCapital: firstSameCapital,
		procLessCapital:  100 * float64(firstSameCapital) / float64(total),

		firstMoreCapital: firstMoreCapital,
		procMoreCapital:  100 * float64(total-firstMoreCapital) / float64(total),

		percentiles: percentiles,
	}
}
