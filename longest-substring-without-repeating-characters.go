func lengthOfLongestSubstring(s string) int {
	charInCurrent := make(map[byte]bool, 256)
	longestSoFar := 0
	currentCount := 0
	left := 0
	for i := 0; i < len(s); i++ {
		if charInCurrent[s[i]] {
			for charInCurrent[s[i]] {
				currentCount--
				charInCurrent[s[left]] = false
				left++
			}
		}
		currentCount++
		if currentCount > longestSoFar {
			longestSoFar = currentCount
		}
		charInCurrent[s[i]] = true
	}
	return longestSoFar
}

