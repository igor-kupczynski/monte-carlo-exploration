package main

import (
	"flag"
	"fmt"
	"log"
	"path"

	"github.com/igor-kupczynski/monte-carlo-exploration/cointoss"
	"github.com/igor-kupczynski/monte-carlo-exploration/montecarlo"
	"github.com/igor-kupczynski/monte-carlo-exploration/pi"

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
	Pi         *pi.Args
}

func parseConfig() config {
	flag.Parse()
	if *configFile == "" {
		log.Fatalf("Select simulation to run with `--conf path-to-file.toml`")
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
			log.Fatalf("cointoss simulation requires a [cointoss] section")
		}
		experiment = cointoss.New(cfg.CoinToss)
		fmt.Printf("## %s\n", cfg.CoinToss)
	case "pi":
		if cfg.Pi == nil {
			log.Fatalf("pi estimation requires a [pi] section")
		}
		// make the image file relative to the config file
		cfg.Pi.Image = path.Join(path.Dir(*configFile), cfg.Pi.Image)
		var err error
		if experiment, err = pi.New(cfg.Pi); err != nil {
			log.Fatalf("pi estimation failed: %v", err)
		}
		fmt.Printf("## %s\n", cfg.Pi)
	default:
		log.Fatalf("Unkown simulation type: %s\n", cfg.Simulation)
	}
	return experiment
}
