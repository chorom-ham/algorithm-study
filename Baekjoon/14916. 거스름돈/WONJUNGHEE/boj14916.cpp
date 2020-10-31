#include <iostream>
#include <sstream>

using namespace std;

int N;

int main()
{
    cin >> N;
    int money = N;
    int count = 0;
    if (money <= 1 or money == 3)
    {
        count = -1;
        money = 0;
    }
    while (money != 0)
    {
        if (money >= 5)
        {
            money -= 5;
            count += 1;
        }
        else if (money <= 4 and money >= 2)
        {
            money -= 2;
            count += 1;
        }
        else
        {
            money += 3;
        }
    }

    cout << count;

    return 0;
}
