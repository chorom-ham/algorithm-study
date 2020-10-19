#include <cstdio>

int dp[1001];

int main(void) {
	int n;
	scanf("%d", &n);

	for (int i = 1; i < n + 1; ++i) {
		if (i == 1) {
			dp[i] = 1;
			continue;
		}
		else if (i == 2) {
			dp[i] = 3;
			continue;
		}
		else if (i > 2) {
			dp[i] = (dp[i - 1] + dp[i - 2] * 2) % 10007;
		}
	}

	printf("%d", dp[n]);

	return 0;
}