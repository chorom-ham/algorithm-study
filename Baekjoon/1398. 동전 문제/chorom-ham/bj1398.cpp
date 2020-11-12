#include <cstdio>
#include <cstring>
#include <algorithm>

using namespace std;

int t; // 테스트케이스 개수
long long ans;
long long price;
int dp[100]; 

// dp 배열 값 계산
void calc() {
	for (int i = 0; i < 100; i++) {
		if (i >= 25) {
			dp[i] = min(dp[i - 25] + 1, dp[i - 10] + 1);
			dp[i] = min(dp[i], dp[i - 1] + 1);
		}
		else if (i >= 10) {
			dp[i] = min(dp[i - 10] + 1, dp[i - 1] + 1);
		}
		else{
			dp[i] = i;
		}
	}
}

// 화폐 단위가 100단위로 곱해진다는 것에 주목 {1,10,25}*(100^k)
void solve() {
	long long tmp = price % 100;
	ans += dp[tmp];
	price /= 100;
	if (price > 0) {
		solve();
	}
}

int main() {
	// dp 배열 계산하고 시작
	memset(dp, 100, sizeof(dp));
	calc();

	scanf("%d", &t);
	for (int i = 0; i < t; i++) {
		scanf("%lld", &price);
		ans = 0;
		solve();
		printf("%lld\n", ans);
	}

	return 0;
}
