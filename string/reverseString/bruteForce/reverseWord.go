package main

import (
	"strings"
)

func main() {
	str := "  hello world  "
	_ = reverseWords(str)

}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	var stack []string
	var output string

	for ind := range words {
		if words[ind] == "" {
			continue
		}
		stack = append(stack, words[ind])
	}

	for ind := len(stack) - 1; ind > -1; ind-- {
		if ind == 0 {
			output += stack[ind]
			return output
		}
		output += stack[ind] + " "
	}
	return output
}
