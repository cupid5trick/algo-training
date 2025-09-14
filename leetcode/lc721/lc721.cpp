#include <bits/stdc++.h>

using namespace std;

const static int N = 1e3 + 1;
vector<int> uf(N);
vector<int> sz(N, 1);
void init(int n) {
    for (int i = 0; i < n; i++) {
        uf[i] = i;
    }
}
int find(int u) {
    if (uf[u] != u) {
        uf[u] = find(uf[u]);
    }
    return uf[u];
}
void join(int u, int v) {
    u = find(u);
    v = find(v);
    if (u == v)
        return;
    if (sz[v] > sz[u])
        swap(u, v);
    uf[v] = u;
    sz[u] += sz[v];
}
class Solution {

public:
    vector<vector<string>> accountsMerge(vector<vector<string>>& ety) {
        unordered_map<string, int> ids{};
        int n = ety.size();
        init(n);
        // 一个邮箱一定对应唯一一个帐户名
        for (int j = 0; j < n; j++) {
            auto elem = ety[j];
            int m = elem.size();
            for (int i = 1; i < m; i++) {
                string& mail = elem[i];
                // 合并实体组 (stk.back(), j)
                if (ids.find(mail) != ids.end()) {
                    join(ids[mail], j);
                } else {
                    ids[mail] = j;
                }
            }
        }
        map<int, vector<string>> ans{};
        for (auto& [mail, i] : ids) {
            int m = ety[i].size();
            auto& elem = ety[i];
            auto& st = ans[find(i)];
            // for (int j = 1; j < m; j ++) {
            //     st.insert(elem[j]);
            // }
            st.push_back(mail);
        }
        // 收集答案 (帐户名，邮箱列表)
        vector<vector<string>> _ans{};
        _ans.reserve(ans.size());
        for (auto& p : ans) {
            auto& mails = p.second;
            int m = mails.size() + 1;
            auto& vec = _ans.emplace_back();
            // 学习使用 reserve 减少insert过程中的内存分配
            // 适用于需要逐步插入但是提前预知所需内存大小的情况
            vec.reserve(m);
            vec.push_back(ety[p.first][0]);
            sort(mails.begin(), mails.end());
            vec.insert(vec.end(), mails.begin(), mails.end());
            
            // vec.insert(vec.begin(), ety[p.first][0]);
        }
        return _ans;
    }
};