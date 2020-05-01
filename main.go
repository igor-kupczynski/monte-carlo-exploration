package main

import (
	"flag"
	"fmt"
	"log"

	"monte-carlo-exploration/cointoss"
	"monte-carlo-exploration/montecarlo"

	"github.com/BurntSushi/toml"
)

var configFile = flag.String("conf", "", "Simulation configuration file")

func main() {
	cfg := parseConfig()
	experiment := selectExperiment(cfg)
	results := montecarlo.Run(experiment)
	fmt.Print(results)
}

// config specifies the simulation to run
type config struct {
	Simulation string
	CoinToss   *cointoss.Args
}

func parseConfig() config {
	flag.Parse()
	if *configFile == "" {
		log.Fatalf("Select simulation to run with `-conf path-to-file.toml`")
	}
	fmt.Printf("# Run simulation %s\n", *configFile)

	var cfg config
	if _, err := toml.DecodeFile(*configFile, &cfg); err != nil {
		log.Fatalf("Can't parse %s: %v\n", *configFile, err)
	}
	return cfg
}

func selectExperiment(cfg config) montecarlo.Experiment {
	var experiment montecarlo.Experiment
	switch cfg.Simulation {
	case "cointoss":
		if cfg.CoinToss == nil {
			log.Fatalf("cointoss simulation requires a [cointoss] section\n")
		}
		experiment = cointoss.New(cfg.CoinToss)
		fmt.Printf("## %s\n", cfg.CoinToss)
	default:
		log.Fatalf("Unkown simulation type: %s\n", cfg.Simulation)
	}
	return experiment
}
