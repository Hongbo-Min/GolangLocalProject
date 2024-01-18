#include <iostream>

using namespace std;

class Solution
{
public:
    bool isRobotBounded(string instructions)
    {
        int direction = 0; // 0 up 1 right 2 down 3 left
        int x = 0;
        int y = 0;
        for (char c : instructions)
        {
            switch (c)
            {
            case 'R':
                direction++;
                break;

            case 'L':
                direction--;
                break;

            case 'G':
            {
                int value = abs(direction % 4);
                switch (value)
                {
                case 0:
                    y++;
                    break;

                case 1:
                    x++;
                    break;

                case 2:
                    y--;
                    break;

                case 3:
                    x--;
                    break;

                default:
                    break;
                }
            }

            default:
                break;
            }
        }
        return (direction % 4) != 0 || (x == 0 && y == 0);
    }
};

int main()
{
    Solution s;
    bool res = s.isRobotBounded("GGLLGG");
    cout << res << endl;
    return 0;
}