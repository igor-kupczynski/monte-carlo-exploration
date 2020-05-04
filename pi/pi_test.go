package pi

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/igor-kupczynski/monte-carlo-exploration/stats"
)

func Test_experiment_Results(t *testing.T) {
	scaled_3_1416 := int64(math.Round(3.1416 * Scale))
	diff := float64(scaled_3_1416-baseline) * 100 / float64(baseline)

	tests := []struct {
		name   string
		states []*state
		want   fmt.Stringer
	}{
		{
			name: "Should produce the results based on the samples",
			states: []*state{
				{
					hit:   7854,
					total: 10000,
				},
			},
			want: &stats.Summary{
				Total:    1,
				Baseline: baseline,
				Avg:      3.1416 * Scale,
				Min:      scaled_3_1416,
				Max:      scaled_3_1416,
				Below:    0,
				At:       0,
				Above:    100.0,
				Percentiles: map[int]int64{
					1:  scaled_3_1416,
					5:  scaled_3_1416,
					10: scaled_3_1416,
					25: scaled_3_1416,
					50: scaled_3_1416,
					75: scaled_3_1416,
					90: scaled_3_1416,
					95: scaled_3_1416,
					99: scaled_3_1416,
				},
				PercentileDiffs: map[int]float64{
					1:  diff,
					5:  diff,
					10: diff,
					25: diff,
					50: diff,
					75: diff,
					90: diff,
					95: diff,
					99: diff,
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
				t.Errorf("Results() = %v, want %v", spew.Sdump(got), spew.Sdump(tt.want))
			}
		})
	}
}
