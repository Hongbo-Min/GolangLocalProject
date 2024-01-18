#include <algorithm>
#include <iostream>
#include <map>
#include <vector>

using namespace std;

class Solution {
 public:
  int bestTeamScore(vector<int> &scores, vector<int> &ages) {
    int n = scores.size();
    vector<pair<int, int>> people;
    for (int i = 0; i < n; ++i) {
      people.push_back({scores[i], ages[i]});
    }
    sort(people.begin(), people.end());
    vector<int> dp(n, 0);
    int res = 0;
    for (int i = 0; i < n; ++i) {
      for (int j = i - 1; j >= 0; --j) {
        if (people[j].second <= people[i].second) {
          dp[i] = max(dp[i], dp[j]);
        }
      }
      dp[i] += people[i].first;
      res = max(res, dp[i]);
    }
    return res;
  }
};

int main() {
  Solution s;
  std::vector<int> scores = {1, 1, 1, 1, 1, 1, 1, 1, 1, 1};
  std::vector<int> ages = {811, 364, 124, 873, 790, 656, 581, 446, 885, 134};
  int res = s.bestTeamScore(scores, ages);
  std::cout << res << std::endl;
  return 0;
}