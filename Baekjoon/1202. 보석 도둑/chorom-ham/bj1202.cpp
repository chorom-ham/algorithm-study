#include <cstdio>
#include <utility>
#include <queue>
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

	priority_queue<int> pq;
	for (int i = 0, j = 0; i < k; ++i) {
		while (j < n && jewelry[j].first <= bag[i]) {
			pq.push(jewelry[j++].second);
		}
		if (!pq.empty()) {
			ans += pq.top();
			pq.pop();
		}
	}

	printf("%lld", ans);
}
