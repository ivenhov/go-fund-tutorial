package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := rnd.Int()
	fmt.Printf("rand: %d", i)
}
