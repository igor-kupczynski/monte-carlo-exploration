package cointoss

import (
	"reflect"
	"sort"
	"testing"
)

func TestCoinTossState_nextRound(t *testing.T) {
	tests := []struct {
		name    string
		process *State
		heads   bool
		want    *State
	}{
		{
			name: "Winning round on heads",
			process: &State{
				Capital:     10,
				Ruined:      false,
				LastRoundNo: 0,
			},
			heads: true,
			want: &State{
				Capital:     11,
				Ruined:      false,
				LastRoundNo: 1,
			},
		},
		{
			name: "Loosing round on tails",
			process: &State{
				Capital:     11,
				Ruined:      false,
				LastRoundNo: 1,
			},
			heads: false,
			want: &State{
				Capital:     10,
				Ruined:      false,
				LastRoundNo: 2,
			},
		},
		{
			name: "Ruinous round on low Capital and tails",
			process: &State{
				Capital:     1,
				Ruined:      false,
				LastRoundNo: 3,
			},
			heads: false,
			want: &State{
				Capital:     0,
				Ruined:      true,
				LastRoundNo: 4,
			},
		},
		{
			name: "Skip round if Ruined",
			process: &State{
				Capital:     0,
				Ruined:      true,
				LastRoundNo: 4,
			},
			heads: true,
			want: &State{
				Capital:     0,
				Ruined:      true,
				LastRoundNo: 4,
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

func TestByCapital(t *testing.T) {
	highCapital := &State{
		Capital:     100,
		Ruined:      false,
		LastRoundNo: 100,
	}
	lowCapital := &State{
		Capital:     1,
		Ruined:      false,
		LastRoundNo: 100,
	}
	ruinedEarly := &State{
		Capital:     0,
		Ruined:      true,
		LastRoundNo: 20,
	}
	ruinedLate := &State{
		Capital:     0,
		Ruined:      true,
		LastRoundNo: 90,
	}
	xs := []*State{highCapital, lowCapital, ruinedLate, ruinedEarly}
	want := []*State{ruinedEarly, ruinedLate, lowCapital, highCapital}
	sort.Sort(ByCapital(xs))
	if !reflect.DeepEqual(xs, want) {
		t.Errorf("Got: %v, want: %v", xs, want)
	}
}
