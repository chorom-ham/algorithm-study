#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int cache[501][501];
vector<pair<int, int>> arr;
int MIN_MIN = 9384920123;

int chain(int a, int b)
{
    int &ret = cache[a][b];
    if (ret != 0)
        return ret;
    if (a == b)
        return 0;
    ret = MIN_MIN;

    for (int i = a; i < b; i++)
    {
        ret = min(ret, chain(a, i) + chain(i + 1, b) + arr[a].first * arr[i].second * arr[b].second);
    }

    return ret;
}

int main()
{
    int N;
    int x, y;
    cin >> N;
    for (int i = 0; i < N; i++)
    {
        cin >> x >> y;
        arr.push_back({x, y});
    }
    int ret = chain(0, N - 1);

    cout << ret;
    return 0;
}