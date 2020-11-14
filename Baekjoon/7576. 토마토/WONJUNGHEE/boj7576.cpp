
#include <iostream>
#include <queue>

using namespace std;
int map[1001][1001];
bool visited[1001][1001];
queue<pair<pair<int, int>, int>> tmt;
int dx[4] = {0, 0, -1, 1};
int dy[4] = {-1, 1, 0, 0};
int cnt;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(0);
    cout.tie(0);

    int a, b;
    cin >> a >> b;
    for (int i = 0; i < b; i++)
    {
        for (int j = 0; j < a; j++)
        {
            cin >> map[i][j];
            if (map[i][j] == 1)
            {
                visited[i][j] = true;
                tmt.push(make_pair(make_pair(i, j), 0));
            }
        }
    }

    while (!tmt.empty())
    {
        int x = tmt.front().first.first;
        int y = tmt.front().first.second;
        cnt = tmt.front().second;

        tmt.pop();

        for (int i = 0; i < 4; i++)
        {
            int nx = x + dx[i];
            int ny = y + dy[i];

            if (nx < 0 || nx >= b || ny < 0 || ny >= a)
                continue;

            if (map[nx][ny] == 0 && !visited[nx][ny])
            {
                visited[nx][ny] = true;
                tmt.push(make_pair(make_pair(nx, ny), cnt + 1));
            }
        }
    }

    for (int i = 0; i < b; i++)
    {
        for (int j = 0; j < a; j++)
        {
            if (map[i][j] == 0 && !visited[i][j])
            {
                cnt = -1;
                break;
            }
        }
    }
    cout << cnt;
    return 0;
}