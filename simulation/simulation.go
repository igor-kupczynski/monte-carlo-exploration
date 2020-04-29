package simulation

// Simulate runs the Monte Carlo experiment on the collection of samples.
//
// It returns the end states of the simulated histories.
func Simulate(samples []Sample) {
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
}
