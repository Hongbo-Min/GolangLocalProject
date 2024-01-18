#include <iostream>
#include <vector>

using namespace std;

class Solution
{
public:
    int arithmeticTriplets(vector<int> &nums, int diff)
    {
        int res = 0;
        int n = nums.size();
        if (n < 3)
        {
            return res;
        }
        for (int first_index = 0; first_index < n - 2; ++first_index)
        {
            for (int second_index = first_index + 1; second_index < n - 1; ++second_index)
            {
                if (nums[second_index] - nums[first_index] > diff)
                {
                    break;
                }
                else if (nums[second_index] - nums[first_index] == diff)
                {
                    for (int third_index = second_index + 1; third_index < n; ++third_index)
                    {
                        if (nums[third_index] - nums[second_index] > diff)
                        {
                            break;
                        }
                        else if (nums[third_index] - nums[second_index] == diff)
                        {
                            res++;
                        }
                        else
                        {
                            continue;
                        }
                    }
                }
                else
                {
                    continue;
                }
            }
        }
        return res;
    }
};

int main()
{
    Solution s;
    vector<int> nums = {0, 1, 4, 6, 7, 10};
    int res = s.arithmeticTriplets(nums, 3);
    cout << res << endl;

    return 0;
}