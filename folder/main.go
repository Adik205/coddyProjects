package main

import (
	"coddy"
	"fmt"
)

func main() {
	numRows := 5

	res := coddy.GeneratePascalTriangle(numRows)
	fmt.Println(res)
}
