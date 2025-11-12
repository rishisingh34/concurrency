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
		go printSomething(fmt.Sprintf("%d: %s", index, word),&wg)
	}

	wg.Wait()

	
	wg.Add(1)

	printSomething("Second String", &wg)
}