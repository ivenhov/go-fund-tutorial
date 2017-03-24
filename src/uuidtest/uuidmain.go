package main

import (
	"fmt"
	"uuid"
)

func main() {
	o := uuid.TimeUUID()
	s := o.String()
	//	o :=
	//	fmt.Println("me")
	//	o := "Test"
	fmt.Printf("type: %T\n", o)
	fmt.Println(s)
}
