package funding

import (
	"sync"
	"testing"
)

const numWorkers = 10
const wantedLimit = 10

func BenchmarkTransactStopsAtLimit(b *testing.B) {

	initialBalance := b.N
	if initialBalance < 20 {
		return
	}
	s := NewFundServer(initialBalance)
	var wg sync.WaitGroup

	limitReached := false

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < initialBalance; i++ {
				s.Transact(func(fund *Fund) {
					if fund.Balance() <= wantedLimit {
						limitReached = true
						return
					}
					fund.Withdraw(1)
				})
			}
		}()
	}

	wg.Wait()

	bal := s.Balance()
	if bal != wantedLimit {
		b.Error("Wanted limit", wantedLimit, "but got", bal)
	}

}
