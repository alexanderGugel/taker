package taker

// Task represents a job that can be run.
// You are encouraged to amend this interface accordingly in order to reference
// results of the task.
type Task interface {
	// Run synchronously executes the task.
	// In order to simulate multiple return values while guaranteeing
	// type-safety, amend the interface with additional properties.
	Run() error
}
