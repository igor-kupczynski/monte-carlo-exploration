package main

import (
	"flag"
	"fmt"
	"log"

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

	fmt.Printf("## Simulating %d executions of %d round coin toss with starting capital of $%d\n",
		histories, args.Rounds, args.InitialCapital)

	// Generate initial state of coin toss samples
	tosses := make([]*cointoss.State, histories)
	for i := range tosses {
		tosses[i] = args.InitState()
	}

	// Run the simulation
	experiment := cointoss.Experiment(tosses)
	simulation.Simulate(experiment.Samples())
	experiment.PrintReport(args.InitialCapital)
}
