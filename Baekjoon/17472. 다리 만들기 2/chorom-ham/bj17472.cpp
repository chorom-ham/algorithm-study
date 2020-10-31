//잘 모르겠어서 https://yabmoons.tistory.com/292 의 글을 참고했다

#include<cstdio>
#include<queue>
#include<vector>
#include<cstring>

#define MAX 10
#define ISLAND_MAX 6 + 1  //섬은 1번부터 라벨링(0번 X)
#define INF 1000

using namespace std;

int n, m, islandNum, answer = INF;
int map[MAX][MAX];                  // 입력 받아 지도 값 저장(0:바다/1:육지)
int labeledMap[MAX][MAX];                // 각 섬마다 번호를 붙이기 위해 사용한 맵
// 각 섬 사이의 최단거리를 저장하기 위한 배열 
int dist[ISLAND_MAX][ISLAND_MAX];   // dist[a][b] = c : a와 b의 최단거리는 c

bool visit[MAX][MAX];                    // BFS탐색 시, 방문 여부 체크용(섬에 번호 붙일 때 사용)
bool connect[ISLAND_MAX][ISLAND_MAX];    // 연결관계 체크를 위한 배열
bool connectIsland[ISLAND_MAX];          // BFS탐색 시, 방문 여부 체크용(연결관계를 통해, 정답을 도출하기 위한 BFS탐색 시 사용)
bool select[ISLAND_MAX * ISLAND_MAX];    // 조합 추출에서 중복 추출을 막기 위한 배열
/* Select배열의 크기 : 7 * 7
	 섬이 N(2<=N<=6)개 존재하고, 이 섬들을 연결한다고 가정했을 때
	 나올 수 있는 간선의 최대 갯수는 N(N-1)/2 개 (조합)
*/

vector<pair<int, int>> v;                       // 입력 받으면서 모든 섬의 좌표들을 저장하기 위한 벡터
vector<pair<int, int>> islandPosition[MAX]; // 각 섬에 존재하는 땅의 좌표들을 저장하기 위한 벡터 배열
vector<pair<pair<int, int>, int>> bridgeList;   // 섬과 섬을 연결하는 다리의 길이와, 그 다리가 어떤 섬을 연결하는지 저장하기 위한 벡터.

int dx[] = { 0, 0, 1, -1 };
int dy[] = { 1, -1, 0, 0 };

void input() {
	for (int i = 0; i < 7; i++) {
		for (int j = 0; j < 7; j++)	{
			dist[i][j] = INF;
		}
	}

	scanf("%d %d", &n, &m);
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < m; j++) {
			scanf("%d", &map[i][j]);
			if (map[i][j] == 1) {
				v.push_back(make_pair(i, j)); //섬의 좌표들 저장
			}
		}
	}
}

//연결된 땅 탐색하면서 라벨링. (a,b): 탐색 시작 좌표, num: 탐색할 섬에 붙일 번호
void BFS(int a, int b, int num) {
	queue<pair<int, int>> q;
	q.push(make_pair(a, b));
	visit[a][b] = true;
	labeledMap[a][b] = num;
	islandPosition[num].push_back(make_pair(a, b));

	//bfs
	while (!q.empty()) {
		int x = q.front().first;
		int y = q.front().second;
		q.pop();

		for (int i = 0; i < 4; i++) {
			int nx = x + dx[i];
			int ny = y + dy[i];

			if (nx >= 0 && ny >= 0 && nx < n && ny < m) { //범위 체크
				if (!visit[nx][ny] && map[nx][ny] == 1) { //(nx,ny)가 방문하지 않은 땅이면
					visit[nx][ny] = true;
					labeledMap[nx][ny] = num;
					q.push(make_pair(nx, ny));
					islandPosition[num].push_back(make_pair(nx, ny));
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
			BFS(x, y, num++);
		}
	}
	islandNum = num - 1;
}

void DFS(int x, int y, int dir, int bridgeSize, int start, int end) {
	int nx = x + dx[dir];
	int ny = y + dy[dir];

	if (nx < 0 || ny < 0 || nx >= n || ny >= m) return; //범위 체크

	if (map[nx][ny] == 0) {
		DFS(nx, ny, dir, bridgeSize + 1, start, end); //아직 바다라면 다리 개수 하나 추가하고 계속 탐색
	}
	else if (map[nx][ny] == 1) { //육지에 도착했으면
		if (labeledMap[nx][ny] == end) { //도착섬에 도착했으면
			if (bridgeSize > 1) {  //다리 최소 개수 == 2
				if (dist[start][end] > bridgeSize) {  //최소 길이로 갱신
					dist[start][end] = bridgeSize;
					dist[end][start] = bridgeSize;
				}
				return;
			}
		}
		else return;
	}
}

//start섬에서 end섬까지의 최단 거리 탐색(다리)
void makeBridge(int start, int end) {
	//start섬의 모든 좌표로부터 탐색 시작
	for (int i = 0; i < islandPosition[start].size(); i++) {
		int x = islandPosition[start][i].first;
		int y = islandPosition[start][i].second;

		for (int j = 0; j < 4; j++) {
			int nx = x + dx[j];
			int ny = y + dy[j];

			if (nx >= 0 && ny >= 0 && nx < n && ny < m) { //범위 체크
				if (map[nx][ny] == 0) {
					DFS(x, y, j, 0, start, end);
				}
			}
		}
	}
}

//DFS 사용하여 각 섬 사이의 최단 거리 계산
void calculateDistance() {
	for (int i = 1; i < islandNum + 1; i++) {
		for (int j = i + 1; j < islandNum + 1; j++) {
			makeBridge(i, j); //i섬과 j사이의 최단 거리(다리)계산 >> dist배열 갱신
			if (dist[i][j] == INF) continue;
			bridgeList.push_back(make_pair(make_pair(i, j), dist[i][j]));
		}
	}
}

//선택한 다리들의 조합으로 탐색할 수 있는 섬의 개수 == 전체 섬의 개수 여부 리턴
bool checkPossible() {
	memset(connect, false, sizeof(connect));
	memset(connectIsland, false, sizeof(connectIsland));

	for (int i = 0; i < bridgeList.size(); i++) {
		if (select[i]) {
			int x = bridgeList[i].first.first;
			int y = bridgeList[i].first.second;

			// 선택한 다리가 연결하는 섬들의 연결관계를 표시
			connect[x][y] = true;
			connect[y][x] = true;
		}
	}

	// 이후 BFS탐색을 통해서 탐색할 수 있는 섬의 갯수를 세기
	queue<int> q;
	q.push(1);
	connectIsland[1] = true;

	int islandCnt = 1;
	bool flag = false;

	while (!q.empty()) {
		int current = q.front();
		q.pop();

		if (islandCnt == islandNum)	{
			flag = true;
			break;
		}

		for (int i = 1; i < islandNum + 1; i++) {
			if (current == i) {
				continue;
			}
			if (connect[current][i] && !connectIsland[i]) {
				connectIsland[i] = true;
				q.push(i);
				islandCnt++;
			}
		}
	}
	return flag;
}

void solution(int index, int count, int sum) {
	//Combination 활용해 정답 찾기
	if (count > 0) { //다리 1개 이상 뽑았으면 확인해보기
		if (checkPossible()) {
			if (answer > sum) {
				answer = sum; //최소값 갱신
			}
		}
	}

	for (int i = index; i < bridgeList.size(); i++) {
		if (select[i]) continue;
		select[i] = true;
		solution(i, count + 1, sum + bridgeList[i].second);
		select[i] = false;
	}
}

int main(void) {
	input(); //입력 받기
	label(); //섬에 번호 붙이기(1,2,3...)
	calculateDistance(); //각 섬 사이의 최단 거리 계산
	solution(0, 0, 0); //정답 계산

	if (answer == INF) {
		printf("-1");
	}
	else {
		printf("%d", answer);
	}

	return 0;
}
