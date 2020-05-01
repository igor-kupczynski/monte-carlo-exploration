package cointoss

import (
	"fmt"
	"strings"
)

var wantPercentiles = []int{1, 5, 10, 25, 50, 75, 90, 95, 99}

// Results represents coin toss experiment Results
type Results struct {
	total int

	firstNotRuined int
	procRuined     float64

	firstSameCapital int
	procLessCapital  float64

	firstMoreCapital int
	procMoreCapital  float64

	percentiles map[int]int
}

func (r *Results) String() string {
	var buf strings.Builder

	buf.WriteString(fmt.Sprintf("ruined: %f%% (%d / %d)\n", r.procRuined, r.firstNotRuined, r.total))
	buf.WriteString(fmt.Sprintf("Less capital: %f%% (%d / %d)\n", r.procLessCapital, r.firstSameCapital, r.total))
	buf.WriteString(fmt.Sprintf("More capital: %f%% (%d / %d)\n",
		r.procMoreCapital, r.total-r.firstMoreCapital, r.total))

	for _, p := range wantPercentiles {
		buf.WriteString(fmt.Sprintf("p%02d $%d\n", p, r.percentiles[p]))
	}

	return buf.String()
}
