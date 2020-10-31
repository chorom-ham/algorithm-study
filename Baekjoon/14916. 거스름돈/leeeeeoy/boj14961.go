package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 || n == 3 { //	1 or 3 일때 불가능
		fmt.Print(-1)
	} else if n == 4 { //	4일 때는 2
		fmt.Println(2)
	} else { //	5이상일 때는 5로 나누어주고, 나머지는 2로 나누어줌, 홀수면 5를 더해서 2로 나눠줌
		count := 0
		count += n / 5
		n %= 5
		count += n / 2
		n %= 2
		if n == 1 {
			n += 5
			count += n/2 - 1
			n %= 2
		}
		fmt.Print(count)
	}
}
