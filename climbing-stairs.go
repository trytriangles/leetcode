// # Problem statement
//
// You are climbing a staircase. It takes n steps
// to reach the top. Each time you can either climb
// 1 or 2 steps. In how many distinct ways can you
// climb to the top?
//
// # Constraints
// 
// 1 <= n <= 45
//
// # Thoughts
//
// This is a second-order recurrence relation. We can
// reach step n from either step n-1 or step n-2;
// therefore the result for step in is the sum of the
// results of steps n-1 and n-2.
//
// The base cases are steps 1 and 2, which have only 1
// and 2 ways of reaching them respectively, therefore
// f(x) = f(x-1) + f(x-2).
//
// A recursive solution would run in exponential time,
// a memoized recursive solution would run in linear
// time but require linear space. The iterative
// solution below runs in linear time and constant
// space.

func climbStairs(n int) int {
    a := 0
    b := 1
    for i := 0; i < n; i++ {
        a, b = b, a + b
    }
    return b
}

