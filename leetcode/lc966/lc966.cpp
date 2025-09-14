#include <bits/stdc++.h>

using namespace std;

unordered_set<char> voccel = {'a', 'e', 'i', 'o', 'u'};
string uppercase(string& str) {
    string tmp(str);
    for (auto& c : tmp) {
        if (c < 'a' && c != '-') {
            c += 'a' - 'A';
        }
    }
    return tmp;
}
string rmv_voccel(string& str) {
    string tmp((str));
    for (auto& c : (tmp)) {
        if (voccel.find(c) != voccel.end()) {
            c = '-';
        }
    }
    return tmp;
}
class Solution {
public:
    vector<string> spellchecker(vector<string>& wl, vector<string>& qs) {
        int n = wl.size();
        // 第一次出现位置
        unordered_set<string> table(wl.begin(), wl.end());
        // 2. 转化为小写
        unordered_map<string, string> st(n);
        // 3. 等价于相同下标的辅音字母一一对应
        unordered_map<string, string> tb3(n);

        for (int i = 0; i < n; i++) {
            auto& str = wl[i];
            string lower = uppercase(str);
            if (st.find(lower) == st.end())
                st[lower] = ((str));
            string tmp = rmv_voccel(lower);
            if (tb3.find(tmp) == tb3.end())
                tb3[(tmp)] = ((str));
            // cout << (tmp) << "\t" << str << endl;
        }
        int m = qs.size();
        vector<string> ans(m);
        for (int i = 0; i < m; i++) {
            auto& q = qs[i];
            auto qk = uppercase(q);
            if (table.find(q) != table.end()) {
                ans[i] = q;
                continue;
            }
            if (!st.empty() && st.find(qk) != st.end()) {
                ans[i] = st[qk];
                continue;
            }
            string tmp = rmv_voccel(qk);
            if (!tb3.empty() && tb3.find((tmp)) != tb3.end()) {
                ans[i] = tb3[(tmp)];
            }
        }
        return ans;
    }
};