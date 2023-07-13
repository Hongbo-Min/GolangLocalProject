#include <algorithm>
#include <iostream>
#include <vector>

class Solution {
 public:
  int minFallingPathSum(std::vector<std::vector<int>>& matrix) {
    int n = matrix.size();
    int result = INT32_MAX;
    if (n == 1) return matrix[0][0];
    for (int i = 1; i < n; ++i) {
      for (int j = 0; j < n; j++) {
        if (j == 0) {
          matrix[i][j] = std::min(matrix[i - 1][j], matrix[i - 1][j + 1]) + matrix[i][j];
        } else if (j == n - 1) {
          matrix[i][j] = std::min(matrix[i - 1][j], matrix[i - 1][j - 1]) + matrix[i][j];
        } else {
          matrix[i][j] =
              std::min(std::min(matrix[i - 1][j - 1], matrix[i - 1][j]), matrix[i - 1][j + 1]) + matrix[i][j];
        }
        if (i == n - 1) {
          result = std::min(result, matrix[i][j]);
        }
      }
    }
    return result;
  }
};

int main() {
  Solution s;
  std::vector<std::vector<int>> vec{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}};
  int res = s.minFallingPathSum(vec);
  std::cout << res << std::endl;
  return 0;
}