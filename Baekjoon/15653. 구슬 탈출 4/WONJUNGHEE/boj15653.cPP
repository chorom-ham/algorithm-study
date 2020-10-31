#include <iostream>
#include <queue>
#include <vector>
#include <string>
using namespace std;

int H, W;
bool check[101][101];
int dx[] = {1, -1, 0, 0};
int dy[] = {0, 0, 1, -1};
char map[101][101];
int dcount[101][101];

int main()
{
    ios::sync_with_stdio(0), cin.tie(0), cout.tie(0);

    cin >> W >> H;
    string s;
    vector<pair<int, int>> cpoint;
    for (int i = 0; i < H; i++)
    {
        for (int j = 0; j < W; j++)
        {
            cin >> map[i][j];
            if (map[i][j] == 'C')
            {
                cpoint.push_back(make_pair(i, j));
            }
        }
    }
    queue<pair<int, int>> que;
    int cx = cpoint[0].first;
    int cy = cpoint[0].second;
    que.push(make_pair(cx, cy));
    check[cx][cy] = true;
    dcount[cx][cy] = 0;
    while (!que.empty())
    {
        int x = que.front().first;
        int y = que.front().second;
        que.pop();
        for (int k = 0; k < 4; k++)
        {
            int nx = x + dx[k];
            int ny = y + dy[k];
            while (nx >= 0 && nx < H && ny >= 0 && ny < W)
            {
                if (map[nx][ny] == '*')
                    break;
                if (!check[nx][ny])
                {
                    check[nx][ny] = true;
                    dcount[nx][ny] = dcount[x][y] + 1;
                    que.push(make_pair(nx, ny));
                }
                nx += dx[k];
                ny += dy[k];
            }
        }
    }
    cout << dcount[cpoint[1].first][cpoint[1].second] - 1;
    return 0;
}
