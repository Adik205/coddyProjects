package main

import (
	"coddy"
	"fmt"
)

func main() {
	s := "aabb"
	t := "abbb"

	res := coddy.IsAnagram(s, t)
	fmt.Println(res)
}
