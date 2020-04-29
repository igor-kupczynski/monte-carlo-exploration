package simulation

import (
	"monte-carlo-exploration/cointoss"
)

// Simulate runs the coin toss simulation.
//
// - initial is the initial state factory,
// - histories is the number of histories to simulate,
// - rounds is the number of rounds to play.
//
// It returns the end states of the simulated histories.
func Simulate(initial func() *cointoss.State, histories int, rounds int) []*cointoss.State {
	// Generate initial states
	states := make([]*cointoss.State, histories, histories)
	for i := 0; i < histories; i++ {
		states[i] = initial()
	}

	done := make(chan struct{})

	// Play for n rounds
	for i := 0; i < histories; i++ {
		state := states[i]
		go func() {
			state.Play(rounds)
			done <- struct{}{}
		}()
	}

	// Wait until done
	for i := 0; i < histories; i++ {
		<-done
	}

	return states
}
