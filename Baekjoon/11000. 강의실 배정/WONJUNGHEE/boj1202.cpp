#include <iostream>
#include <queue>
#include <algorithm>

using namespace std;

int N;
int answer;
priority_queue<int, vector<int>, greater<int>> ing;

int main()
{
    pair<int, int> time[200001];
    cin >> N;
    for (int i = 0; i < N; i++)
    {
        cin >> time[i].first >> time[i].second;
    }
    sort(time, time + N);

    ing.push(time[0].second);
    for (int i = 1; i < N; i++)
    {
        if (ing.top() <= time[i].first)
        {
            ing.pop();
            ing.push(time[i].second);
        }
        else
        {
            ing.push(time[i].second);
        }
    }

    cout << ing.size();

    return 0;
}
