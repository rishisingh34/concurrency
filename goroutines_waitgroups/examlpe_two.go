package grw

import "time"

// ExampleTwo demonstrates using time.Sleep to give a goroutine
// time to finish before the main function exits.
func ExampleTwo() {
	// Starts a new goroutine that runs concurrently.
	go printSomething("This is the first thing")

	// Pause the current goroutine (usually main) for 1 second
	// to let the other goroutine finish its work.
	time.Sleep(1 * time.Second)

	// Runs in the current goroutine.
	printSomething("This is the second thing")

	// ⚠️ WARNING:
	// Sleeping is not a reliable way to synchronize goroutines.
	// The first goroutine might still not complete in time if the system
	// is slow or the task takes longer than 1 second.
	// Always use sync.WaitGroup or channels to coordinate goroutines properly.
}
