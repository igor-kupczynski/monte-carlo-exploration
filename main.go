package main

import (
	"fmt"
	"monte-carlo-exploration/cointoss"
	"monte-carlo-exploration/simulation"
)

func main() {

	// Run 100 games for 100 rounds; starting capital 10
	states := simulation.Simulate(
		func() *cointoss.State {
			return &cointoss.State{
				Capital:     10,
				Ruined:      false,
				LastRoundNo: 0,
			}
		},
		100,
		100,
	)

	for i, state := range states {
		fmt.Printf("%02d: capital=$%3d, ruined=%t, last round=%d\n", i, state.Capital, state.Ruined, state.LastRoundNo)
	}
}
