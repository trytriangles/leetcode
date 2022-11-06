func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	var minimums = make([]int, len(cost))
	minimums[0] = 0 // Costs 0 to start at 0.
	minimums[1] = 0 // Costs 0 to start at 1.
	// Reaching 2 has the cost of departing from the cheaper starting point.
	minimums[2] = min(cost[0], cost[1])
	for i := 3; i < len(cost); i++ {
		a := cost[i-1] + minimums[i-1]
		b := cost[i-2] + minimums[i-2]
		minimums[i] = min(a, b)
	}
	top := len(cost)
	a := cost[top-1] + minimums[top-1]
	b := cost[top-2] + minimums[top-2]
	return min(a, b)
}
