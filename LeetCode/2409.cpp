#include <iostream>
#include <vector>
#include <string>
#include <numeric>
#include <math.h>
#include <algorithm>
using namespace std;

constexpr int day[]{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};

class Solution
{
public:
    int countDaysTogether(string arriveAlice, string leaveAlice, string arriveBob, string leaveBob)
    {
        auto getDay = [](auto &s)
        { return std::accumulate(day, day + 10 * (s[0] - '0') + (s[1] - '0') - 1, 10 * (s[3] - '0') + (s[4] - '0')); };
        int aa = getDay(arriveAlice);
        int la = getDay(leaveAlice);
        int ab = getDay(arriveBob);
        int lb = getDay(leaveBob);
        return std::max(0, std::min({la - aa, lb - ab, la - ab, lb - aa}) + 1);
    }
};

int main()
{
    string arriveAlice = "08-15";
    string leaveAlice = "08-18";
    string arriveBob = "08-16";
    string leaveBob = "08-19";
    Solution s;
    int res = s.countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob);
    cout << res << endl;
    return 0;
}