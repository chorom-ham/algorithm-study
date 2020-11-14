#include <string>
#include <iostream>
#include <string>

using namespace std;

int main()
{
    int num, strnum;
    string answer = "";
    int mod, key, tmp;
    cin >> num >> strnum;
    if (strnum > num * 26 || num > strnum)
    {
        cout << "!";
        return 0;
    }
    if (num * 26 == strnum)
    {
        for (int i = 0; i < num; i++)
        {
            answer += 'Z';
        }
        cout << answer;
        return 0;
    }
    tmp = strnum - num;
    mod = tmp % 25;
    key = tmp / 25;

    for (int i = 0; i < num; i++)
    {
        answer += 'A';
    }
    int index = num - 1;
    while (key > 0)
    {
        answer[index] = 'Z';
        key--;
        index--;
    }
    answer[index] = char(65 + mod);
    cout << answer << "\n";
    return 0;
}
