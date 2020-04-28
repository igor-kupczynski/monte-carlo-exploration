package cointoss

import (
	"math/rand"
	"time"
)

// State represents a state of the coin toss game.
type State struct {
	Capital     int
	Ruined      bool
	LastRoundNo int
}

// ByCapital defines  a sort order.
//
// The state with higher capital is preferred.
// If both states are ruins, the state which stated in the game longer is preferred.
type ByCapital []*State

func (xs ByCapital) Len() int { return len(xs) }

func (xs ByCapital) Less(i, j int) bool {
	if xs[i].Ruined && xs[j].Ruined {
		return xs[i].LastRoundNo < xs[j].LastRoundNo
	}
	return xs[i].Capital < xs[j].Capital
}

func (xs ByCapital) Swap(i, j int) { xs[i], xs[j] = xs[j], xs[i] }

// nextRound advances the game state depending on the coin toss outcome.
//
// Heads - we win, tails - we lose. If the Capital reaches zero we are Ruined. We don't play anymore.
func (s *State) nextRound(heads bool) {
	if s.Ruined {
		return
	}
	s.LastRoundNo += 1
	if heads {
		s.Capital += 1
	} else {
		s.Capital -= 1
	}
	if s.Capital == 0 {
		s.Ruined = true
	}
}

// Play plays the game for given number of rounds from the initial state s
func (s *State) Play(rounds int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for n := 0; n < rounds; n++ {
		s.nextRound(toss(rng))
	}
}

func toss(rng *rand.Rand) bool {
	// TODO: There maybe some problems with math/rand, see https://github.com/golang/go/issues/21835
	return rng.Uint32()&(1<<31) == 0
}
