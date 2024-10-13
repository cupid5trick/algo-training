/**
 * 题目链接：P9749 [CSP-J 2023] 公路 - 洛谷 | 计算机科学教育新生态: https://www.luogu.com.cn/problem/P9749
 * 题解: https://www.luogu.com.cn/problem/solution/P9749
 */
#include <bits/stdc++.h>
using namespace std;


int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);

    long long n, d;
    cin >> n >> d;

    vector<long long> v(n);
    vector<long long> a(n + 1);
    vector<long long> f(n + 1);
    vector<long long> premin(n + 1);
    vector<long long> remain(n + 1);

    premin[0] = 1e18;

    for (long long i = 1; i < n; i++) {
        cin >> v[i];
    }

    for (long long i = 1; i <= n; i++) {
        cin >> a[i];
        premin[i] = min(premin[i - 1], a[i]);
    }

    for (long long i = 2; i <= n; i++) {
        long long ceilValue = (v[i - 1] - remain[i - 1] + d - 1) / d;
        f[i] = f[i - 1] + ceilValue * premin[i - 1];
        remain[i] = ceilValue * d - (v[i - 1] - remain[i - 1]);
    }

    cout << f[n] << endl;

    return 0;
}