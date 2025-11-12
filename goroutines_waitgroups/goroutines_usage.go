// Package goroutines_go_waitgroups : this provides examples of using go keyword, goroutines, waitgroups
package grw

import (
	"fmt"
)

// GoroutinesUsage is chapter one of concurrency with Go.
func GoroutinesUsage(){
	var option int
	fmt.Println("Inside the GoRoutines Usage function")
	fmt.Println("Enter the option from 1..3: ")
	_, _ = fmt.Scan(&option)

	switch(option){
	case 1:
		// 1. go routine dies before it could do its work
		SimpleGoroutine()
	case 2:
		// 2. time.Sleep() to wait for go routine to fininsh ( it may or may not work )
		GoroutineWithSleep()
	case 3:
		// 3. Waitgroup
		GoroutineWaitGroup()
	default:
		fmt.Println("Wrong option")
	}

}