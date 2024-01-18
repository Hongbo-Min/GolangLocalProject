#include <algorithm>
#include <iostream>
#include <vector>

class Solution {
public:
  std::vector<bool> checkArithmeticSubarrays(std::vector<int> &nums,
                                             std::vector<int> &l,
                                             std::vector<int> &r) {
    int n = nums.size();
    int m = l.size();
    std::vector<bool> result(m, true);
    for (int i = 0; i < m; ++i) {
      std::vector<int> sub_vec;
      int left = l[i];
      int right = r[i];
      for (int j = left; j <= right; ++j) {
        sub_vec.push_back(nums[j]);
      }
      std::sort(sub_vec.begin(), sub_vec.end());
      std::cout << "sub vector : " << std::endl;
      for (auto v : sub_vec) {
        std::cout << v << " ";
      }
      std::cout << std::endl;
      int sub_value = sub_vec.size() > 2 ? sub_vec[1] - sub_vec[0] : 0;
      for (int j = 1; j < right - left + 1; ++j) {
        if (sub_vec[j] - sub_vec[j - 1] != sub_value) {
          result[i] = false;
          break;
        }
      }
    }
    return result;
  }
};

int main() {
  Solution s;
  std::vector<int> nums = {4, 6, 5, 9, 3, 7};
  std::vector<int> l = {0, 0, 2};
  std::vector<int> r = {2, 3, 5};

  std::vector<bool> result = s.checkArithmeticSubarrays(nums, l, r);
  for (auto r : result) {
    std::cout << r << " ";
  }
  return 0;
}