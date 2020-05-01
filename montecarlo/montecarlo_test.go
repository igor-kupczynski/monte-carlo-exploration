package montecarlo

import (
	"fmt"
	"reflect"
	"testing"
)

type mockSample struct {
	run bool
}

func (s *mockSample) Run() {
	s.run = true
}

type mockResult struct {
	runSamples []bool
}

func (r *mockResult) String() string {
	return fmt.Sprintf("%v", r.runSamples)
}

type mockExperiment struct {
	samples []*mockSample
}

func (e *mockExperiment) Samples() []Sample {
	r := make([]Sample, len(e.samples))
	for i, s := range e.samples {
		r[i] = Sample(s)
	}
	return r
}

func (e *mockExperiment) Results() fmt.Stringer {
	rs := make([]bool, len(e.samples))
	for i, s := range e.samples {
		rs[i] = s.run
	}
	return &mockResult{rs}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		exp  Experiment
		want fmt.Stringer
	}{
		{
			name: "run each of the samples",
			exp:  &mockExperiment{samples: []*mockSample{{}, {}, {}}},
			want: &mockResult{runSamples: []bool{true, true, true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.exp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
