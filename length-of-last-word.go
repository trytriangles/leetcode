func lengthOfLastWord(s string) int {
	wordLength := 0
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == ' ' {
			if wordLength > 0 {
				return wordLength
			}
			continue
		}
		wordLength++
		if i == 0 {
			return wordLength
		}
	}
	return -1
}
