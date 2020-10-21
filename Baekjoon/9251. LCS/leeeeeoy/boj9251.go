package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br *bufio.Reader = bufio.NewReader(os.Stdin)
	bw *bufio.Writer = bufio.NewWriter(os.Stdout)
	dp [1001][1001]int
)

//LCS(최장공통부분수열) --> 이차원 배열 이용, 공식에 코드 대입...
func main() {
	defer bw.Flush()
	var s1, s2 string
	fmt.Fscan(br, &s1, &s2)

	sl1, sl2 := len(s1), len(s2)
	ar1, ar2 := []rune(s1), []rune(s2)

	//	---공식 이용---
	//	1. 문자가 같으면 좌측위  + 1
	//	2. 문자가 다르면 좌측 or 위의 숫자 중 큰값, 동일할 경우엔 1번처럼 처리
	for i := 1; i <= sl1; i++ {
		for j := 1; j <= sl2; j++ {
			if ar1[i-1] == ar2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	fmt.Fprintln(bw, dp[sl1][sl2])
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
