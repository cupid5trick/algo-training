#include <bits/stdc++.h>

using namespace std;

template<typename R,typename T>
void debug(const unordered_map<R,T>& nums) {
    for (auto& [k, v]: nums) {
        cout << k << ": " << v << " " ;
    }
    cout << endl;
}
class Solution {
public:
    int minArrivalsToDiscard(vector<int>& ts, int w, int m) {
        int n = ts.size();
        unordered_map<int,int> cnt(n);
        vector<bool> drop(n);
        int mx {}, minor {};
        int ans = 0;
        for (int i=0; i < n; i ++) {
            auto t = ts[i];
            // if (i>=w && t == ts[i-w]) {continue;}
            if (i >= w && !drop[i-w]) {
                auto tmp = cnt[ts[i-w]];
                cnt[ts[i-w]] --;
                if (tmp == mx && tmp > minor) {
                    mx --;
                }
            }
            // 否则丢弃
            if (cnt[t] < m) {
                // cout << i << "\t"; debug(cnt);
                cnt[t] ++;
            } else {
                ans ++;
                drop[i] = true;
                continue;
            }
            if (cnt[t] > mx) {
                mx = cnt[t];
            }
            else if (cnt[t] >= minor) {
                minor = cnt[t];
            }
        }
        return ans;
    }
};
