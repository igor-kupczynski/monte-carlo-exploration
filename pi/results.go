package pi

import (
	"fmt"
	"strings"
)

// results of the Pi estimation.
//
// results are given per-sample
// TODO: store aggregated results instead of per-samples, e.g. min/max/avg and percentiles
type results struct {
	hit   []int
	total []int
	pi    []float64
	err   []float64
}

func (r *results) String() string {
	var buf strings.Builder

	for i := 0; i < len(r.pi); i++ {
		fmt.Printf("%02d: Estimated Pi = %f (err %f)\n", i, r.pi[i], r.err[i])
	}

	return buf.String()
}
