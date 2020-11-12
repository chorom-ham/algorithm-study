#include <cstdio>
#include <vector>
#include <utility>
#include <algorithm>
#include <queue>

using namespace std;

int n, s, t;
vector<pair<int, int>> classes;

struct compare {
	bool operator()(int a, int b) {
		return a > b;
	}
};

// 강의 끝나는 시간 저장(빨리 끝날 수록 우선순위 높음)
// pq 원소 값 : 해당 인덱스의 강의실에서 가장 마지막으로 끝나는 강의 시간
priority_queue<int, vector<int>, compare> pq; 

int main() {
	scanf("%d", &n);
	for (int i = 0; i < n; i++) {
		scanf("%d %d", &s, &t);
		classes.push_back({ s,t });
	}

	// 시작 시간 기준으로 강의 정렬
	sort(classes.begin(), classes.end());

	pq.push(classes[0].second);
	
	for (int i = 1; i < n; i++) {
		// 가장 빨리 끝나는 강의 시간 <= 현재 넣으려는 강의 시작 시간
		if (pq.top() <= classes[i].first) { 
			pq.pop();
			pq.push(classes[i].second);
		}
		else {
			pq.push(classes[i].second);
		}
	}

	printf("%d", pq.size());

	return 0;
}
