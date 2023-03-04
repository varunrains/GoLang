package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source  string
	balance float64
}

func main() {
	//variable for bank balance
	bankBalance := 0.00
	//print out starting values
	fmt.Printf("Starting balance is %f", bankBalance)
	//define weekly revenue

	weeklyEarnings := []Income{
		{
			Source:  "Main Job",
			balance: 10,
		},
		{
			Source:  "Part time Job",
			balance: 500,
		},
		{
			Source:  "Wed Job",
			balance: 100,
		},
		{
			Source:  "Side Job",
			balance: 50,
		},
	}

	//loop through 52 weeks and print out how much is made; keep running total
	var mu sync.Mutex
	for i := 1; i <= 52; i++ {
		wg.Add(len(weeklyEarnings))
		go func(ind int, wg *sync.WaitGroup) {
			for _, earning := range weeklyEarnings {
				fmt.Printf("%d week earning from %s is %f.00", ind, earning.Source, earning.balance)
				fmt.Println()
				mu.Lock()
				bankBalance += earning.balance
				mu.Unlock()
				defer wg.Done()
			}

		}(i, &wg)

	}

	//print out final balance
	wg.Wait()
	fmt.Println("Total yearly earning: ", bankBalance)

}
