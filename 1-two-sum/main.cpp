#include <iostream>
#include <vector>
#include <map>
using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int, int> m;
        vector<int> r;
        for(int i = 0; i < nums.size(); i++){
            if (m.find(nums[i]) != m.end()){
                r.push_back(m[target-nums[i]]);
                r.push_back(i);
                break;
            }
            m[target - nums[i]] = i;
        }
        return r;
    }
};

int main(){
    int a[5] = {2,7,10,4,6};
    vector<int> b(a,a+5);
    Solution *s = new Solution();
    s->twoSum(b,9);
}