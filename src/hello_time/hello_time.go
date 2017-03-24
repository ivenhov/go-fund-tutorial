package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("Hello World! Now is: %s\n", t)
	utc := t.UTC()
	fmt.Printf("UTC time now: %s\n", utc)

	var d time.Duration
	d = time.Since(t)
	fmt.Printf("Duration since: %s\n", d)
}
