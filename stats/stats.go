// Package stats offers convenience functions to calculate descriptive result properties
package stats

import (
	"fmt"
	"sort"
	"strings"
)

var wantPercentiles = []int{1, 5, 10, 25, 50, 75, 90, 95, 99}

// Summary represents the basic statistics on the given dataset
type Summary struct {
	Total    int
	Baseline int64
	Avg      float64
	Min, Max int64
	// Percentage of the dataset below, above, and at the baseline
	Below, At, Above float64
	Percentiles      map[int]int64
	PercentileDiffs  map[int]float64
}

func (s *Summary) String() string {
	var buf strings.Builder

	buf.WriteString(fmt.Sprintf("Dataset [len=%d, baseline=%d]\n", s.Total, s.Baseline))
	buf.WriteString(fmt.Sprintf("* Avg=%f\tMin=%d\tMax%d\n", s.Avg, s.Min, s.Max))
	buf.WriteString(fmt.Sprintf("* %% of items below=%f%%\tat=%f%%\tabove=%f%% baseline\n", s.Below, s.At, s.Above))
	buf.WriteString(fmt.Sprintf("* Percentiles:\n"))

	for _, p := range wantPercentiles {
		buf.WriteString(fmt.Sprintf("\t- p%02d%%: %d\tbaseline diff: %f%%\n", p, s.Percentiles[p], s.PercentileDiffs[p]))
	}

	return buf.String()
}

// Describe provides basic statistics on the given dataset.
//
// The dataset will be sorted in place
func Describe(results []int64, baseline int64) *Summary {
	total := len(results)
	Int64Slice(results).Sort()

	var sum int64
	for _, x := range results {
		sum += x
	}
	avg := float64(sum) / float64(total)

	firstAt := sort.Search(total, func(i int) bool { return results[i] >= baseline })
	firstAbove := sort.Search(total, func(i int) bool { return results[i] > baseline })

	percentiles := make(map[int]int64)
	pErrors := make(map[int]float64)
	for _, p := range wantPercentiles {
		x := results[p*total/100]
		percentiles[p] = x
		pErrors[p] = float64(x-baseline) * 100 / float64(baseline)
	}

	return &Summary{
		Total:           total,
		Baseline:        baseline,
		Min:             results[0],
		Max:             results[total-1],
		Avg:             avg,
		Below:           100 * float64(firstAt) / float64(total),
		At:              100 * float64(firstAbove-firstAt) / float64(total),
		Above:           100 * float64(total-firstAbove) / float64(total),
		Percentiles:     percentiles,
		PercentileDiffs: pErrors,
	}
}
