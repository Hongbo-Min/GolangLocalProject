#include <iostream>
#include <vector>
#include <map>
using namespace std;

class Solution
{
public:
    int mostFrequentEven(vector<int> &nums)
    {
        map<int, int, less<int>> m;
        for (auto num : nums)
        {
            if (num % 2 == 0)
            {
                m[num]++;
            }
        }
        int value = -1;
        int number = 0;
        for (auto it : m)
        {
            if (it.first % 2 == 0)
            {
                if (it.second > number)
                {
                    value = it.first;
                    number = it.second;
                }
            }
        }
        return value;
    }
};

int main()
{
    vector<int> nums = {29, 47, 21, 41, 13, 37, 25, 7};
    Solution s;
    int res = s.mostFrequentEven(nums);
    cout << res << endl;
    return 0;
}