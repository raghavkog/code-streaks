package codedaily

import (
	"fmt"
	"strings"
)

func Reversestring(s string) {
	var revstr = " "
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	revstr = strings.Join(words, " ")

	fmt.Println("Reverse of the string", s, " is :", revstr)
}
