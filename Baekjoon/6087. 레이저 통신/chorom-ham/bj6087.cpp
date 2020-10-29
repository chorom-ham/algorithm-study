#include <iostream>
#include <utility>
#include <algorithm>
#include <cstring>
#include <queue>
using namespace std;

char arr[100][100];
int w, h;
pair<int, int> c1, c2; //c 좌표(c1출발 c2도착)
int visit[100][100]; //visit[i][j]: (i,j)로 가는 데 필요한 거울 최소 개수

//상 하 좌 우
int dx[4] = { 0, 0, -1, 1 };
int dy[4] = { 1, -1, 0, 0 };

int search(int a, int b) {
	//q: {{현재 x좌표, 현재 y좌표}, {현재 진행 방향, 지금까지 놓은 거울 개수}}
	queue<pair<pair<int, int>, pair<int, int>>> q;

	for (int i = 0; i < 4; i++) {
		q.push(make_pair(make_pair(a, b), make_pair(i, 0)));
	}
	visit[a][b] = 0;

	while (!q.empty()) {
		pair<pair<int, int>, pair<int, int>> p = q.front();
		int x = p.first.first;
		int y = p.first.second;
		int dir = p.second.first;
		int mirror = p.second.second;
		q.pop();

		for (int i = 0; i < 4; i++) {
			int nx = x + dx[i];
			int ny = y + dy[i];
			int current = mirror;

			//다음 위치가 범위 안에 있는지 체크
			if (nx < 0 || ny < 0 || nx > h - 1 || ny > w - 1) continue;
			//다음 위치가 벽이 아닌지 체크
			if (arr[nx][ny] == '*') continue;
			//원래 가던 방향과 지금 가려는 방향이 다르면 거울 개수 추가
			if (dir != i) current++;
			//탐색하고자 하는 경로가 현재 놓은 거울 개수보다 큰지 확인하고 탐색 여부 결정(큐에 넣어 탐색)
			if (visit[nx][ny] >= current) {
				visit[nx][ny] = current;
				q.push(make_pair(make_pair(nx, ny), make_pair(i, current)));
			}
		}
	}
	return visit[c2.first][c2.second];
}

int main() {
	cin >> w >> h;
	int tmp = 0;
	for (int i = 0; i < h; i++) {
		for (int j = 0; j < w; j++) {
			cin >> arr[i][j];
			if (arr[i][j] == 'C') {
				if (tmp == 0) {
					c1 = make_pair(i, j);
					tmp++;
				}
				else {
					c2 = make_pair(i, j);
				}
			}
		}
	}
	memset(visit, 10000, sizeof(visit));

	cout << search(c1.first, c1.second);
}