#include <cstdio>
#include <cstring>
#include <algorithm>
#include <vector>
#define MAX 10001
using namespace std;

int n;						// 정점 개수
int vertex[MAX];			// 정점의 가중치
vector<int> edge[MAX];		// 그래프
//dp[i][0]: i번째 노드를 선택하지 않는 경우 최대 독립 집합
//dp[i][0]: i번째 노드를 선택하는 경우 최대 독립 집합
int dp[MAX][2];
vector<int> ans; //최대 독립집합에 속하는 정점

int solution(int prev, int now, bool include) {
	int& result = dp[now][include];

	if (result != -1) {
		return result;
	}
	if (include) {
		result = vertex[now];
	}
	else {
		result = 0;
	}

	for (int i = 0; i < edge[now].size(); i++) {
		int next = edge[now][i];
		if (next == prev) {
			continue;
		}
		if (include) {
			result += solution(now, next, 0);
		}
		else {
			// 이 노드를 선택 안한 경우
			result += max(solution(now, next, 0), solution(now, next, 1));
		}
	}
	return result;
}

void tracking(int prev, int now, bool include) {
	if (include) {
		ans.push_back(now);
		for (auto next : edge[now]) {
			if (next == prev) {
				continue;
			}
			tracking(now, next, 0);
		}
	}
	else {
		for (auto next : edge[now]) {
			if (next == prev) {
				continue;
			}
			if (dp[next][1] >= dp[next][0]) {
				tracking(now, next, 1);
			}
			else {
				tracking(now, next, 0);
			}
		}
	}
}

int main() {
	memset(dp, -1, sizeof(dp));

	scanf("%d", &n);

	for (int i = 1; i <= n; i++) {
		scanf("%d", &vertex[i]);
	}

	for (int i = 0; i < n - 1; i++) {
		int v, e;
		scanf("%d %d", &v, &e);
		edge[v].push_back(e);
		edge[e].push_back(v);
	}

	int temp1 = solution(-1, 1, 0);
	int temp2 = solution(-1, 1, 1);

	// dp는 첫번째 노드부터 각 노드의 정점까지의 독립집합 최댓값. 최댓값 따라가며 트래킹.
	if (temp1 > temp2) {
		tracking(-1, 1, 0);
	}
	else {
		tracking(-1, 1, 1);
	}

	sort(ans.begin(), ans.end());

	printf("%d\n", max(temp1, temp2));

	for (auto i : ans) {
		printf("%d ", i);
	}

	return 0;
}
