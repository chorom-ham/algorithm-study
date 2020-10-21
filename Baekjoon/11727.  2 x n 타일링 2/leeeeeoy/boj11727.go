package main

import "fmt"

//점화식 이용
func main() {
	var n int
	fmt.Scan(&n)
	dp := [1001]int{}
	dp[1], dp[2] = 1, 3
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + 2*dp[i-2]) % 10007 // 문제 조건 조심
	}
	fmt.Print(dp[n])
}
