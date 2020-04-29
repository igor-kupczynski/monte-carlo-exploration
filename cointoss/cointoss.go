package cointoss

import (
	"fmt"
	"math/rand"
	"time"
)

// State represents a state of the coin toss game.
type State struct {
	Capital     int
	Ruined      bool
	LastRoundNo int
	TotalRounds int
}

// Args represents the parameters of the coin toss game
type Args struct {
	Rounds         int
	InitialCapital int
}

func ParseArgs(args map[string]interface{}) (*Args, error) {
	var rounds, initialCapital int64
	var ok bool

	var rawRounds interface{}
	if rawRounds, ok = args["rounds"]; !ok {
		return nil, fmt.Errorf("missing required argument 'rounds'")
	}
	if rounds, ok = rawRounds.(int64); !ok {
		return nil, fmt.Errorf("'rounds' needs to be int")
	}

	var rawInitialCapital interface{}
	if rawInitialCapital, ok = args["initial_capital"]; !ok {
		return nil, fmt.Errorf("missing required argument 'initial_capital'")
	}
	if initialCapital, ok = rawInitialCapital.(int64); !ok {
		return nil, fmt.Errorf("'initial_capital' needs to be int")
	}

	return &Args{
		Rounds:         int(rounds),
		InitialCapital: int(initialCapital),
	}, nil
}

// InitState creates an initial state based on the Args
func (a *Args) InitState() *State {
	return &State{
		Capital:     a.InitialCapital,
		Ruined:      false,
		LastRoundNo: 0,
		TotalRounds: a.Rounds,
	}
}

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

// Run plays the game for given number of rounds from the initial state s
func (s *State) Run() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for n := 0; n < s.TotalRounds; n++ {
		s.nextRound(toss(rng))
	}
}

func toss(rng *rand.Rand) bool {
	// TODO: There maybe some problems with math/rand, see https://github.com/golang/go/issues/21835
	return rng.Uint32()&(1<<31) == 0
}
