#include <cstdio>

int main() {
	int n;
	scanf("%d", &n);

	int answer = -1;
	for (int i = n / 5; i >= 0; i--) {
		if ((n - i * 5) % 2 == 0) {
			answer = i + (n - i * 5) / 2;
			break;
		}
	}

	printf("%d", answer);

	return 0;
}