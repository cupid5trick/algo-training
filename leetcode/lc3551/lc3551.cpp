#include <bits/stdc++.h>

using namespace std;

// 方法二：根据“数位和”排序，在”排序前位置“-”排序后位置“之间连边。
// 这个图上的每一个联通块恰好构成一个置换环，每个“置换环”的交换次数为k-1。
// 运行并查集维护联通块数量，答案为n-CC。
int sumd(int x)
{
    int s {};
    while (x) {
        s += x % 10;
        x /= 10;
    }
    return s;
}
class UnionFind {
public:
    vector<int> uf {};
    vector<int> sz {};
    int cnt {};
    UnionFind(int n)
    {
        cnt = n;
        uf.reserve(n);
        sz.reserve(n);
        for (int i = 0; i < n; i++) {
            uf[i] = i;
            sz[i] = 1;
        }
    }
    int find(int u)
    {
        if (uf[u] != u) {
            uf[u] = find(uf[u]);
        }
        return uf[u];
    }
    void join(int u, int v)
    {
        u = find(u);
        v = find(v);
        if (u == v)
            return;
        if (sz[v] > sz[u])
            swap(u, v);
        sz[u] += sz[v];
        uf[v] = u;
        cnt--;
    }
};

class Solution {
public:
    int minSwaps(vector<int>& nums)
    {
        int n = nums.size();
        vector<int> seq(n);
        UnionFind uf(n);
        vector<array<int, 3>> sums(n);
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            int s = sumd(x);
            sums[i] = { s, x, i };
        }
        sort(sums.begin(), sums.end());
        for (int i = 0; i < n; i++) {
            uf.join(i, sums[i][2]);
        }
        return n - uf.cnt;
    }
};
// 方法一：离散化+置换环。但是没有利用好“互不相同”；
class Solution2 {
public:
    int minSwaps(vector<int>& nums)
    {
        // 离散化
        // for (int i=0; i<n; i ++) {
        //     seq[i] = i;
        // }
        // sort(seq.begin(), seq.end(), [&](int a, int b) -> bool {
        //     int x = sumd(nums[a]), y = sumd(nums[b]);
        //     if (x == y) {
        //         return nums[a] < nums[b];
        //     }
        //     return x < y;
        // });
        // vector<bool> used(n);
        // int  ans = n;
        // for (int i=0; i<n; i ++) {
        //     int x = seq[i];
        //     if (used[x]) continue;
        //     while (!used[x]) {
        //         used[x] = true;
        //         x = seq[x];
        //     }
        //     ans --;
        // }
        // return ans;
    }
};
