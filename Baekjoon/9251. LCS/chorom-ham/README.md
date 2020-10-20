# 9251. LCS

## 풀이 방법

dp[i][j] : 1번째 문자열의 첫 i개의 문자열과 2번째 문자열의 첫 j개의 문자열의 LCS.

### dp[i][j]를 구하는 법
```cpp
if (str1[i] == str2[j]) {
	dp[i][j] = dp[i - 1][j - 1] + 1; //표의 왼쪽 대각선 위 값 + 1
}
else {
	dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]); // 왼쪽 값과 윗쪽 값 중 큰 것 
}   
```

+) 문자열 앞에 '0'을 넣어준 이유?  배열은 0부터 시작한다. 첫번째 문자를 arr[1]로 사용하고 싶었기 때문에 배열의 0번 인덱스의 아무 의미 없는 값을 넣었다.


## 어려웠던 부분 혹은 새로 알게 된 내용

subsequence와 substring의 개념이 헷갈려 처음에 문제 이해를 제대로 못했다.

- subsequence: 연속적일 필요는 없지만 순서는 지켜야 하는 부분 문자열.

- substring: 연속적인 부분 문자열. 순서 지켜야 함.

- 예시) iamchorom의 subsequence는 achm, substring은 amch

어떻게 푸는 지 감이 안 잡혀서 인터넷에 푸는 방법을 검색해 보았다. (https://twinw.tistory.com/126)

이전에 풀어봤던 edit distance문제랑 접근법이 비슷했다. 두 문자열 비교하는 문제는 일단 표를 만들어서 규칙을 찾아보자.