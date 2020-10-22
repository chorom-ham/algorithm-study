#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int dist[16][16];
int cache[16][1 << 16];
int N;
int INF = 100000007;

int dp(int here, int start, int visited)
{
    if (visited == (1 << N) - 1)
    {
        if (dist[here][start] != 0)
            return dist[here][start];
        else
        {
            return INF;
        }
    }

    int &ret = cache[here][visited];
    if (ret > 0)
        return ret;

    ret = INF;
    for (int next = 0; next < N; next++)
    {
        if (visited & (1 << next) || dist[here][next] == 0)
            continue;
        ret = min(ret, dist[here][next] + dp(next, start, visited + (1 << next)));
    }

    return ret;
}

int main()
{
    cin >> N;
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; j < N; j++)
        {
            cin >> dist[i][j];
        }
    }
    int ret = INF;

    cout << dp(0, 0, 1);

    return 0;
}