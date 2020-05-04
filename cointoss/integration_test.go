package cointoss

import (
	"testing"

	"github.com/igor-kupczynski/monte-carlo-exploration/montecarlo"
)

func Test_cointoss_integration(t *testing.T) {
	// Run the cointoss game
	cfg := &Args{
		Histories:      10000,
		Rounds:         100,
		InitialCapital: 10,
	}
	experiment := New(cfg)
	results := montecarlo.Run(experiment).(*Results)

	// Verify basic properties of the results
	for description, property := range map[string]func(r *Results) bool{
		"ruin between 28% and 35% times": func(r *Results) bool {
			return r.procRuined > 28 && r.procRuined < 35
		},
		"chance of less capital > chance of more capital": func(r *Results) bool {
			return r.summary.Below > r.summary.Above
		},
		"lower percentiles at $0": func(r *Results) bool {
			return r.summary.Percentiles[1] == 0 &&
				r.summary.Percentiles[5] == 0 &&
				r.summary.Percentiles[10] == 0 &&
				r.summary.Percentiles[25] == 0
		},
		"median at $10": func(r *Results) bool {
			return r.summary.Percentiles[50] == 10
		},
		"higher percentiles with more capital": func(r *Results) bool {
			return r.summary.Percentiles[75] > 10 &&
				r.summary.Percentiles[90] > r.summary.Percentiles[75] &&
				r.summary.Percentiles[95] > r.summary.Percentiles[90] &&
				r.summary.Percentiles[99] > r.summary.Percentiles[95] &&
				r.summary.Percentiles[99] > 30
		},
	} {
		if !property(results) {
			t.Errorf("Property doesn't hold: %s\nResults:\n%s", description, results)
		}
	}
}
