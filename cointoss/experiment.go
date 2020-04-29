package cointoss

import (
	"fmt"
	"sort"

	"monte-carlo-exploration/simulation"
)

// Experiment is a collection of coin toss instances we want to simulate
type Experiment []*State

// Samples returns coin toss instances as slice of samples to be simulated
func (e Experiment) Samples() []simulation.Sample {
	samples := make([]simulation.Sample, len(e))
	for i, state := range e {
		samples[i] = simulation.Sample(state)
	}
	return samples
}

// PrintReport prints the capital and ruin report to stdout
func (e Experiment) PrintReport(initialCapital int) {
	total := len(e)

	// Sort the results by capital
	sort.Sort(e)

	firstNotRuined := sort.Search(total, func(i int) bool { return !e[i].Ruined })
	procRuined := 100 * float64(firstNotRuined) / float64(total)
	firstSameCapital := sort.Search(total, func(i int) bool { return e[i].Capital >= initialCapital })
	procLessCapital := 100 * float64(firstSameCapital) / float64(total)
	firstMoreCapital := sort.Search(total, func(i int) bool { return e[i].Capital > initialCapital })
	procMoreCapital := 100 * float64(total-firstMoreCapital) / float64(total)

	wantPercentiles := []int{1, 5, 10, 50, 90, 95, 99}
	percentiles := make(map[int]int)
	for _, p := range wantPercentiles {
		percentiles[p] = e[p*total/100].Capital
	}

	fmt.Printf("Ruined: %f%% (%d / %d)\n", procRuined, firstNotRuined, total)
	fmt.Printf("Less capital: %f%% (%d / %d)\n", procLessCapital, firstSameCapital, total)
	fmt.Printf("More capital: %f%% (%d / %d)\n", procMoreCapital, total-firstMoreCapital, total)

	for _, p := range wantPercentiles {
		fmt.Printf("p%02d $%d\n", p, percentiles[p])
	}
}

func (e Experiment) Len() int { return len(e) }

func (e Experiment) Less(i, j int) bool {
	if e[i].Ruined && e[j].Ruined {
		return e[i].TotalRounds < e[j].TotalRounds
	}
	return e[i].Capital < e[j].Capital
}

func (e Experiment) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
