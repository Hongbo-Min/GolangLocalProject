#include <iostream>
#include <map>
#include <vector>

using namespace std;

class Solution {
 public:
  int maximalNetworkRank(int n, vector<vector<int>> &roads) {
    vector<vector<bool>> mark_edge(n, vector<bool>(n, false));
    vector<int> mark_count(n);
    for (auto road : roads) {
      mark_edge[road[0]][road[1]] = true;
      mark_edge[road[1]][road[0]] = true;
      mark_count[road[0]]++;
      mark_count[road[1]]++;
    }
    int res = 0;
    // 直接去双重遍历，暴力寻找每两个点周围的道路，如果该两点中有道路就 - 1
    for (int i = 0; i < n; ++i) {
      for (int j = i + 1; j < n; ++j) {
        int temp = mark_count[i] + mark_count[j] - mark_edge[i][j];
        res = max(res, temp);
      }
    }
    return res;
  }
};

int main() {
  Solution s;
  std::vector<std::vector<int>> roads = {{0, 1}, {0, 3}, {1, 2}, {1, 3}};
  int res = s.maximalNetworkRank(4, roads);
  std::cout << res << std::endl;
  return 0;
}