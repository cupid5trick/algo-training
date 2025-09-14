#include <bits/stdc++.h>

using namespace std;

class XorBasis {
    vector<int> bs;

public:
    XorBasis(int n) { bs = vector<int>(n); }
    void insert(int x) {
        int sz = bs.size();
        for (int i = sz - 1; i >= 0; i--) {
            if ((x >> i) & 1) {
                if (bs[i] == 0) {
                    bs[i] = x;
                    return;
                }
                x ^= bs[i];
            }
        }
    }
    int max_xor() {
        int ans{};
        // 贪心：如果异或能变大，就加入x这个基
        int sz = bs.size();
        for (int i = sz - 1; i >= 0; i--) {
            auto& x = bs[i];
            ans = max(ans ^ x, ans);
        }
        return ans;
    }
};

class Solution {
public:
    int maxXorSubsequences(vector<int>& nums) {
        // f[i+1] = max(f[j] ^ nums[i], f[j]), j <= i
        // f[0] = 0
        // 线性基
        XorBasis bx(31);
        for (auto& x : nums) {
            bx.insert(x);
        }
        return bx.max_xor();
    }
};