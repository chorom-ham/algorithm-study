#include <string>
#include <iostream>

using namespace std;

int main()
{
    string str;
    cin >> str;

    int p = 0;
    for (int i = 0; i < str.length(); i++)
    {
        if (str[i] == 'P')
        {
            p++;
            continue;
        }
        if (p >= 2 && str[i + 1] == 'P')
        {
            p--;
            i++;
        }
        else
        {
            cout << "NP";
            return 0;
        }
    }

    if (p == 1)
    {
        cout << "PPAP";
    }
    else
    {
        cout << "NP";
    }

    return 0;
}
