package pi

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/igor-kupczynski/monte-carlo-exploration/montecarlo"
)

func TestNew(t *testing.T) {

	_, testFname, _, _ := runtime.Caller(0)

	tests := []struct {
		name    string
		args    *Args
		want    montecarlo.Experiment
		wantErr bool
	}{
		{
			name: "should parse image",
			args: &Args{
				Histories: 1,
				Rounds:    17,
				Image:     filepath.Join(filepath.Dir(testFname), "test.png"),
			},
			want: &experiment{
				states: []*state{
					{
						img: [][]bool{
							{true, true, true, true, true, true, false, false, false},
							{true, true, true, true, true, true, false, false, false},
							{true, true, true, true, true, true, false, false, false},
							{true, true, true, true, true, true, false, false, false},
							{true, true, true, true, true, true, false, false, false},
							{true, true, true, true, true, true, false, false, false},
							{true, true, true, true, true, true, false, false, false},
							{true, true, true, true, true, true, false, false, false},
							{true, true, true, true, true, true, false, false, false},
						},
						hit:        0,
						total:      0,
						wantRounds: 17,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", spew.Sdump(got), spew.Sdump(tt.want))
			}
		})
	}
}
