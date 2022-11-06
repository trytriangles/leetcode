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
	y := 0                     // Cost of reaching pos 1.
	z := min(cost[0], cost[1]) // Cost of reaching pos 2.
	for i := 3; i < len(cost); i++ {
		a := cost[i-1] + z
		b := cost[i-2] + y
		z, y = min(a, b), z
	}
	top := len(cost)
	return min(cost[top-1]+z, cost[top-2]+y)
}