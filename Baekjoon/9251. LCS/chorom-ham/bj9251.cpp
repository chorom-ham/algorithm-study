#include <cstdio>
#include <iostream>
#include <string>
#include <algorithm>
using namespace std;

int dp[1001][1001];

int main() {
	string str1; 
	string str2;
	
	getline(cin, str1);
	getline(cin, str2);

	int iMax = str1.size();
	int jMax = str2.size();

	str1 = '0' + str1;
	str2 = '0' + str2;

	for (int i = 1; i < iMax + 1; i++) {
		for (int j = 1; j < jMax + 1; j++) {
			if (str1[i] == str2[j]) {
				dp[i][j] = dp[i - 1][j - 1] + 1;
			}
			else {
				dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]);
			}
		}
	}
	
	printf("%d", dp[iMax][jMax]);

	return 0;
}