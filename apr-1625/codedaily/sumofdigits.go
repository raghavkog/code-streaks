package codedaily

import "fmt"

func Sumofdigits(n int) {
	sum := 0
	num := n
	for n > 0 {
		rem := n % 10
		sum += rem
		n /= 10
	}
	fmt.Println("Sum of digits in ", num, " : ", sum)
}
