package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

	"monte-carlo-exploration/cointoss"
	"monte-carlo-exploration/simulation"

	"github.com/BurntSushi/toml"
)

var configFile = flag.String("conf", "examples/cointoss.toml", "Simulation configuration file")

// config specifies the simulation to run
type config struct {
	Simulation string
	Histories  int
	Args       map[string]interface{}
}

func main() {
	flag.Parse()
	fmt.Printf("# Run simulation %s\n", *configFile)

	var cfg config
	if _, err := toml.DecodeFile(*configFile, &cfg); err != nil {
		log.Fatalf("Can't parse %s: %v\n", *configFile, err)
	}
	histories := cfg.Histories

	var args *cointoss.Args
	var err error
	if args, err = cointoss.ParseArgs(cfg.Args); err != nil {
		log.Fatalf("Can't parse %s: %v\n", *configFile, err)
	}

	states := simulation.Simulate(func() *cointoss.State { return args.InitState() }, histories)

	// Sort the results by capital
	sort.Sort(cointoss.ByCapital(states))

	firstNotRuined := sort.Search(histories, func(i int) bool { return !states[i].Ruined })
	procRuined := 100 * float64(firstNotRuined) / float64(histories)
	firstSameCapital := sort.Search(histories, func(i int) bool { return states[i].Capital >= args.InitialCapital })
	procLessCapital := 100 * float64(firstSameCapital) / float64(histories)
	firstMoreCapital := sort.Search(histories, func(i int) bool { return states[i].Capital > args.InitialCapital })
	procMoreCapital := 100 * float64(histories-firstMoreCapital) / float64(histories)

	wantPercentiles := []int{1, 5, 10, 50, 90, 95, 99}
	percentiles := make(map[int]int)
	for _, p := range wantPercentiles {
		percentiles[p] = states[p*histories/100].Capital
	}

	fmt.Printf("## Simulating %d executions of %d round coin toss with starting capital of $%d\n",
		histories, args.Rounds, args.InitialCapital)
	fmt.Printf("Ruined: %f%% (%d / %d)\n", procRuined, firstNotRuined, histories)
	fmt.Printf("Less capital: %f%% (%d / %d)\n", procLessCapital, firstSameCapital, histories)
	fmt.Printf("More capital: %f%% (%d / %d)\n", procMoreCapital, histories-firstMoreCapital, histories)

	for _, p := range wantPercentiles {
		fmt.Printf("p%02d $%d\n", p, percentiles[p])
	}
}
