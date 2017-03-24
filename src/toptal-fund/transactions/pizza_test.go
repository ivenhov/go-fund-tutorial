package funding

import (
	"sync"
	"testing"
)

const workers = 30
const limit = 10

func BenchmarkStopsAtLimit(b *testing.B) {
	initFund := 1000

	s := NewFundServer(initFund)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for l := 0; l < initFund; l++ {
				if s.Balance() <= limit {
					break
				}
				s.Withdraw(1)
			}
		}()
	}
	wg.Wait()

	bal := s.Balance()
	if bal != 10 {
		b.Error("Expected balance ", limit, ", was:", bal)
	}
}
