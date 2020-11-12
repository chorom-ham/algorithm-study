package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
	t  int
)

//	10^k, 25 * 100^k
func main() {
	defer bw.Flush()
	fmt.Fscan(br, &t)
	a := []int{25, 10}
	coin := [100]int{} //	100단위로 끊어진다

	for i := 1; i < 100; i++ {
		coin[i] = i
	}

	for i := 0; i < 2; i++ {
		for j := a[i]; j < 100; j++ {
			coin[j] = minInt(coin[j], coin[j-a[i]]+1)
		}
	}

	for i := 0; i < t; i++ {
		var car, count int
		fmt.Fscan(br, &car)
		for car > 0 {
			count += coin[car%100]
			car /= 100
		}
		fmt.Fprintln(bw, count)
	}
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//	34 -> 25 + 9  =  10
//	   -> 10+10+10+4 = 7

// func 오답() {
// 	defer bw.Flush()
// 	fmt.Fscan(br, &t)
// 	coin := []int{
// 		1, 10, 25, 100, 1000, 2500,
// 		10000, 100000, 250000, 1000000, 10000000,
// 		25000000, 100000000, 1000000000, 2500000000,
// 		10000000000, 100000000000, 250000000000, 1000000000000,
// 		10000000000000, 25000000000000, 100000000000000, 1000000000000000,
// 	}
// 	for i := 0; i < t; i++ {
// 		var car, count int
// 		fmt.Fscan(br, &car)
// 		for j := 22; j >= 0; j-- {
// 			if coin[j] <= car {
// 				count += car / coin[j]
// 				car %= coin[j]
// 			}
// 			if car == 0 {
// 				break
// 			}
// 		}
// 		fmt.Fprintln(bw, count)
// 	}
// }
