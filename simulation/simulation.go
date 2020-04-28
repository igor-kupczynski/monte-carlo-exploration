package simulation

import (
	"monte-carlo-exploration/cointoss"
	"sort"
)

// Simulate runs the coin toss simulation.
//
// - initial is the initial state factory,
// - histories is the number of histories to simulate,
// - rounds is the number of rounds to play.
//
// It returns the end states of the simiulated histories sorted ByCapital.
func Simulate(initial func() *cointoss.State, histories int, rounds int) []*cointoss.State {
	// Generate initial states
	states := make([]*cointoss.State, histories, histories)
	for i := 0; i < histories; i++ {
		states[i] = initial()
	}

	// Play for n rounds
	// TODO: In goroutines
	for i := 0; i < histories; i++ {
		states[i].Play(rounds)
	}

	// Sort them by fitness (capital)
	sort.Sort(cointoss.ByCapital(states))

	return states
}
