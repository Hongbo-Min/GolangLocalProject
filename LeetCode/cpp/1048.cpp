#include <algorithm>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
 public:
  int longestStrChain(vector<string> &words) {
    unordered_map<string, int> mp;
    sort(words.begin(), words.end(), [](const string &s1, const string &s2) {
      return s1.size() < s2.size();
    });
    int res = 0;
    for (string word : words) {
      mp[word] = 1;
      for (int i = 0; i < word.size(); i++) {
        string subStr = word.substr(0, i) + word.substr(i + 1);
        if (mp.count(subStr)) {
          mp[word] = max(mp[word], mp[subStr] + 1);
        }
        res = max(res, mp[word]);
      }
    }
    return res;
  }
};