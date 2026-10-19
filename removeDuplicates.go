package coddy

import "slices"

func RemoveDuplicates(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		if !slices.Contains(result, nums[i]) {
			result = append(result, nums[i])
		}
	}
	return result
}
