package main

import (
	"fmt"
	"sync"

	"github.com/harturk/go-first/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var waitingGroup sync.WaitGroup
	// cannot use go, to threadling, in every call because the main function
	// will ends before the threads. waitgroup prevents this, by doing
	// something like a sleep but without a fixed time.
	for _, currency := range currencies {
		waitingGroup.Add(1)
		// "Async" function, similar to javacript syntax
		go func(currency string) {
			getCurrencyData(currency)
			waitingGroup.Done()
		}(currency)
	}
	waitingGroup.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	}
}
