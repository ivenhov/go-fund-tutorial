package funding

import ()

type FundServer struct {
	commands chan transactionCommand
	fund     *Fund
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan transactionCommand),
		fund:     NewFund(initialBalance),
	}

	go server.loop()

	return server
}

func (s *FundServer) Transact(t Transactor) {
	cmd := transactionCommand{Transactor: t, Done: make(chan bool)}
	s.commands <- cmd
	<-cmd.Done
}

func (s *FundServer) Balance() int {
	var balance int
	s.Transact(func(f *Fund) {
		balance = f.Balance()
	})
	return balance
}

func (s *FundServer) Withdraw(amount int) {
	s.Transact(func(f *Fund) {
		f.Withdraw(amount)
	})
}

func (s *FundServer) loop() {
	for tx := range s.commands {
		// handle command
		tx.Transactor(s.fund)
		tx.Done <- true
	}
}

type Transactor func(fund *Fund)

type transactionCommand struct {
	Transactor Transactor
	Done       chan bool
}
