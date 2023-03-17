#include <iostream>
#include <vector>
#include <algorithm>

class Solution
{
public:
    std::vector<int> answerQueries(std::vector<int> &nums, std::vector<int> &queries)
    {
        int n = nums.size();
        int m = queries.size();
        std::vector<int> pre_sum(n + 1); // nums的前缀和
        sort(nums.begin(), nums.end());
        for (int i = 0; i < n; i++)
        {
            pre_sum[i + 1] = pre_sum[i] + nums[i];
        }

        std::vector<int> res(m);
        for (int i = 0; i < m; ++i)
        {
            res[i] = std::upper_bound(pre_sum.begin(), pre_sum.end(), queries[i]) - pre_sum.begin() - 1;
        }

        return res;
    }
};

int main()
{
    Solution s;
    std::vector<int> nums = {4, 5, 2, 1};
    std::vector<int> queries = {3, 10, 21};
    std::vector<int> res = s.answerQueries(nums, queries);
    for (int r : res)
    {
        std::cout << r << std::endl;
    }
    return 0;
}