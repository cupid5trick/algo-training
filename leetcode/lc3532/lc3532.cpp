#include <bits/stdc++.h>

using namespace std;

class Solution {
    vector<int> uf;
    vector<int> sz;

    void join(int u, int v) {
        u = find(u);
        v = find(v);
        if (u == v)
            return;
        if (sz[v] > sz[u]) {
            swap(v, u);
        }
        uf[v] = u;
        sz[u] += sz[v];
    }
    int find(int rt) {
        if (uf[rt] != rt) {
            uf[rt] = find(uf[rt]);
        }
        return uf[rt];
    }

public:
    Solution() = default;
    vector<bool> pathExistenceQueries(int n, vector<int>& nums, int maxDiff,
                                      vector<vector<int>>& queries) {
        uf = vector<int>(n);
        sz = vector<int>(n, 1);
        for (int i = 0; i < n; i++) {
            uf[i] = i;
        }
        for (int i = 1; i < n; i++) {
            if (abs(nums[i] - nums[i - 1]) <= maxDiff) {
                join(i, i - 1);
            }
        }
        int m = queries.size();

        vector<bool> ans(m);
        for (int i = 0; i < m; i++) {
            auto [u, v] = make_tuple(queries[i][0], queries[i][1]);
            // cout << u << "\t" << v << endl;
            ans[i] = find(u) == find(v);
        }
        return ans;
    }
};