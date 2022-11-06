func firstBadVersion(n int) int {
	midpoint := 0
	floor := 1
	ceiling := n
	currentGuess := 1
	for floor <= ceiling {
		midpoint = floor + (ceiling-floor)/2
		if isBadVersion(midpoint) {
			currentGuess = midpoint
			ceiling = midpoint - 1
		} else {
			floor = midpoint + 1
		}
	}
	return currentGuess
}