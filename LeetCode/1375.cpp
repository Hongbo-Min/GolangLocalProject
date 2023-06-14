#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  int numTimesAllBlue(vector<int>& flips) {
    int n = flips.size();
    if (n <= 0) {
      return 0;
    }

    int res = 0;
    int right = 0;
    for (int i = 0; i < n; ++i) {
      right = max(right, flips[i]);
      if (right == i + 1) {
        res++;
      }
    }
    return res;
  }
};

int main(int argc, char const* argv[]) {
  Solution s;
  //   vector<int> flips = {3, 2, 4, 1, 5};
  vector<int> flips = {4, 1, 2, 3};
  cout << s.numTimesAllBlue(flips) << endl;
  return 0;
}
