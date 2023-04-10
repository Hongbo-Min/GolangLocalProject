#include <iostream>
#include <vector>
#include <stack>
using namespace std;

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */

struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution
{
public:
    vector<int> nextLargerNodes(ListNode *head)
    {
        vector<int> res;
        stack<pair<ListNode *, int>> stk;
        ListNode *cur = head;
        int idx = -1;
        while (cur)
        {
            idx++;
            res.push_back(0);
            while (!stk.empty() && stk.top().first->val < cur->val)
            {
                res[stk.top().second] = cur->val;
                stk.pop();
            }
            stk.emplace(cur, idx);
            cur = cur->next;
        }
        return res;
    }
};

int main()
{
    Solution s;

    ListNode *head = new ListNode(2);
    ListNode *temp = head;
    temp->next = new ListNode(1);
    temp = temp->next;
    temp->next = new ListNode(5);
    temp = temp->next;

    vector<int> res = s.nextLargerNodes(head);
    for (auto value : res)
    {
        cout << value << " ";
    }
    return 0;
}