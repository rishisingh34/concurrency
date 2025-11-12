package main

import (
	"fmt"
	grw "github.com/rishisingh34/concurrency_with_go/goroutines_waitgroups"
)

func main() {
	fmt.Println("This repository is for learning \"Concurrency in Go\"")

	var option int
	fmt.Println("enter the concept number to go to: ")
	_, _ = fmt.Scan(&option)

	switch(option){
	case 1:
		// 1. Go keyword, goroutines, waitgroups
		grw.GoroutinesUsage()
	case 2:
		// 2.
	case 3:

	}
}