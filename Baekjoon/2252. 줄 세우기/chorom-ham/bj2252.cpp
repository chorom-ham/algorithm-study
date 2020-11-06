#include<cstdio>
#include<queue>
#include<vector>

#define MAX 32001

using namespace std;

int indegree[MAX];
vector<int> v[MAX];

int main() {
	int n, m, a, b;
	scanf("%d %d", &n, &m);
	for (int i = 0; i < m; i++) {
		scanf("%d %d", &a, &b);
		indegree[b]++;
		v[a].push_back(b);
	}

	queue<int> q;
	for (int i = 1; i <= n; i++) {
		if (indegree[i] == 0) q.push(i);
	}

	while (!q.empty()) {
		int i = q.front();
		q.pop();
		printf("%d ", i);

		for (int j = 0; j < v[i].size(); j++) {
			if (--indegree[v[i][j]] == 0) {
				q.push(v[i][j]);
			}
		}
	}
}
