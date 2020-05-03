package pi

import (
	"testing"
)

func Test_state_Run(t *testing.T) {
	tests := []struct {
		name         string
		img          [][]bool
		noRounds     int
		lowEstimate  float64
		highEstimate float64
	}{
		{
			name: "Only hits",
			img: [][]bool{
				{true, true, true, true},
				{true, true, true, true},
				{true, true, true, true},
				{true, true, true, true},
			},
			noRounds:     10,
			lowEstimate:  1,
			highEstimate: 1,
		},
		{
			name: "Only miss",
			img: [][]bool{
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			},
			noRounds:     10,
			lowEstimate:  0,
			highEstimate: 0,
		},
		{
			name: "Simple checker",
			img: [][]bool{
				{true, false},
				{false, true},
			},
			noRounds:     1000,
			lowEstimate:  0.45,
			highEstimate: 0.55,
		},
		{
			name: "6x6 circle",
			img: [][]bool{
				{false, false, true, true, false, false},
				{false, true, true, true, true, false},
				{true, true, true, true, true, true},
				{true, true, true, true, true, true},
				{false, true, true, true, true, false},
				{false, false, true, true, false, false},
			},
			noRounds:     1000,
			lowEstimate:  0.62, // 2/3 is the exact answer
			highEstimate: 0.72,
		},
		{
			name: "Go logo",
			img: [][]bool{
				{false, true, true, false, false, false, true, false},
				{true, false, false, true, false, true, false, true},
				{true, false, false, false, false, true, false, true},
				{true, false, true, true, false, true, false, true},
				{false, true, true, false, false, false, true, false},
			},
			noRounds:     10000,
			lowEstimate:  0.41, // 45% is the exact answer
			highEstimate: 0.49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state{
				img:        tt.img,
				wantRounds: tt.noRounds,
			}
			s.Run()
			gotEstimate := float64(s.hit) / float64(s.total)
			if gotEstimate < tt.lowEstimate || gotEstimate > tt.highEstimate {
				t.Errorf("Got estimate: %f, want in range: [%f, %f]",
					gotEstimate, tt.lowEstimate, tt.highEstimate)
			}
		})
	}
}
