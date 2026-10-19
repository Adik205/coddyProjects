package coddy

import (
	"maps"
)

func IsAnagram(s string, t string) bool {
	a := make(map[byte]int)
	b := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		a[s[i]]++
		b[t[i]]++
	}

	if maps.Equal(a, b) {
		return true
	}
	return false
}
