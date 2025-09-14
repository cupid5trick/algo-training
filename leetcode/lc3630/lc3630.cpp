#include <bits/stdc++.h>

using namespace std;

int bit_width(int x) {
    int i {};
    for (i=31; i > 0; i --) {
        if (x >> (i-1)) {
            break;
        }
    }
    return i;
}
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
        int sz = bs.size();
        int ans{};
        for (int i = sz - 1; i >= 0; i--) {
            ans = max(ans, ans ^ bs[i]);
        }
        return ans;
    }
};
class Solution {
public:
    long long maximizeXorAndXor(vector<int>& nums) {
        int n = nums.size();
        int u = (1 << n);
        int mx = *max_element(nums.begin(), nums.end());
        int sz = bit_width(mx);
        vector<int> bit_and(u), bit_xor(u);
        bit_and[0] = -1;
        // 预处理所有子集的AND、XOR
        // SUM(i*2^i) = O()
        //  
        for (int i = 0; i < n; i++) {
            auto& x = nums[i];
            int upper = (1 << i);
            // DP: 新加入x，计算AND、XOR
            for (int j = 0; j < upper; j++) {
                bit_and[upper | j] = bit_and[j] & x;
                bit_xor[upper | j] = bit_xor[j] ^ x;
            }
        }
        bit_and[0] = 0;
        // cout << bit_and[1] << endl;

        // target = XOR(A) + AND(B) + XOR(C)
        // target = AND(B) + XOR(S) + 2*XOR(A')
        // S是B集合之外的元素集、A‘表示"相加时1的个数为偶数的集合"
        // xa' = xa & ~ xor_s
        auto&& max_xor = [&](int u) -> long long {
            int _xor = bit_xor[u];
            auto bx = XorBasis(sz);
            for (int i = 0; i < n; i++) {
                auto x = nums[i] & (~_xor);
                if ((u >> i) & 1) {
                    bx.insert(x);
                }
            }
            int _xor_a = bx.max_xor();
            return _xor + 2LL * _xor_a;
        };
        long long ans{};
        for (int i = 0; i < u; i++) {
            ans = max(ans, bit_and[i] + max_xor((u - 1) ^ i));
        }
        return ans;
    }
};