package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	n  int
	m  [500][2]int
	dp [500][500]int
)

func main() {
	n = nextInt()
	for i := 0; i < n; i++ {
		m[i][0], m[i][1] = nextInt(), nextInt()
	}
	//	dp[i][j] -> i번과 j번의 최소 곱셈 횟수
	//	if n -> 3... [0 1] [1 2] [0 2]
	//	if n -> 4... [0 1] [1 2] [2 3] [0 2] [1 3] [0 3]
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			dp[j][j+i] = math.MaxInt64
			for k := 0; k < i; k++ {
				c := dp[j][j+k] + dp[j+k+1][j+i] + m[j][0]*m[j+k][1]*m[j+i][1]
				dp[j][j+i] = minInt(dp[j][j+i], c)
			}
		}
	}
	fmt.Print(dp[0][n-1])
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
