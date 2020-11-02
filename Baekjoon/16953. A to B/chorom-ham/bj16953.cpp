#include <cstdio>

int main() {
	int a, b, count = 1;
	scanf("%d %d", &a, &b);

	//거꾸로 계산해 나가기
	while (a <= b) {
		if (a == b) {
			printf("%d", count);
			return 0;
		}
		if (b % 2 == 0) {
			b /= 2;
			count++;
		}
		else {
			if (b % 10 == 1) {
				b /= 10;
				count++;
			}
			else {
				break;
			}
		}
	}

	printf("-1");
	return 0;
}
