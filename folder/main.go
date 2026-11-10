package main

import (
	"coddy"
	"fmt"
)

func main() {
	s := "aaaaaaa"
	wordDict := []string{"aaaa", "aaa"}

	res := coddy.WordBreak(s, wordDict)
	fmt.Println(res)
}
