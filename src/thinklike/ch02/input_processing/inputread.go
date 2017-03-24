package input_processing

import (
	"bufio"
	"os"
)

type StdioScanner struct {
	scan *bufio.Scanner
}

func NewStdioScanner() *StdioScanner {
	er := new(StdioScanner)
	er.scan = bufio.NewScanner(os.Stdin)
	return er
}

func (er *StdioScanner) Scan() string {
	er.scan.Scan()
	return er.scan.Text()
}
