package mutex

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

type Income struct {
	Source string 
	Amount int
}


// print out starting values
// define weekly revenue
// loop through 52 weeks and print out how much is made; keep a running total
// print out final balance

func ComplexMutexExample() {
	// var for bank balance
	var bankBalance int 
	var balance sync.Mutex
	
	fmt.Printf("Initial Account Balance: $%d.00", bankBalance) 
	fmt.Println()
	
	// weekly revenue 
	incomes := []Income{
		{Source: "Main Job", Amount: 500}, 
		{Source: "Gifts", Amount: 10},
		{Source: "Part Time Job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}
	
	wg2.Add(len(incomes)) 
	
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg2.Done() 
			for week := 1; week <= 52; week++ {
				
				// locking the balance for safe access to critical section ( prevents race cond. )
				balance.Lock()
				temp := bankBalance
				temp += income.Amount 
				bankBalance = temp 
				balance.Unlock()
				// balance unlocked for other goroutines to access 
				
				fmt.Printf("On Week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source) 
			}
		}(i, income)
	}
	
	wg2.Wait()
	fmt.Printf("Final Bank Balance: $%d.00\n", bankBalance) 
}