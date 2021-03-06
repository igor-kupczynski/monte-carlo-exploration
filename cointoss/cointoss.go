// cointoss is a coin toss Monte Carlo experiment
//
// We toss a coin for n-rounds. Heads - we get a $1, tails - we lose a $1. If out capital reaches
// zero we are ruined. We can't play anymore.
package cointoss

import (
	"fmt"
	"sort"

	"github.com/igor-kupczynski/monte-carlo-exploration/montecarlo"
	"github.com/igor-kupczynski/monte-carlo-exploration/stats"
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

// Results returns the experiment summary
func (e *experiment) Results() fmt.Stringer {
	total := len(e.states)
	sort.Sort(stateSlice(e.states))

	firstNotRuined := sort.Search(total, func(i int) bool { return !e.states[i].ruined })

	results := make([]int64, total)
	for i, s := range e.states {
		results[i] = int64(s.capital)
	}
	baseline := int64(e.states[0].initialCapital)

	return &Results{
		summary:    stats.Describe(results, baseline),
		procRuined: 100 * float64(firstNotRuined) / float64(total),
	}
}
