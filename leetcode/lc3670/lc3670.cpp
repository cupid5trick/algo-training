#include <bits/stdc++.h>

using namespace std;

int bit_width(int x) {
    int i {};
    for (i = 31; i > 0; i --) {
        if ((x >> (i-1))) {
            break;
        }
    }
    return i;
}
class Solution {
public:
    long long maxProduct(vector<int>& nums) {
        // a_i&a_j == 0 且 a_i*a_j 最大
        int n = nums.size();
        const int&& U = bit_width(*max_element(nums.begin(), nums.end()));
        int u = 1 << U;
        // unordered_map<int, int> dp (u);
        vector<int> dp (u);

        // 把 nums[i] 看成掩码：需要满足 nums[i]&nums[j] == 0
        // 从1置位的视角：“最多”和nums[i]的补码有相同数量的1
        // 1101_0110:    0010_1001   0010_1000  0000_1001   ...
        // 因此：预处理dp[x]表示满足掩码x的最大数字
        // dp[x] = max(dp[x], dp[y]) y = x+lowbit(~x&(u-1))

        for (int i=0;i <n; i ++) {
            auto& x = nums[i];
            dp[x] = x;
        }
        for (int x=0; x < u; x ++) {
            int cx = x;

            while (cx) {
                int lb = (cx&-cx);
                int y = x^lb;
                // bitset<17> b1(1ULL*x), b2(1ULL*y);
                // cout << b1.to_string() << "\t" << b2.to_string() << endl;
                dp[x] = max(dp[x], dp[y]);
                cx ^= lb;
            }
        }
        long long ans {};
        for (auto& x : nums) {
            int y = ~x & (u -1);
            ans = max(ans, 1LL*x*dp[y]);
        }
        return ans;
    }
};