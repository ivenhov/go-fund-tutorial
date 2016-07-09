package funding

import (
	"fmt"
)

type FundServer struct {
	commands chan TransactionCommand
	fund     *Fund
}

// Add a new command type with a callback and a semaphore channel
type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

// Typedef the callback for readability
type Transactor func(fund *Fund)

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		// make() creates builtins like channels, maps, and slices
		commands: make(chan TransactionCommand),
		fund:     NewFund(initialBalance),
	}

	fmt.Printf("Created fund server with balance %d\n", initialBalance)
	// Spawn off the server's main loop immediately
	go server.loop()
	return server
}

func (s *FundServer) Balance() int {
	var balance int
	s.Transact(func(fund *Fund) {
		balance = fund.Balance()
	})
	return balance
}

func (s *FundServer) Withdraw(amount int) {
	s.Transact(func(fund *Fund) {
		fund.Withdraw(amount)
	})
}

func (s *FundServer) Transact(transactor Transactor) {
	command := TransactionCommand{
		Transactor: transactor,
		Done:       make(chan bool),
	}
	s.commands <- command
	<-command.Done
}

func (s *FundServer) loop() {
	for transaction := range s.commands {
		transaction.Transactor(s.fund)
		transaction.Done <- true
	}
}
