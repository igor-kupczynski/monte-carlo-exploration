package pi

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_experiment_Results(t *testing.T) {
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
			want: &results{
				hit:   []int{7854},
				total: []int{10000},
				pi:    []float64{3.1416},
				err:   []float64{3.1416 - math.Pi},
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
