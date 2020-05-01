package cointoss

import (
	"math/rand"
	"time"
)

// state represents a state of the coin toss game.
type state struct {
	capital     int
	ruined      bool
	lastRoundNo int

	wantRounds     int
	initialCapital int
}

// nextRound advances the game state depending on the coin toss outcome.
//
// Heads - we win, tails - we lose. If the capital reaches zero we are ruined. We don't play anymore.
func (s *state) nextRound(heads bool) {
	if s.ruined {
		return
	}
	s.lastRoundNo += 1
	if heads {
		s.capital += 1
	} else {
		s.capital -= 1
	}
	if s.capital == 0 {
		s.ruined = true
	}
}

// Run plays the game for given number of rounds from the initial state s
func (s *state) Run() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for n := 0; n < s.wantRounds; n++ {
		s.nextRound(toss(rng))
	}
}

func toss(rng *rand.Rand) bool {
	// TODO: There maybe some problems with math/rand, see https://github.com/golang/go/issues/21835
	return rng.Uint32()&(1<<31) == 0
}

// stateSlice provides sort ordering for the slices of *state
type stateSlice []*state

func (s stateSlice) Len() int { return len(s) }

func (s stateSlice) Less(i, j int) bool {
	if s[i].ruined && s[j].ruined {
		// The earlier we are ruined the worse we do
		return s[i].lastRoundNo < s[j].lastRoundNo
	}
	return s[i].capital < s[j].capital
}

func (s stateSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
