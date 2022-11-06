func arrange(str string) string {
	result := str
	chars := []byte(str)
	chars = append(chars, chars...)
	strLen := len(str)
	for i := 1; i < strLen; i++ {
		sub := string(chars[i : i+strLen])
		if result > sub {
			result = sub
		}
	}
	return result
}

func sortChars(str string) string {
	chars := []byte(str)
	sort.Slice(chars, func(i int, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func orderlyQueue(s string, k int) string {
	if k == 1 {
		return arrange(s)
	} else {
		return sortChars(s)
	}
}
