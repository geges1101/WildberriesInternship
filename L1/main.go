package main

import (
	"fmt"
	"strings"
)

func validate(s string) bool {
	m := make(map[string]int)

	for i := 0; i != len(s); i++ {
		c := strings.ToLower(string(s[i]))
		if m[c] == 0 {
			m[c]++
			continue
		}
		if m[c] > 0 {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(validate(s))
}
