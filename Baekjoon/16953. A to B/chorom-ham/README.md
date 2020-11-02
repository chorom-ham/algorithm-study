처음에는 간단하게 b >> a로 만드는 데 몇번의 연산이 필요한 지 구하는 알고리즘을 짰다.

반복해서 b가 짝수면 2로 나누고, 1로 끝나면 오른쪽에서 1을 빼주는(10으로 나눠서) 방식을 사용했다.

a == b가 되면 답을 프린트하고, a>b가 되면 바꿀 수 없는 경우로 -1을 출력했다.

```cpp
#include <cstdio>

int main() {
	int a, b, count = 1;
	scanf("%d %d", &a, &b);

	//거꾸로 계산해 나가기
	while (a <= b) {
		if (a == b) {
			printf("%d", count);
			return 0;
		}
		if (b % 2 == 0) {
			b /= 2;
			count++;
		}
		if (b % 10 == 1) {
			b /= 10;
			count++;
		}
	}

	printf("-1");
	return 0;
}
```

백준에서 채점해보니 시간 초과가 났다.