#include<iostream>
#include<string>
using namespace std;

string s;

int main() {
	getline(cin, s);

	int count = 0;
	for (int i = 0; i < s.length(); i++) {
		if (s[i] == 'P') {
			count++;
			continue;
		}
		//앞에 P가 2번 이상 나오고, 지금 A이며 다음이 P라면
		if (s[i + 1] == 'P' && count >= 2) {
			// PPAP는 P로 치환(처리해준 애들은 없어졌다고 생각)
			count--; 
			// 뒤에 P까지 확인했으므로 i++ 
			i++;
		}
		else{
			cout << "NP";
			return 0;
		}
	}

	if (count == 1) {
		cout << "PPAP";
	}
	else {
		cout << "NP";
	}
	
	return 0;
}