package cointoss

import (
	"reflect"
	"sort"
	"testing"
)

func Test_state_Run(t *testing.T) {

	tests := []struct {
		name  string
		state *state
		want  *state
	}{
		{
			name: "should make progress for wantRounds",
			state: &state{
				capital:     10,
				ruined:      false,
				lastRoundNo: 0,

				wantRounds:     5,
				initialCapital: 10,
			},
			want: &state{
				ruined:      false,
				lastRoundNo: 5,

				wantRounds:     5,
				initialCapital: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.state.Run()

			// ignore the end capital
			tt.want.capital = tt.state.capital

			if !reflect.DeepEqual(tt.state, tt.want) {
				t.Errorf("Got: %+v, want: %+v", tt.state, tt.want)
			}
		})
	}
}

func Test_state_nextRound(t *testing.T) {
	tests := []struct {
		name    string
		process *state
		heads   bool
		want    *state
	}{
		{
			name: "Winning round on heads",
			process: &state{
				capital:     10,
				ruined:      false,
				lastRoundNo: 0,
			},
			heads: true,
			want: &state{
				capital:     11,
				ruined:      false,
				lastRoundNo: 1,
			},
		},
		{
			name: "Loosing round on tails",
			process: &state{
				capital:     11,
				ruined:      false,
				lastRoundNo: 1,
			},
			heads: false,
			want: &state{
				capital:     10,
				ruined:      false,
				lastRoundNo: 2,
			},
		},
		{
			name: "Ruinous round on low capital and tails",
			process: &state{
				capital:     1,
				ruined:      false,
				lastRoundNo: 3,
			},
			heads: false,
			want: &state{
				capital:     0,
				ruined:      true,
				lastRoundNo: 4,
			},
		},
		{
			name: "Skip round if ruined",
			process: &state{
				capital:     0,
				ruined:      true,
				lastRoundNo: 4,
			},
			heads: true,
			want: &state{
				capital:     0,
				ruined:      true,
				lastRoundNo: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.process.nextRound(tt.heads)
			if !reflect.DeepEqual(tt.process, tt.want) {
				t.Errorf("Got: %+v, want: %+v", tt.process, tt.want)
			}
		})
	}
}

func test_state_sort_order(t *testing.T) {
	highCapital := &state{
		capital:     100,
		ruined:      false,
		lastRoundNo: 100,
	}
	lowCapital := &state{
		capital:     1,
		ruined:      false,
		lastRoundNo: 100,
	}
	ruinedEarly := &state{
		capital:     0,
		ruined:      true,
		lastRoundNo: 20,
	}
	ruinedLate := &state{
		capital:     0,
		ruined:      true,
		lastRoundNo: 90,
	}
	xs := []*state{highCapital, lowCapital, ruinedLate, ruinedEarly}
	want := []*state{ruinedEarly, ruinedLate, lowCapital, highCapital}
	sort.Sort(stateSlice(xs))
	if !reflect.DeepEqual(xs, want) {
		t.Errorf("Got: %v, want: %v", xs, want)
	}
}
