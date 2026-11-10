import "slices"

func canPartition(nums []int) (bool) {
    var sum int
	for _, num := range nums {
		sum += num
	}

	if sum % 2 == 1 {
		return false
	}
	isReach := make([]bool, sum/2+1) // true false
	usedPos := make([][]int, sum/2+1) // [] 
	isReach[0] = true
	for i := 1; i < len(isReach); i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 && isReach[i-nums[j]] == true && !slices.Contains(usedPos[i-nums[j]], j) {
				usedPos[i] = append(usedPos[i], usedPos[i-nums[j]]...)
				usedPos[i] = append(usedPos[i], j)
				isReach[i] = true
				break
			}
		}
	}

	if isReach[sum/2] {
		return true
	}
	return false
}