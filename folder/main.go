package main

import (
	"coddy"
	"fmt"
)

func main() {
	nums := []int{1, 1, 2, 2, 2, 3, 4}

	res := coddy.RemoveDuplicates(nums)
	fmt.Println(res)
}
