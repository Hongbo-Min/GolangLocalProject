#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

class Solution
{
public:
    vector<int> numMovesStonesII(vector<int> &stones)
    {
        int res_min = 0;
        int res_max = 0;
        int length = stones.size();
        sort(stones.begin(), stones.end());
        if (stones[length - 1] - stones[0] == length - 1)
        {
            return {res_min, res_max};
        }

        res_max = max((stones[length - 1] - stones[1] + 1), (stones[length - 2] - stones[0] + 1)) - (length - 1);
        res_min = length;
        for (int i = 0, j = 0; i < length && j + 1 < length; ++i)
        {
            while (j + 1 < length && stones[j + 1] - stones[i] + 1 <= length)
            {
                ++j;
            }
            if (j - i + 1 == length - 1 && stones[j] - stones[i] + 1 == length - 1)
            {
                res_min = min(res_min, 2);
            }
            else
            {
                res_min = min(res_min, length - (j - i + 1));
            }
        }

        return {res_min, res_max};
    }
};

int main()
{
    Solution s;
    vector<int> stones = {7, 4, 9}; // {6,5,4,3,10} {100,101,104,102,103}
    vector<int> res = s.numMovesStonesII(stones);
    for (size_t i = 0; i < res.size(); i++)
    {
        cout << res[i] << " ";
    }
    return 0;
}