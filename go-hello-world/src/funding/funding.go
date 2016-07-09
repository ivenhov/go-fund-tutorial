package funding

import (
	"fmt"
	"sync"
)

const workers = 10

func main() {

	N := 10000

	// Skip N = 1
	if N < workers {
		return
	}

	v := N
	fmt.Printf("iterations bN: %d\n", v)

	// Add as many dollars as we have iterations this run
	//	fund := NewFund(b.N)

	// Casually assume b.N divides cleanly
	dollarsPerFounder := N / workers

	// WaitGroup structs don't need to be initialized
	// (their "zero value" is ready to use).
	// So, we just declare one and then use it.
	var wg sync.WaitGroup

	server := NewFundServer(N)

	// Spawn off the workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < dollarsPerFounder; i++ {
				if server.Balance() <= 10 {
					break
				}
				//				server.Commands <- WithdrawCommand{Amount: 1}
				server.Withdraw(1)
			}
		}()
	}

	// Wait for all the workers to finish
	wg.Wait()

	//	balanceResponseChan := make(chan int)
	//	server.Commands <- BalanceCommand{Response: balanceResponseChan}
	//	balance := <-balanceResponseChan
	balance := server.Balance()

	if balance != 10 {
		panic(fmt.Sprint("Balance unexpected:", balance))
	}
}
