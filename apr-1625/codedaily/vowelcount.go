package codedaily

import (
	"fmt"
	"strings"
)

func Countowels(s string) {
	vowels := "aeiou"
	count := 0
	s = strings.ToLower(s)

	for _, ch := range s {
		if strings.ContainsRune(vowels, ch) {
			count++
		}
	}
	fmt.Println("No. of vowels in ", s, " are: ", count)
}
