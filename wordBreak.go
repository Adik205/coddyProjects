package coddy

func WordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	for i := 1; i < len(dp); i++ {
		copyS := s[:i]
		for j := i - 1; j > 0; j-- {
			if dp[j] == true {
				processedS := copyS[j:]
				for k := 0; k < len(wordDict); k++ {
					if processedS == wordDict[k] {
						dp[i] = true
						break
					}
				}
			}
		}
		for k := 0; k < len(wordDict); k++ {
			if copyS == wordDict[k] {
				dp[i] = true
				break
			}
		}
	}

	if dp[len(s)] == true {
		return true
	} else {
		return false
	}

}
