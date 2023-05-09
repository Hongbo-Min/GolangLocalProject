#include <iostream>
#include <string>

class Solution
{
public:
    int countTime(std::string time)
    {
        int count = 0;
        for (char c : time)
        {
            if (c == '?')
            {
                count++;
            }
        }
        if (count == 0)
        {
            return 0;
        }
        int res = 1;

        char h_first = time[0];
        char h_second = time[1];
        char m_first = time[3];
        char m_second = time[4];

        if (m_second == '?')
        {
            res *= 10;
        }
        if (m_first == '?')
        {
            res *= 6;
        }
        if (h_first == '?' && h_second == '?')
        {
            res *= 24;
        }
        else if (h_second == '?')
        {
            if (h_first < '2')
            {
                res *= 10;
            }
            else
            {
                res *= 4;
            }
        }
        else if (h_first == '?')
        {
            if (h_second >= '4')
            {
                res *= 2;
            }
            else
            {
                res *= 3;
            }
        }
        return res;
    }
};

int main()
{
    Solution s;
    std::cout << s.countTime("?5:00") << std::endl;
    std::cout << s.countTime("0?:0?") << std::endl;
    std::cout << s.countTime("??:??") << std::endl;

    return 0;
}