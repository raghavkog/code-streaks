package codedaily

import (
	"fmt"
)

func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func Main2() {
	nums := []int{1, 2, 2, 3, 4, 3, 5}
	unique := removeDuplicates(nums)
	fmt.Println("After removing duplicates:", unique)
}
