package funding

import (
	"fmt"
)

type FundServer struct {
	commands chan interface{}
	fund     *Fund
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}

	go server.loop()

	return server
}

func (s *FundServer) Balance() int {
	respChan := make(chan int)
	s.commands <- balanceCommand{response: respChan}
	bal := <-respChan
	return bal
}

func (s *FundServer) Withdraw(amount int) {
	s.commands <- withdrawCommand{amount: amount}
}

func (s *FundServer) loop() {
	for command := range s.commands {
		// handle command
		switch command.(type) {
		case withdrawCommand:
			withdrawal := command.(withdrawCommand)
			s.fund.Withdraw(withdrawal.amount)
		case balanceCommand:
			getBalance := command.(balanceCommand)
			bal := s.fund.Balance()
			getBalance.response <- bal
		default:
			panic(fmt.Sprint("Unrecognized command: %v", command))
		}
	}
}

type withdrawCommand struct {
	amount int
}

type balanceCommand struct {
	response chan int
}
