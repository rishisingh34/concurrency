package grw

import "fmt"

// SimpleGoroutine demonstrates what happens when a goroutine runs
// without synchronization or waiting for it to complete.
func SimpleGoroutine() {
	// Starts a new goroutine that runs independently.
	go fmt.Println("This is the first thing")

	// Runs in the current goroutine (current execution thread that called the first go routine) .
	fmt.Println("This is the second thing")

	// ⚠️ The program may end before the first goroutine runs,
	// so "This is the first thing" might not print.
}