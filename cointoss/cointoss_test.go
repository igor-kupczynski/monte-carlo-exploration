package cointoss

import (
	"fmt"
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
		want   fmt.Stringer
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
				total:            1,
				firstNotRuined:   0,
				procRuined:       0,
				firstSameCapital: 0,
				procLessCapital:  0,
				firstMoreCapital: 1,
				procMoreCapital:  0,
				percentiles: map[int]int{
					1:  10,
					5:  10,
					10: 10,
					25: 10,
					50: 10,
					75: 10,
					90: 10,
					95: 10,
					99: 10,
				},
			},
		},
		{
			name:   "increasing capital",
			states: incCap(),
			want: &Results{
				total:            100,
				firstNotRuined:   1,
				procRuined:       1,
				firstSameCapital: 50,
				procLessCapital:  50,
				firstMoreCapital: 51,
				procMoreCapital:  49,
				percentiles: map[int]int{
					1:  1,
					5:  5,
					10: 10,
					25: 25,
					50: 50,
					75: 75,
					90: 90,
					95: 95,
					99: 99,
				},
			},
		},
		{
			name:   "30% ruined",
			states: ruined30(),
			want: &Results{
				total:            100,
				firstNotRuined:   30,
				procRuined:       30,
				firstSameCapital: 30,
				procLessCapital:  30,
				firstMoreCapital: 100,
				procMoreCapital:  0,
				percentiles: map[int]int{
					1:  0,
					5:  0,
					10: 0,
					25: 0,
					50: 10,
					75: 10,
					90: 10,
					95: 10,
					99: 10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &experiment{
				states: tt.states,
			}
			if got := e.Results(); !reflect.DeepEqual(got, tt.want) {
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
