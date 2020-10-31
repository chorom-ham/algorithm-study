#include <iostream>
#include <queue>
using namespace std;

int N, M;
int rx, ry;
int bx, by;
int goalx, goaly;
bool flag;
int answer = 0;
char map[11][11];
int dx[4] = {-1, 1, 0, 0};
int dy[4] = {0, 0, -1, 1};
int visited[11][11][11][11];
queue<pair<int, int>> redball;
queue<pair<int, int>> blueball;

int main()
{
    ios::sync_with_stdio(0), cin.tie(0), cout.tie(0);

    cin >> N >> M;
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; j < M; j++)
        {
            cin >> map[i][j];
            if (map[i][j] == 'R')
            {
                redball.push(make_pair(i, j));
                rx = i, ry = j;
            }
            if (map[i][j] == 'B')
            {
                blueball.push(make_pair(i, j));
                bx = i, by = j;
            }
            if (map[i][j] == 'O')
                goalx = i, goaly = j;
        }
    }
    visited[rx][ry][bx][by] = true;

    while (!redball.empty())
    {
        int qsize = redball.size();
        while (qsize--)
        {
            rx = redball.front().first;
            ry = redball.front().second;
            bx = blueball.front().first;
            by = blueball.front().second;
            redball.pop();
            blueball.pop();

            if (rx == goalx && ry == goaly)
            {
                flag = true;
                break;
            }

            for (int i = 0; i < 4; i++)
            {
                int nrx = rx + dx[i];
                int nry = ry + dy[i];
                int nbx = bx + dx[i];
                int nby = by + dy[i];

                while (1)
                {
                    if (map[nrx][nry] == '#')
                    {
                        nrx -= dx[i];
                        nry -= dy[i];
                        break;
                    }
                    if (map[nrx][nry] == 'O')
                    {
                        break;
                    }
                    nrx += dx[i];
                    nry += dy[i];
                }
                while (1)
                {
                    if (map[nbx][nby] == '#')
                    {
                        nbx -= dx[i];
                        nby -= dy[i];
                        break;
                    }
                    if (map[nbx][nby] == 'O')
                    {
                        break;
                    }
                    nbx += dx[i];
                    nby += dy[i];
                }

                if (nbx == goalx && nby == goaly)
                    continue;

                if (nrx == nbx && nry == nby)
                {
                    switch (i)
                    {
                    case 0:
                        rx > bx ? nrx++ : nbx++;
                        break;
                    case 1:
                        rx < bx ? nrx-- : nbx--;
                        break;
                    case 2:
                        ry > by ? nry++ : nby++;
                        break;
                    case 3:
                        ry < by ? nry-- : nby--;
                        break;
                    }
                }
                if (visited[nrx][nry][nbx][nby])
                    continue;
                redball.push(make_pair(nrx, nry));
                blueball.push(make_pair(nbx, nby));
                visited[nrx][nry][nbx][nby] = true;
            }
        }

        if (flag)
            break;
        else
            answer++;
    }
    if (flag)
        cout << answer;
    else
        cout << -1;

    return 0;
}
