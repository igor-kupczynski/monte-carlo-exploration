// montecarlo provides primitives to run Monte Carlo experiments
package montecarlo

import "fmt"

// Experiment represents a Monte Carlo Experiment
type Experiment interface {

	// Samples returns the execution samples we run in the experiment. Multiple samples maybe run in parallel.
	Samples() []Sample

	// Results returns experiment results as something that can be printed
	Results() fmt.Stringer
}

// Sample is a single instance, or "history", of the process being simulated
type Sample interface {

	// Run executes this sample
	//
	// Multiple samples are run in parallel, but a single sample is run only from a single go routine.
	Run()
}

// Simulate runs given Monte Carlo experiment
func Run(e Experiment) fmt.Stringer {

	samples := e.Samples()
	done := make(chan struct{})

	for i := 0; i < len(samples); i++ {
		sample := samples[i]
		go func() {
			sample.Run()
			done <- struct{}{}
		}()
	}

	// Wait until done
	for i := 0; i < len(samples); i++ {
		<-done
	}

	return e.Results()
}
