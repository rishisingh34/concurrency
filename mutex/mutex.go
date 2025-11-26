package mutex

import (
	"fmt"
)

func MutexUsage(){
	for true {
		var option int
		fmt.Println("Inside Ch2. Mutex Usage")
		fmt.Println("**** Choose an option number ****")
		if _, err := fmt.Scan(&option); err != nil { return }
		switch(option) {
		case 0:
			return
		case 1:
			// 1. Race Condition
			RaceCondition()
		case 2:
			// 2. Sync Mutex to prevent race condition
			SyncMutex()
		case 3: 
			// complex example for mutex 
			ComplexMutexExample()
		}

		fmt.Println("*************************")
	}
}
