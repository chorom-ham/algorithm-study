#include<cstdio>
#include<queue>
#include<vector>
#include<cstring>
#include<algorithm>

#define MAX 100
#define INF 10000

using namespace std;

int n, islandNum, answer;
int map[MAX][MAX];                  // 입력 받아 지도 값 저장(0:바다/1:육지) >> 이후 섬 번호 붙여줌

bool visit[MAX][MAX];                    // BFS탐색 시, 방문 여부 체크용
vector<pair<int, int>> v;                // 입력 받으면서 모든 섬의 좌표들을 저장하기 위한 벡터

int dx[] = { 0, 0, 1, -1 };
int dy[] = { 1, -1, 0, 0 };

void input() {
	scanf("%d", &n);
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			scanf("%d", &map[i][j]);
			if (map[i][j] == 1) {
				v.push_back(make_pair(i, j)); //섬의 좌표들 저장
				map[i][j] = -1; //나중에 섬 번호 붙여줄 때 1번부터 붙여주기 위함
			}
		}
	}
}

//연결된 땅 탐색하면서 라벨링. (a,b): 탐색 시작 좌표, num: 탐색할 섬에 붙일 번호
void BFSlabel(int a, int b, int num) {
	queue<pair<int, int>> q;
	q.push(make_pair(a, b));
	visit[a][b] = true;
	map[a][b] = num;

	while (!q.empty()) {
		int x = q.front().first;
		int y = q.front().second;
		q.pop();

		for (int i = 0; i < 4; i++) {
			int nx = x + dx[i];
			int ny = y + dy[i];

			if (nx >= 0 && ny >= 0 && nx < n && ny < n) { //범위 체크
				if (!visit[nx][ny] && map[nx][ny] == -1) { //(nx,ny)가 방문하지 않은 땅이면
					visit[nx][ny] = true;
					map[nx][ny] = num;
					q.push(make_pair(nx, ny));
				}
			}
		}
	}
}

//입력 시 저장해놓은 v를 순차적으로 탐색하면서 아직 방문하지 않은 좌표에 대해서 BFS 탐색 및 라벨링
void label() {
	int num = 1;
	for (int i = 0; i < v.size(); i++) {
		int x = v[i].first;
		int y = v[i].second;

		if (!visit[x][y]) {
			BFSlabel(x, y, num++);
		}
	}
	islandNum = num - 1;
}

int BFS(int start) {
	int bridgeSize = 0;
	queue<pair<int, int>> q;

	//BFS돌리기 전 큐에 같은 섬 좌표 다 넣기
	for (int i = 0; i < n; i++)	{
		for (int j = 0; j < n; j++)	{
			if (map[i][j] == start)	{
				visit[i][j] = true;
				q.push(make_pair(i, j));
			}
		}
	}

	while (!q.empty()) {
		int size = q.size();
		for (int k = 0; k < size; k++) {
			int x = q.front().first;
			int y = q.front().second;
			q.pop();

			for (int i = 0; i < 4; i++) {
				int nx = x + dx[i];
				int ny = y + dy[i];

				if (nx >= 0 && ny >= 0 && nx < n && ny < n) { //범위 체크
					if (map[nx][ny] != 0 && map[nx][ny] != start) { //시작 섬이 아닌 다른 섬에 도착했으면
						return bridgeSize;
					}
					if (!visit[nx][ny] && map[nx][ny] == 0) { //(nx,ny)가 방문한 적 없는 바다면
						visit[nx][ny] = true;
						q.push(make_pair(nx, ny));
					}
				}
			}
		}
		bridgeSize++;
	}
}

int main(void) {
	input(); //입력 받기
	label(); //섬에 번호 붙이기(1,2,3...)
	answer = INF;
	for (int i = 1; i < islandNum + 1; i++) {
		memset(visit, false, sizeof(visit));
		answer = min(answer, BFS(i)); //i섬 출발 최단 거리(다리)계산(최소값 갱신)
	}

	printf("%d", answer);

	return 0;
}
