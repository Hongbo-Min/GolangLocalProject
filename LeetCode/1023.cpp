#include <iostream>
#include <vector>
#include <string>
using namespace std;

class Solution
{
public:
    vector<bool> camelMatch(vector<string> &queries, string pattern)
    {
        int n = queries.size();
        vector<bool> res(n, true);
        for (int i = 0; i < n; i++)
        {
            int index = 0;
            for (char c : queries[i])
            {
                if (index < pattern.size() && c == pattern[index])
                {
                    index++;
                }
                else if (c >= 'a' && c <= 'z')
                {
                    continue;
                }
                else
                {
                    res[i] = false;
                    break;
                }
            }
            if (index < pattern.size())
            {
                res[i] = false;
            }
        }

        return res;
    }
};

int main()
{
    vector<string> queries = {"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"};
    string pattern = "FB";
    Solution s;
    vector<bool> res = s.camelMatch(queries, pattern);
    for (bool val : res)
    {
        cout << val << " ";
    }
    return 0;
}