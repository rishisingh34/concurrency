package grw

// ExampleOne demonstrates what happens when a goroutine runs
// without synchronization or waiting for it to complete.
func ExampleOne() {
	// Starts a new goroutine that runs independently.
	go printSomething("This is the first thing")

	// Runs in the current goroutine (current execution thread that called the first go routine) .
	printSomething("This is the second thing")

	// ⚠️ The program may end before the first goroutine runs,
	// so "This is the first thing" might not print.
}