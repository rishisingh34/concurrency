package main

import (
	"fmt"
	"github.com/rishisingh34/concurrency/mutex"
	grw "github.com/rishisingh34/concurrency/waitgroups"
	"os"
)

func main() {
	fmt.Println("-----------------------------------------------------")
	fmt.Println("This repository is for learning \"Concurrency in Go\"")
	fmt.Println("-----------------------------------------------------")

	for true {
		var option int
		fmt.Println("$ Choose an option (type 0 to terminate the program): ")
		_, _ = fmt.Scan(&option)

		switch option {
		case 0:
			os.Exit(0)
		case 1:
			// 1. Go keyword, goroutines, waitgroups
			grw.GoroutinesUsage()
		case 2:
			// 2. Mutex, Race Condition
			mutex.MutexUsage()
		case 3:

		}
		fmt.Println("*************************")
	}
}
