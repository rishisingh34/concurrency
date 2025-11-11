// Package goroutines_go_waitgroups : this provides examples of using go keyword, goroutines, waitgroups
package grw

import (
	"fmt"
)

func GoRoutinesUsage(){
	var option int
	_, _ = fmt.Scan(&option)

	switch(option){
	case 1:
		// 1. go routine dies before it could do its work
		ExampleOne()
	case 2:
		// 2. time.Sleep() to wait for go routine to fininsh ( it may or may not work )
		ExampleTwo()
	case 3:

	}

}