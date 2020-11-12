한번에 짠 아래 코드.

테스트케이스는 문제 없이 동작했는데 채점 돌리니까 7%에서 시간 초과 났다.

고민하다가 결국 우선순위 큐 써서 해결했다.

```cpp
#include <cstdio>
#include <utility>
#include <algorithm>
#define MAX 300000

using namespace std;

long long ans = 0;
int n, k;
pair<int, int> jewelry[MAX]; // {무게, 가치}
int bag[MAX];
bool include[MAX]; // 보석 넣었는지

int main() {
	scanf("%d %d", &n, &k);
	for (int i = 0; i < n; i++) {
		scanf("%d %d", &jewelry[i].first, &jewelry[i].second);
	}
	sort(jewelry, jewelry + n); // 무게 작은 순으로 정렬

	for (int i = 0; i < k; i++) {
		scanf("%d", &bag[i]);
	}
	sort(bag, bag + k); //오름차순 정렬(가방 크기 작은 순부터 큰 순으로)

	int index, maxV = 0; // 최대 가치 보석 인덱스와 가치
	// 가방 안에 들어가는 보석 중에서 가장 가치가 높은 보석 넣기
	for (int i = 0; i < k; i++) {
		for (int j = 0; j < n; j++) {
			if (jewelry[j].first > bag[i]) { //보석 무게가 가방 크기를 넘어서면
				break;
			}
			if (include[j]) {
				continue;
			}
			if (maxV < jewelry[j].second) {
				maxV = jewelry[j].second;
				index = j;
			}
		}
		ans += maxV;
		include[index] = true;
		maxV = 0;
	}

	printf("%lld", ans);
}
```