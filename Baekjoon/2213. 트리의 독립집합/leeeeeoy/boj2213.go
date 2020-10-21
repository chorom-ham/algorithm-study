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

//다시 풀어볼것
func main() {
	n = nextInt()
	tree = make([][]int, n+1)
	w = make([]int, n+1)
	dp = make([][2]int, n+1)
	check = make([]bool, n+1)
	for i := 1; i <= n; i++ {
		w[i] = nextInt()
	}
	for i := 1; i < n; i++ {
		a, b := nextInt(), nextInt()
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}
	ret := maxInt(findDP(1, 0, -1), findDP(1, 1, -1))
	fmt.Println(ret)
	dfs(1, 0)
	sort.Ints(answer)
	for i := 0; i < len(answer); i++ {
		fmt.Print(answer[i], " ")
	}
}
func dfs(now, parent int) {
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
func findDP(now, isInclude, parent int) int {
	cur := dp[now][isInclude]
	if cur != 0 {
		return cur
	}
	if isInclude == 1 {
		cur = w[now]
		for _, next := range tree[now] {
			if next != parent {
				cur += findDP(next, 0, now)
			}
		}
	} else {
		cur = 0
		for _, next := range tree[now] {
			if next != parent {
				cur += maxInt(findDP(next, 1, now), findDP(next, 0, now))
			}
		}
	}
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
