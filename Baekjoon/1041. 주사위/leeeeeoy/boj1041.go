package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

//	하드코딩...
func main() {
	n := nextInt()
	sum := 0
	//	A B C D E F
	//	A - F, D - C, B-E
	// 0 - 5, 2 - 3, 1 - 4
	dice := make([]int, 6)
	for i := 0; i < 6; i++ {
		dice[i] = nextInt()
	}

	if n == 1 {
		sort.Ints(dice)
		for _, num := range dice {
			sum += num
		}
		fmt.Print(sum - dice[5])
		return
	}
	s := make([]int, 3) //	마주보는 양면 중 작은 값
	r := make([]int, 3) //	각 면마다 보이는 주사위 갯수
	f := make([]int, 3) //	주사위의 최솟값

	s[0] = minInt(dice[0], dice[5])
	s[1] = minInt(dice[1], dice[4])
	s[2] = minInt(dice[2], dice[3])

	sort.Ints(s)

	r[0] = (n-1)*(n-2)*4 + (n-2)*(n-2) //	1면이 보이는 주사위 개수
	r[1] = (n-1)*4 + (n-2)*4           //	2면이 보이는 주사위 개수
	r[2] = 4                           //	3면이 보이는 주사위 개수

	f[0] = s[0]               //	1면이 보이는 주사위의 최솟값
	f[1] = s[0] + s[1]        //	2면이 보이는 주사위의 최솟값
	f[2] = s[0] + s[1] + s[2] //	3면이 보이는 주사위의 최솟값

	sum += f[0] * r[2]
	sum += f[1] * r[1]
	sum += f[2] * r[0]

	fmt.Print(sum)
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
