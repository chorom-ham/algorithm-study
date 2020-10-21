package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	n   int
	w   [16][16]int
	dp  [16][1 << 16]int //	dp[i][j] -> 현재 위치:i, 방문한 도시: j
	max = math.MaxInt32
)

//	비트마스크 이용
//	외판원 순회(TSP)
func main() {
	n = nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			w[i][j] = nextInt()
		}
	}
	//	1번부터 출발
	fmt.Print(solve2(0, 1))
}
func solve2(now, v int) int {
	if v == (1<<n - 1) {
		if w[now][0] > 0 {
			return w[now][0]
		}
		return max
	}
	//	max 아닐 때 추가
	if dp[now][v] != max && dp[now][v] > 0 {
		return dp[now][v]
	}
	dp[now][v] = max
	for i := 0; i < n; i++ {
		next := v | (1 << i)
		if (v&(1<<i)) == 0 && w[now][i] != 0 {
			cur := solve2(i, next) + w[now][i]
			if dp[now][v] > cur {
				dp[now][v] = cur
			}
		}
	}
	return dp[now][v]
}

//	시간초과...
//	now: 현재 위치, v: 방문한 도시 집합
func solve(now, v int) int {
	v |= (1 << now)
	if v == (1<<n)-1 { //	모든 도시 방문
		if w[now][0] > 0 { //	돌아가는 길이 있음
			return w[now][0]
		}
		return max //	돌아가는 길이 없음
	}
	if dp[now][v] > 0 {
		return dp[now][v]
	}
	dp[now][v] = max
	for i := 0; i < n; i++ {
		//	방문x, 가는 경로가 존재
		if i != now && (v&(1<<i)) == 0 && w[now][i] > 0 {
			cur := solve(i, v) + w[now][i]
			//	최소 비용 갱신
			if dp[now][v] > cur {
				dp[now][v] = cur
			}
		}
	}
	return dp[now][v]
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
