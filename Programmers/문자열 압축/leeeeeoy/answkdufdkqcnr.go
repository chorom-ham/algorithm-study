package main

import (
	"fmt"
	"strconv"
)

//	java코드로 작성 후
//	Go 코드로 다시 작성
func main() {
	var str string
	fmt.Scan(&str)
	fmt.Print(solution(str))
}
func solution(str string) int {
	answer := len(str) // 초기값 문자열 길이
	//	1,2,3... -> 한번에 끊을 문자열, 최대 절반까지
	for i := 1; i < len(str)/2; i++ {
		var temp string
		//	처음부터 i씩 귾어서 비교
		for j := 0; j < len(str); j += i {
			cur := ""
			if i+j >= len(str) { //	최대길이가 되면 끝까지 자르기
				cur = str[j:len(str)]
			} else {
				cur = str[j : j+i] //	i칸씩 자르기
			}
			count := 1
			var sb string

			for k := i + j; k < len(str); k += i {
				cur2 := ""

				if k+i >= len(str) { //	최대 길이가 되면 끝까지 자르기
					cur2 = str[k:len(str)]
				} else {
					cur2 = str[k : k+i] //	i칸씩 자르기
				}

				if cur == cur2 { //	같으면 횟수를 증가시키고 자를 위치를 옴긴다
					count++
					j = k
				} else { //	틀리면 다시 다음 문자부터 비교
					break
				}
			}
			if count == 1 {
				sb = cur
			} else {
				sb = strconv.Itoa(count) + cur
			}
			temp = sb
		}
		if answer > len(temp) {
			answer = len(temp)
		}
	}
	return answer
}
