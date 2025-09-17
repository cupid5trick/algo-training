#include <bits/stdc++.h>

using namespace std;

// 辗转相除法：(a, b) -> (b, a%b)
int gcd(int a, int b) {
    if (a < b) swap(a,b);
    while (b) {
        int tmp = b;
        b = a%b;
        a = tmp;
    }
    return a;
}

class Solution {
public:
    vector<int> replaceNonCoprimes(vector<int>& nums) {
        // 以 任意 顺序替换相邻的非互质数都可以得到相同的结果
        // 2 35 9 18 15 55
        // 1. “连续”的相邻“非互质数”：不断迭代(a,b) -> a/gcd(a,b)*b
        // 2. 邻项消除：直到当前位置找不到连续相邻的“非互质数”
        // 2 35 9 18 15 -> 2 35 9 90 -> ...
        int n = nums.size();
        vector<int> ans {};
        for (int i=0; i<n; i ++) {
            auto x = nums[i];
            int tmp {};
            while (!ans.empty() && (tmp=gcd(x, ans.back())) > 1) {
                // 先除后乘防止溢出
                x = x / tmp * ans.back();
                ans.pop_back();
            }
            ans.push_back(x);
        }
        return ans;
    }
};