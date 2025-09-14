#include <bits/stdc++.h>

using namespace std;

const static int N=26;
vector<int> uf(N);
vector<int> sz(N, 1);
void init(int n) {
    for (int i=0; i<n; i++) {
        uf[i] = i;
    }
    // itoa(uf.begin(), uf.end(), 0);
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
    if (u == v) return;
    if (sz[v] > sz[u]) swap(u, v);
    uf[v] = u;
    sz[u] += sz[v];
}
class Solution {
public:
    bool equationsPossible(vector<string>& equations) {
        int n = equations.size();
        init(N);
        // cout << find('x'-'a') << "\t" << find('z'-'a') << endl;
        for (int i = 0; i < n; i ++) {
            auto& eq = equations[i];
            if (eq[1] == '=') {
                join(eq[0]-'a', eq[3]-'a');
            }
        }
        for (int i = 0; i < n; i ++) {
            auto& eq = equations[i];
            auto&& [u, v] = make_tuple(eq[0]-'a', eq[3]-'a');
            if (eq[1] == '!' && find(u) == find(v)) {
                return false;
            }
        }
        return true;
    }
};