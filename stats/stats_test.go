package stats

import (
	"reflect"
	"testing"
)

type dataset struct {
	results  []int64
	baseline int64
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		name    string
		dataset dataset
		want    *Summary
	}{
		{
			name: "simplified example",
			dataset: dataset{
				results:  []int64{10},
				baseline: 10,
			},
			want: &Summary{
				Total:    1,
				Baseline: 10,
				Min:      10,
				Max:      10,
				Avg:      10.0,
				Below:    0.0,
				At:       100.0,
				Above:    0.0,
				Percentiles: map[int]int64{
					1:  10,
					5:  10,
					10: 10,
					25: 10,
					50: 10,
					75: 10,
					90: 10,
					95: 10,
					99: 10,
				},
				PercentileDiffs: map[int]float64{
					1:  0.0,
					5:  0.0,
					10: 0.0,
					25: 0.0,
					50: 0.0,
					75: 0.0,
					90: 0.0,
					95: 0.0,
					99: 0.0,
				},
			},
		},
		{
			name:    "increasing capital",
			dataset: incCap(),
			want: &Summary{
				Total:    100,
				Baseline: 50,
				Min:      0,
				Max:      99,
				Avg:      49.5,
				Below:    50.0,
				At:       1.0,
				Above:    49.0,
				Percentiles: map[int]int64{
					1:  1,
					5:  5,
					10: 10,
					25: 25,
					50: 50,
					75: 75,
					90: 90,
					95: 95,
					99: 99,
				},
				PercentileDiffs: map[int]float64{
					1:  -98.0,
					5:  -90.0,
					10: -80.0,
					25: -50.0,
					50: 0.0,
					75: 50.0,
					90: 80.0,
					95: 90.0,
					99: 98.0,
				},
			},
		},
		{
			name:    "30% ruined",
			dataset: ruined30(),
			want: &Summary{
				Total:    100,
				Baseline: 10,
				Min:      0,
				Max:      10,
				Avg:      7,
				Below:    30.0,
				At:       70.0,
				Above:    0.0,
				Percentiles: map[int]int64{
					1:  0,
					5:  0,
					10: 0,
					25: 0,
					50: 10,
					75: 10,
					90: 10,
					95: 10,
					99: 10,
				},
				PercentileDiffs: map[int]float64{
					1:  -100.0,
					5:  -100.0,
					10: -100.0,
					25: -100.0,
					50: 0.0,
					75: 0.0,
					90: 0.0,
					95: 0.0,
					99: 0.0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Describe(tt.dataset.results, tt.dataset.baseline); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Describe() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

// incCap returns 100 samples with increasing capital $0 to $99, the baseline being $50
func incCap() dataset {
	total := 100
	results := make([]int64, total)

	for i := 0; i < total; i++ {
		results[i] = int64(i)
	}

	return dataset{
		results:  results,
		baseline: 50,
	}
}

// ruined30 returns 100 samples: 30% is ruined and 70% has initial capital $10
func ruined30() dataset {
	total := 100
	ruined := 30
	initial := int64(10)

	results := make([]int64, total)

	for i := 0; i < total; i++ {
		var captial int64
		if i >= ruined {
			captial = initial
		}
		results[i] = captial
	}

	return dataset{
		results:  results,
		baseline: initial,
	}
}
