#include <bits/stdc++.h>

using namespace std;

const int mod = 1e9+7;
class Solution {
public:
    int peopleAwareOfSecret(int n, int d, int fg) {
        deque<pair<int,int>> know {}, share {};
        // the first day, only 1 people know 
        know.emplace_back(0, 1);
        // day-know[0].first >= d: 变成 share 分享者
        // [t+d, t+fg-1]: share.first+d <= day <= share.first+fg-1 生产等量的 know 知情者
        int know_cnt = 1, share_cnt {};
        for (int day=1; day<n; day++) {
            if (!know.empty() && day >= know[0].first + d) {
                // 加入 share 尾部，则share头部是最“老”的人
                share.emplace_back(know[0].first, know[0].second);
                share_cnt = (share_cnt + know[0].second) % mod;
                know_cnt = (know_cnt - know[0].second + mod) % mod;
                know.pop_front();
            }
            while (!share.empty() && (share[0].first+fg-1 < day)) {
                share_cnt = (share_cnt - share[0].second + mod) % mod;
                share.pop_front();
            }
            if (!share.empty() && share[0].first+fg-1 >= day) {
                know.emplace_back(day, share_cnt);
                know_cnt = (know_cnt + share_cnt) % mod;
            }
        }

        return (know_cnt + share_cnt) % mod;
    }
};