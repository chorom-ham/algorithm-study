#include <cstdio>
#include <string>

using namespace std;

int length, value;
string ans;

int main() {
	scanf("%d %d", &length, &value);
	
	if (value<length || value>length * 26) {
		printf("!");
		return 0;
	}

	for (int i = 1; i < length + 1; i++) {
		int tmp = value - (length - i) * 26;
		if (tmp < 1) {
			ans += "A";
			value--;
		}
		else{
			ans += char(tmp + 64);
			value -= tmp;
		}
	}

	printf("%s", ans.c_str());

	return 0;
}
