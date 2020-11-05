package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	sc               = bufio.NewScanner(os.Stdin)
	n, m, start, end int
	h                = [1000]posHeap{}
	dis              = [1000]int{}
	route            = [1000]int{}
	routes           = make([]int, 0)
)

//	다익스트라
//	런타임에러...다시
func main() {
	n, m = nextInt(), nextInt()
	for i := 0; i < m; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		h[a] = append(h[a], pos{b, c})
	}
	start, end = nextInt(), nextInt()
	find(start)

	fmt.Println(dis[end])
	fmt.Println(len(routes))
	for i := len(routes) - 1; i >= 0; i-- {
		fmt.Print(routes[i], " ")
	}
}
func find(x int) {
	for i := 0; i < len(dis); i++ {
		dis[i] = 1000001
	}

	pq := new(posHeap)
	heap.Push(pq, pos{x, 0})
	dis[0] = 0
	route[0] = 0

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(pos)
		node := cur.node
		cost := cur.cost

		for _, curPos := range h[node] {
			nextNode := curPos.node
			nextCost := curPos.cost + cost

			if dis[nextNode] > nextCost {
				dis[nextNode] = nextCost
				heap.Push(pq, pos{nextNode, nextCost})
				route[nextNode] = node
			}
		}
	}

	node := end
	for node != 0 {
		routes = append(routes, node)
		node = route[node]
	}
}

type pos struct {
	node, cost int
}
type posHeap []pos

func (h posHeap) Len() int {
	return len(h)
}
func (h posHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}
func (h posHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *posHeap) Push(x interface{}) {
	*h = append(*h, x.(pos))
}
func (h *posHeap) Pop() interface{} {
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
