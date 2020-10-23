#include <iostream>
#include <cstring>
#include <string>

using namespace std;

int dp[1002][1002];

int main()
{
    char first[1002], second[1002];
    cin >> first >> second;
    for (int i = 1; i <= strlen(first); i++)
    {
        for (int j = 1; j <= strlen(second); j++)
        {
            if (first[i - 1] == second[j - 1])
            {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            }
            else
            {
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]);
            }
        }
    }
    cout << dp[strlen(first)][strlen(second)];
}