package grw

import (
	"fmt"
	"sync"
)
func printSomething(s string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(s)
}
func greekLetters() []string{
	return []string{
		"aplha",
		"beta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}
}