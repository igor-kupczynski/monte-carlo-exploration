package cointoss

import (
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	tests := []struct {
		name           string
		histories      int
		rounds         int
		initialCapital int
		want           *experiment
	}{
		{
			name:           "Should create coin toss experiment with the params",
			histories:      3,
			rounds:         5,
			initialCapital: 7,
			want: &experiment{
				states: []*state{
					{
						capital:        7,
						ruined:         false,
						lastRoundNo:    0,
						wantRounds:     5,
						initialCapital: 7,
					},
					{
						capital:        7,
						ruined:         false,
						lastRoundNo:    0,
						wantRounds:     5,
						initialCapital: 7,
					},
					{
						capital:        7,
						ruined:         false,
						lastRoundNo:    0,
						wantRounds:     5,
						initialCapital: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Args{
				Histories:      tt.histories,
				Rounds:         tt.rounds,
				InitialCapital: tt.initialCapital,
			}
			if got := New(a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_experiment_Results(t *testing.T) {
	tests := []struct {
		name   string
		states []*state
		want   *Results
	}{
		{
			name: "simplified example",
			states: []*state{
				{
					capital:     10,
					ruined:      false,
					lastRoundNo: 6,

					wantRounds:     6,
					initialCapital: 10,
				},
			},
			want: &Results{
				// we are not interested in Results.summary here
				procRuined: 0,
			},
		},
		{
			name:   "increasing capital",
			states: incCap(),
			want: &Results{
				procRuined: 1,
			},
		},
		{
			name:   "30% ruined",
			states: ruined30(),
			want: &Results{
				procRuined: 30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &experiment{
				states: tt.states,
			}
			got := e.Results().(*Results)
			// we don't want to compare got.summary since this is tested elsewhere
			tt.want.summary = got.summary
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Results() = %v, want %v", got, tt.want)
			}
		})
	}
}

// incCap returns 100 samples with increasing capital $0 to $99
func incCap() []*state {
	total := 100
	states := make([]*state, total)

	for i := 0; i < total; i++ {
		states[i] = &state{
			capital:        i,
			ruined:         i == 0,
			lastRoundNo:    total,
			wantRounds:     total,
			initialCapital: total / 2,
		}
	}

	return states
}

// ruined30 returns 100 samples: 30% is ruined and 70% has initial capital
func ruined30() []*state {
	total := 100
	ruined := 30
	initial := 10
	states := make([]*state, total)

	for i := 0; i < total; i++ {
		var captial int
		if i >= ruined {
			captial = initial
		}

		states[i] = &state{
			capital:        captial,
			ruined:         captial == 0,
			lastRoundNo:    4,
			wantRounds:     4,
			initialCapital: initial,
		}
	}

	return states
}
