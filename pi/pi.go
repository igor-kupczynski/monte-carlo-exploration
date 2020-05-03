// Pi estimates the value of Pi using Monte Carlo method.
//
// We throw darts into a circle. We can estimate the area of the circle as area of the image * hits / miss. "Throwing a
//dart" means getting random x / y directions.
//
// Circle is represented by a GIMP-made 1024x1024 png image. Circle is black and the background is white. There is
// some gradient on the border. We consider 50%+ gray to be black.
package pi

import (
	"fmt"
	_ "image/png"
	"math"

	"github.com/igor-kupczynski/monte-carlo-exploration/montecarlo"
)

// experiment is the dart throwing Pi estimation Monte Carlo experiment
type experiment struct {
	states []*state
}

func (e *experiment) Samples() []montecarlo.Sample {
	samples := make([]montecarlo.Sample, len(e.states))
	for i, state := range e.states {
		samples[i] = montecarlo.Sample(state)
	}
	return samples
}

func (e *experiment) Results() fmt.Stringer {
	hits := make([]int, len(e.states))
	totals := make([]int, len(e.states))
	pis := make([]float64, len(e.states))
	errors := make([]float64, len(e.states))

	for i, s := range e.states {
		hits[i] = s.hit
		totals[i] = s.total
		pis[i] = 4 * float64(s.hit) / float64(s.total)
		errors[i] = math.Abs(math.Pi - pis[i]) // We cheat here a little; we know what Pi is, so we calculate the error
	}

	return &results{
		hit:   hits,
		total: totals,
		pi:    pis,
		err:   errors,
	}
}
