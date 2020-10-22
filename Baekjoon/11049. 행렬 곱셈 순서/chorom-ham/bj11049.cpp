#include <cstdio>
#include <algorithm>
using namespace std;

//matrix[i][0]: i-th matrix's row num
//matrix[i][1]: i-th matrix's column num
int matrix[501][2]; 
int N;

//dp[i][j]: i번째 행렬부터 j번째 행렬 사이의 최소 곱셈 연산 횟수(1<=i<=j<=n)
int dp[501][501]; 

int main(void) {
	scanf("%d", &N);

	for (int i = 1; i < N + 1; i++) {
		scanf("%d %d", &matrix[i][0], &matrix[i][1]);
	}

	//i == j일 때 dp[i][j] == 0
	//i+1 == j일 때 dp[i][j]는 아래와 같음
	for (int i = 1; i < N; i++) {
		dp[i][i + 1] = matrix[i][0] * matrix[i][1] * matrix[i + 1][1];
	}

	for (int m = 2; m <= N; m++) { //m == j-i (차)
		for (int i = 1; i <= N-m; i++) {
			int j = i + m;
			for (int k = i; k < j; k++) {
				int value = dp[i][k] + dp[k + 1][j] + matrix[i][0] * matrix[k][1] * matrix[j][1];
				if (!dp[i][j] || dp[i][j] > value) {
					dp[i][j] = value;
				}
			}
		}
	}

	printf("%d", dp[1][N]);

	return 0;
}