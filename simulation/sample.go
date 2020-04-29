package simulation

// Sample is a single instance, or "history", of the process being simulated
type Sample interface {

	// Run executes this sample
	Run()
}
