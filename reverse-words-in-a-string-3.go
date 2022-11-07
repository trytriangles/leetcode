func reverseWords(s string) string {
	chars := make([]byte, len(s))
	lastSpace := -1
	for i := 0; i < len(s); i++ {
		chars[i] = ' '
		if s[i] != ' ' {
			continue
		}
		reversedWord := reverseWord(s[lastSpace+1 : i])
		copy(chars[lastSpace+1:i], reversedWord)
		lastSpace = i
	}
	reversedWord := reverseWord(s[lastSpace+1:])
	copy(chars[lastSpace+1:len(s)], reversedWord)
	return string(chars)
}

func reverseWord(s string) []byte {
	b := []byte(s)
	max := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		op := max - i
		b[i], b[op] = b[op], b[i]
	}
	return b
}

