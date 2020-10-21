#include <cstdio>
#include <cstring>
#include <algorithm>
#define INF 16000000
using namespace std;

int N; 
int W[16][16]; 
int dp[16][1 << 16]; 
int complete;

int dfs(int current, int path) {
	if (path == complete) {
		if (W[current][0] == 0) {
			return INF;
		}
		else {
			return W[current][0];
		}
	}

	if (dp[current][path] != -1) {
		return dp[current][path];
	}

	dp[current][path] = INF;

	for (int i = 0; i < N; ++i) {
		if (W[current][i] == 0) {
			continue;
		}
		if ((path & (1 << i)) == (1 << i)) {
			continue;
		}
		dp[current][path] = min(dp[current][path], W[current][i] + dfs(i, path | 1 << i));
	}
	return dp[current][path];
}

int main(void) {
	scanf("%d", &N);
	complete = (1 << N) - 1;

	for (int i = 0; i < N; ++i) {
		for (int j = 0; j < N; ++j) {
			scanf("%d", &W[i][j]);
		}
	}

	memset(dp, -1, sizeof(dp));

	printf("%d", dfs(0, 1));

	return 0;
}
