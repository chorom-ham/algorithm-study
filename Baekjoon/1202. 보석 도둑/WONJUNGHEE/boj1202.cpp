#include <string>
#include <iostream>
#include <queue>
#include <algorithm>

using namespace std;

int n, k;
vector<pair<int, int>> ju;
priority_queue<int> weight;
int bag[300005];

int main()
{
    cin >> n >> k;
    for (int i = 0; i < n; i++)
    {
        int a, b;
        cin >> a >> b;
        ju.push_back({a, b});
    }
    for (int j = 0; j < k; j++)
    {
        cin >> bag[j];
    }
    sort(ju.begin(), ju.end());
    sort(bag, bag + k);
    int index = 0;
    long long ans = 0;
    for (int i = 0; i < k; i++)
    {
        while (index < n && ju[index].first <= bag[i])
        {
            weight.push(ju[index].second);
            index++;
        }
        if (!weight.empty())
        {
            ans += weight.top();
            weight.pop();
        }
    }
    cout << ans;

    return 0;
}