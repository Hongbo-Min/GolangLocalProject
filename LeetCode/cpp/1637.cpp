#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

bool MyCompare(vector<int> &v1, vector<int> &v2)
{
    return v1[0] > v2[0];
}

class Solution
{
public:
    int maxWidthOfVerticalArea(vector<vector<int>> &points)
    {
        sort(points.begin(), points.end(), MyCompare);
        int result = 0;
        for (int i = 0; i < points.size() - 1; ++i)
        {
            result = max(result, points[i][0] - points[i + 1][0]);
        }
        return result;
    }
};

int main()
{
    Solution s;
    vector<vector<int>> points = {
        {3, 1},
        {9, 0},
        {1, 0},
        {1, 4},
        {5, 3},
        {8, 8}};
    int res = s.maxWidthOfVerticalArea(points);
    cout << res << endl;
    return 0;
}