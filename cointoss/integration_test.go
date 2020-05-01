package cointoss

import (
	"testing"

	"monte-carlo-exploration/montecarlo"
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
			return r.procLessCapital > r.procMoreCapital
		},
		"lower percentiles at $0": func(r *Results) bool {
			return r.percentiles[1] == 0 &&
				r.percentiles[5] == 0 &&
				r.percentiles[10] == 0 &&
				r.percentiles[25] == 0
		},
		"median at $10": func(r *Results) bool {
			return r.percentiles[50] == 10
		},
		"higher percentiles with more capital": func(r *Results) bool {
			return r.percentiles[75] > 10 &&
				r.percentiles[90] > r.percentiles[75] &&
				r.percentiles[95] > r.percentiles[90] &&
				r.percentiles[99] > r.percentiles[95] &&
				r.percentiles[99] > 30
		},
	} {
		if !property(results) {
			t.Errorf("Property doesn't hold: %s\nResults:\n%s", description, results)
		}
	}
}
