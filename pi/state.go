package pi

import (
	"math/rand"
	"time"
)

// state represents a state of the Pi estimation
type state struct {
	// has to be a square
	img [][]bool

	hit   int
	total int

	wantRounds int
}

// Run throws the given number of darts into the circle and saves the score.
func (s *state) Run() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxY := len(s.img)
	maxX := len(s.img[0])
	for n := 0; n < s.wantRounds; n++ {
		x := rng.Intn(maxX)
		y := rng.Intn(maxY)
		if s.img[y][x] {
			s.hit++
		}
		s.total++
	}
}
