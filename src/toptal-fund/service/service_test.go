package funding

import (
	"sync"
	"testing"
)

var test *testing.T

const WORKERS = 10

func TestInitialBalance(t *testing.T) {
	test = t
	initVal := 10
	s := NewFundServer(initVal)
	assertBalance(s, 10)

	s.Withdraw(1)
	assertBalance(s, 9)
	s.Withdraw(9)
	assertBalance(s, 0)
}

func BenchmarkWithService(b *testing.B) {
	initVal := b.N
	s := NewFundServer(initVal)
	perWorker := initVal / WORKERS

	var wg sync.WaitGroup
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for k := 0; k < perWorker; k++ {
				s.Withdraw(1)
			}
		}()
	}
	wg.Wait()

	assertBalance(s, 0)
}

func assertBalance(s *FundServer, expected int) {
	bal := s.Balance()
	if bal != expected {
		test.Error("Balance should be", expected, "but was", bal)
	}
}
