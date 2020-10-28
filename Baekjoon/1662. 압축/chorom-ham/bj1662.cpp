#include <iostream>
#include <string>
#include <stack>
using namespace std;

string s;
int visited[50];
int ans;

//바깥쪽에서 안쪽으로 계산해 나감
int calc(int start, int end) {
	int result = 0; 
	for (int i = start; i < end; i++) { 
		if (s[i] == '(') { 
			int k = s[i - 1] - '0'; 
			result += k * calc(i + 1, visited[i]) - 1;

			i = visited[i]; //처리한 문자열 이후부터 다시 탐색
			continue; 
		} 
		result++; //괄호 안에 있지 않은 문자들 처리
	} 
	return result;
}

int main() {
	getline(cin, s);

	stack<int> st; 

	// 괄호 쌍 위치 전처리
	// 전처리 안해주면 3(972(11)13(2))22처럼 같은 층위의 괄호 안에 여러 괄호쌍이 나오는 문자열 처리 불가
	for (int i = 0; i < s.length(); i++) { 
		if (s[i] == '(') { 
			st.push(i); 
		} else if (s[i] == ')') { 
			visited[st.top()] = i;
			st.pop(); 
		} 
	}

	cout << calc(0, s.length());
	return 0;
}
