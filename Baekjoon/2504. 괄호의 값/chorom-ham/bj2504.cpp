#include<iostream>
#include<stack>
#include<string>
using namespace std;

string str;
stack<char> st;

int main() {
	getline(cin, str);

	int ans = 0;
	int temp = 1;

	for (int i = 0; i < str.length(); i++) {
		if (str[i] == '(') {
			st.push(str[i]);
			temp *= 2;
		}
		else if (str[i] == '[') {
			st.push(str[i]);
			temp *= 3;
		}
		else if (str[i] == ')') {
			if (st.empty()) {
				cout << 0;
				return 0;
			}

			if (st.top() == '(') {
				if (str[i - 1] == '(') { 
					ans += temp;
				}
				st.pop();
				temp /= 2;
			}
			else {
				cout << 0;
				return 0;
			}
		}
		else if (str[i] == ']') {
			if (st.empty()) {
				cout << 0;
				return 0;
			}
			if (st.top() == '[') {
				if (str[i - 1] == '[') {
					ans += temp;
				}
				st.pop();
				temp /= 3;
			}
			else {
				cout << 0;
				return 0;
			}
		}
	}

	if (!st.empty()) {
		cout << 0;
		return 0;
	}

	cout << ans;
	return 0;
}
