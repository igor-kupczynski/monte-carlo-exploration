package cointoss

import (
	"fmt"
	"strings"

	"github.com/igor-kupczynski/monte-carlo-exploration/stats"
)

// Results represents coin toss experiment Results
type Results struct {
	summary *stats.Summary

	procRuined float64
}

func (r *Results) String() string {
	var buf strings.Builder

	buf.WriteString(r.summary.String())
	buf.WriteString(fmt.Sprintf("* %% ruined: %f%%\n", r.procRuined))

	return buf.String()
}
