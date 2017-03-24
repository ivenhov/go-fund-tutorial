package funding

import (
	"sync"
	"testing"
)

const WORKERS = 10

func BenchmarkFundConcurr(b *testing.B) {

	perWorker := b.N / WORKERS

	fund := NewFund(b.N)
	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < perWorker; i++ {
				fund.Withdraw(1)
			}
		}()
	}

	wg.Wait()

	if fund.Balance() != 0 {
		b.Error("Result not zero:", fund.Balance())
	}
}
