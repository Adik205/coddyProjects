package main

import (
	"coddy"
	"fmt"
)

func main() {
	s := "codingisfun"

	res := coddy.PartitionLabels(s)
	fmt.Println("Result: ", res)
}
