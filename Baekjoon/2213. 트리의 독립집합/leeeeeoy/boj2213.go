package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	tree   [][]int
	w      []int
	n      int
	dp     [][2]int
	check  []bool
	answer = make([]int, 0)
)

func main() {
	n = nextInt()
	tree = make([][]int, n+1)
	w = make([]int, n+1)
	dp = make([][2]int, n+1)
	check = make([]bool, n+1)
	for i := 1; i <= n; i++ {
		w[i] = nextInt()
		dp[i][0] = -1
		dp[i][1] = -1
	}
	for i := 1; i < n; i++ {
		a, b := nextInt(), nextInt()
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}
	//	독립집합 합 구하기
	ret := maxInt(findDP(1, 0, -1), findDP(1, 1, -1))
	fmt.Println(ret)
	//	탐색
	dfs(1, 0)
	sort.Ints(answer)
	for _, number := range answer {
		fmt.Print(number, " ")
	}
}

func dfs(now, parent int) {
	//	현재 노드를 포함한 게 포함하지 않은 것보다 크고
	//	이전 방문 정점이 독립 집합에 포함되지 않을 때
	//	독립집합 포함 노드
	if dp[now][1] > dp[now][0] && check[parent] == false {
		answer = append(answer, now)
		check[now] = true
	}
	for _, next := range tree[now] {
		if next != parent {
			dfs(next, now)
		}
	}
}

//	독립 집합의 합 구하기
//	dp[i][0] i를 정점으로 하는 서브트리에서 i 미포함
//	dp[i][1] i를 정점으로 하는 서브트리에서 i 포함
func findDP(now, isInclude, parent int) int {
	cur := dp[now][isInclude]
	if cur != -1 { //	이전에 찾았으면 종료
		return cur
	}
	if isInclude == 1 { //	현재 노드를 포함시키면
		cur = w[now]
		for _, next := range tree[now] {
			if next != parent {
				//	다음노드는 포함시키지 않음
				cur += findDP(next, 0, now)
			}
		}
	} else { //	현재 노드 미포함시
		cur = 0
		for _, next := range tree[now] {
			if next != parent {
				//	포함시킨 것과 안시킨 것 중 큰값
				cur += maxInt(findDP(next, 1, now), findDP(next, 0, now))
			}
		}
	}
	dp[now][isInclude] = cur
	return cur
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
