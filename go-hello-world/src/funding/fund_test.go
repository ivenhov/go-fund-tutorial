package funding

import (
	"fmt"
	"sync"
	"testing"
)

const WORKERS = 10

func _BenchmarkFund(b *testing.B) {

	fund := NewFund(b.N)

	for i := 0; i < b.N; i++ {
		fund.Withdraw(1)
	}

	if fund.Balance() != 0 {
		b.Error("Balance wasn't zero:", fund.Balance())
	}
}

func _BenchmarkWithdrawals(b *testing.B) {
	// Skip N = 1
	if b.N < WORKERS {
		return
	}

	v := b.N
	fmt.Printf("iterations bN: %d\n", v)

	// Add as many dollars as we have iterations this run
	fund := NewFund(b.N)

	// Casually assume b.N divides cleanly
	dollarsPerFounder := b.N / WORKERS

	// WaitGroup structs don't need to be initialized
	// (their "zero value" is ready to use).
	// So, we just declare one and then use it.
	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		// Let the waitgroup know we're adding a goroutine
		wg.Add(1)

		// Spawn off a founder worker, as a closure
		go func() {
			// Mark this worker done when the function finishes
			defer wg.Done()

			for i := 0; i < dollarsPerFounder; i++ {
				fund.Withdraw(1)
			}

		}() // Remember to call the closure!
	}

	// Wait for all the workers to finish
	wg.Wait()

	if fund.Balance() != 0 {
		b.Error("Balance wasn't zero:", fund.Balance())
	}
}

func _BenchmarkWithdrawalsServer(b *testing.B) {
	// Skip N = 1
	if b.N < WORKERS {
		return
	}

	v := b.N
	fmt.Printf("iterations bN: %d\n", v)

	// Casually assume b.N divides cleanly
	dollarsPerFounder := b.N / WORKERS

	// WaitGroup structs don't need to be initialized
	// (their "zero value" is ready to use).
	// So, we just declare one and then use it.
	var wg sync.WaitGroup

	server := NewFundServer(b.N)

	// Spawn off the workers
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < dollarsPerFounder; i++ {
				if server.Balance() <= 10 {
					break
				}
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
		b.Error("Balance unexpected:", balance)
	}
}

func BenchmarkWithdrawalsServerWithTxs(b *testing.B) {
	// Skip N = 1
	if b.N < WORKERS {
		return
	}

	v := b.N
	fmt.Printf("iterations bN: %d\n", v)

	// Casually assume b.N divides cleanly
	dollarsPerFounder := b.N / WORKERS

	// WaitGroup structs don't need to be initialized
	// (their "zero value" is ready to use).
	// So, we just declare one and then use it.
	var wg sync.WaitGroup

	server := NewFundServer(b.N)

	// Spawn off the workers
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pizzaTime := false
			for i := 0; i < dollarsPerFounder; i++ {
				server.Transact(func(fund *Fund) {
					if fund.Balance() <= 10 {
						// Set it in the outside scope
						pizzaTime = true
						return
					}
					fund.Withdraw(1)
				})
				if pizzaTime {
					break
				}
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
		b.Error("Balance unexpected:", balance)
	}
}
