package main

import "fmt"

func main() {
	str := "  hello world  "
	res := reverseWords(str)
	fmt.Println(res)

}

func reverseWords(s string) string {
	out, temp := "", ""
	for i := 0; i <= len(s)-1; i++ {
		if s[i] != ' ' {
			temp += string(s[i])
		} else if s[i] == ' ' {
			if out != "" && temp != "" {
				out = temp + " " + out
				temp = ""
			} else {
				out += temp
				temp = ""
			}
		}

	}
	if out != "" && temp != "" {
		out = temp + " " + out
		temp = ""
	} else {
		out += temp
		temp = ""
	}
	return out
}
