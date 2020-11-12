package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	n  int
)

func main() {
	n = nextInt()
	l := make([]lec, 0)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		l = append(l, lec{a, b})
	}

	//	출발시간을 기준, 오름차순 정렬
	sort.Slice(l, func(i, j int) bool {
		if l[i].s == l[j].s {
			return l[i].t < l[j].t
		}
		return l[i].s < l[j].s
	})

	lpq := new(lHeap) //	강의실 넣기
	heap.Push(lpq, l[0])

	for i := 1; i < n; i++ {
		if lpq.peek() <= l[i].s { // 시작시간이 끝나는 시간이랑 같거나 크면 기존 강의실 이용
			ll := heap.Pop(lpq).(lec)
			ll.t = 0 //	에러처리 코드, 의미없음
		}
		heap.Push(lpq, l[i]) //	새로운 강의실 이용
	}
	fmt.Print(lpq.Len())
}

type lec struct {
	s, t int
}
type lHeap []lec

func (h lHeap) peek() int {
	return h[0].t
}
func (h lHeap) Len() int {
	return len(h)
}
func (h lHeap) Less(i, j int) bool {
	return h[i].t < h[j].t
}
func (h lHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *lHeap) Push(x interface{}) {
	*h = append(*h, x.(lec))
}
func (h *lHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}

//	시간초과코드...

// func 시간초과(){
// 	n = nextInt()
// 	l := make([]lec, 0)
// 	check := make([]bool, n)
// 	count := 0
// 	for i := 0; i < n; i++ {
// 		a, b := nextInt(), nextInt()
// 		l = append(l, lec{a, b})
// 	}
// 	sort.Slice(l, func(i, j int) bool {
// 		return l[i].t < l[j].t
// 	})
// 	for i := 0; i < n; i++ {
// 		if !check[i] {
// 			cur := l[i]
// 			now := cur.t
// 			count++
// 			for j := i + 1; j < n; j++ {
// 				if !check[j] && now == l[j].s {
// 					check[j] = true
// 					now = l[j].s
// 				}
// 			}
// 		}
// 	}
// 	fmt.Print(count)
// }
