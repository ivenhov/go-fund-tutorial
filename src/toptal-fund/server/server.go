package funding

import (
	"fmt"
)

type FundServer struct {
	Commands chan interface{}
	fund     *Fund
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		Commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}

	go server.loop()

	return server
}

func (s *FundServer) loop() {
	for command := range s.Commands {
		// handle command
		switch command.(type) {
		case WithdrawCommand:
			withdrawal := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawal.Amount)
		case BalanceCommand:
			getBalance := command.(BalanceCommand)
			bal := s.fund.Balance()
			getBalance.Response <- bal
		default:
			panic(fmt.Sprint("Unrecognized command: %v", command))
		}
	}
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}
