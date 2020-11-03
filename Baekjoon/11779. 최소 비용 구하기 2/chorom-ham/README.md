아래 알고리즘으로 최소 비용은 구했는데 제대로 경로 저장이 되지 않는다.

계속 디버깅해봤는데 경로 저장을 어느 시점에 어떻게 해야할 지 모르겠다.

일단 경로 저장 코드 지운 코드는 아래와 같다.

```cpp
#include <cstdio>
#include <stack>
using namespace std;
#define MAX 1001
#define INF 999999

int n, m, start, finish;
int w[MAX][MAX]; //w[i][j]: i도시에서 j도시로 가는데 드는 최소 비용

int main() {
	//w[i][j] == INF는 에지가 없는 상태. 에지가 없는 상태로 전처리
	for (int i = 0; i < MAX + 1; i++) {
		for (int j = 0; j < MAX + 1; j++) {
			if (i == j) {
				w[i][j] = 0;
			}
			w[i][j] = INF;
		}
	}

	scanf("%d", &n);
	scanf("%d", &m);
	for (int i = 0; i < m; i++) {
		int a, b, value;
		scanf("%d %d %d", &a, &b, &value);
		w[a][b] = value;
	}
	scanf("%d %d", &start, &finish);

	//다익스트라 알고리즘 사용
	int vnear, min;

	bool visited[MAX]; //이미 방문한 정점인지 확인용
	int cost[MAX]; //start정점에서 부터 다른 정점까지의 최소 비용

	//visited와 cost배열 초기화
	for (int i = 1; i < n + 1; ++i) {
		visited[i] = 0; //방문 여부 다 false로
		cost[i] = w[start][i];  //초기 경로 가중치 최소값은 입력 받은 값으로 초기화
	}
	visited[start] = 1; //시작 정점은 방문여부 = true

	//n-1번 반복
	for (int i = 1; i < n; ++i) {
		min = INF;
		//아직 방문하지 않은 정점 중에 그 정점으로 가는 최소값을 min에, 그 정점의 인덱스를 vnear에 저장
		for (int j = 1; j < n + 1; ++j) {
			if (!visited[j] && cost[j] < min) {
				min = cost[j];
				vnear = j; //거리가 가장 가까운 정점 인덱스
			}
		}
		//현재 거리가 가장 짧은 정점을 선택하고 방문 표시
		visited[vnear] = true;

		for (int j = 1; j < n + 1; ++j) {
			//정점j를 방문하지 않았고 기존 j까지의 거리보다 vnear인덱스의 노드를 거쳐가는 비용이 더 작을 경우 거리 업데이트
			if (!visited[j] && cost[vnear] + w[vnear][j] < cost[j]) {
				cost[j] = cost[vnear] + w[vnear][j];
			}
		}
	}

	printf("%d\n", cost[finish]);

	return 0;
}

```