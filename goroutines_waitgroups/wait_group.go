package grw

import (
	"fmt"
	"sync"
)

func GoroutineWaitGroup(){
	var wg sync.WaitGroup

	var words []string =  greekLetters()

	wg.Add(len(words))

	for index, word := range words {
		go func(){
			defer wg.Done()
			fmt.Printf("index: %d, word: %s", index, word)
			fmt.Println()
		}()
		
	}

	wg.Wait()

	
//	wg.Add(1)

	fmt.Println("Second String")
}