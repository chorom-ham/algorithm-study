## 첫 시도 -- 시간 초과
```cpp
#include <cstdio>
#include <queue>
#include <utility>

#define MAX 1000

using namespace std;

int m, n, day = 0;
//정수 1은 익은 토마토, 정수 0은 익지 않은 토마토, 정수 - 1은 토마토가 들어있지 않은 칸
int arr[MAX][MAX];
queue<pair<int, int>> q;
bool visited[MAX][MAX];
int dx[4] = { 0,0,-1,1 };
int dy[4] = { 1,-1,0,0 };

int main() {
	scanf("%d %d", &m, &n);

	for (int i = 0; i < n; i++) {
		for (int j = 0; j < m; j++) {
			scanf("%d", &arr[i][j]);
			if (arr[i][j] == 1) {
				q.push({ i,j });
			}
		}
	}

	while (!q.empty()) {
		//다 익었는지 확인
		bool success = true;
		for (int i = 0; i < n; i++) {
			for (int j = 0; j < m; j++) {
				if (arr[i][j] == 0) {
					success = false;
				}
			}
		}

		if (success) {
			printf("%d", day);
			return 0;
		}

		day++;
		int size = q.size();
		for (int i = 0; i < size; i++) {
			int x = q.front().first;
			int y = q.front().second;
			q.pop();
			visited[x][y] = true;

			for (int j = 0; j < 4; j++) {
				int nx = x + dx[j];
				int ny = y + dy[j];
				if (nx >= 0 && ny >= 0 && nx < n && ny < m && arr[nx][ny] == 0 && !visited[nx][ny]) {
					arr[nx][ny] = 1;
					q.push({ nx,ny });
				}
			}
		}
	}

	printf("-1");
	return 0;
}
```

## 두번째 시도

다 익었는지 체크하는 부분을 while문 밖으로 빼서 해결했다.