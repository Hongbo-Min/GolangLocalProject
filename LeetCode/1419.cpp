#include <string>
#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

class Solution
{
public:
    int minNumberOfFrogs(string croakOfFrogs)
    {
        if (croakOfFrogs.size() % 5 != 0)
        {
            return -1;
        }
        int res = 0;
        int frogNum = 0;
        vector<int> cnt(4);
        unordered_map<char, int> mp = {{'c', 0}, {'r', 1}, {'o', 2}, {'a', 3}, {'k', 4}};
        for (char c : croakOfFrogs)
        {
            int t = mp[c];
            if (t == 0)
            {
                cnt[t]++;
                frogNum++;
                if (frogNum > res)
                {
                    res = frogNum;
                }
            }
            else
            {
                if (cnt[t - 1] == 0)
                {
                    return -1;
                }
                cnt[t - 1]--;
                if (t == 4)
                {
                    frogNum--;
                }
                else
                {
                    cnt[t]++;
                }
            }
        }

        if (frogNum > 0)
        {
            return -1;
        }
        return res;
    }
};

int main()
{
    Solution s;
    int res = s.minNumberOfFrogs("croakcroak");
    cout << res << endl;
    res = s.minNumberOfFrogs("crcoakroak");
    cout << res << endl;
    res = s.minNumberOfFrogs("croakcrook");
    cout << res << endl;
    return 0;
}