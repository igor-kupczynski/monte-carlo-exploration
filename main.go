package main

import (
	"fmt"
	"monte-carlo-exploration/cointoss"
	"monte-carlo-exploration/simulation"
	"sort"
)

func main() {

	// Run 100 games for 100 rounds; starting capital 10
	histories := 1000000
	rounds := 100
	startingCapital := 10

	states := simulation.Simulate(
		func() *cointoss.State {
			return &cointoss.State{
				Capital:     startingCapital,
				Ruined:      false,
				LastRoundNo: 0,
			}
		},
		histories,
		rounds,
	)

	// Sort the results by capital
	sort.Sort(cointoss.ByCapital(states))

	firstNotRuined := sort.Search(histories, func(i int) bool { return !states[i].Ruined })
	procRuined := 100 * float64(firstNotRuined) / float64(histories)
	firstSameCapital := sort.Search(histories, func(i int) bool { return states[i].Capital >= startingCapital })
	procLessCapital := 100 * float64(firstSameCapital) / float64(histories)

	wantPercentiles := []int{1, 5, 10, 50, 90, 95, 99}
	percentiles := make(map[int]int)
	for _, p := range wantPercentiles {
		percentiles[p] = states[p*histories/100].Capital
	}

	fmt.Printf("# Simulating %d executions of %d round coin toss with starting capital %d\n", histories, rounds, startingCapital)
	fmt.Printf("Ruined: %f%% (%d / %d)\n", procRuined, firstNotRuined, histories)
	fmt.Printf("Less capital: %f%% (%d / %d)\n", procLessCapital, firstSameCapital, histories)
	for _, p := range wantPercentiles {
		fmt.Printf("p%02d captial: %d\n", p, percentiles[p])
	}
}
