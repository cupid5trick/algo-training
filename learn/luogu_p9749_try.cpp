// C++ 功底测试: https://ks.wjx.top/wjx/join/completemobile2.aspx?activityid=mFkwNRx&joinactivity=121642531453&sojumpindex=122&tvd=h%2fx8A0bqRNw%3d&costtime=4980&comsign=D3E684CFCF16BB5C14B5E845AA79D229E25361DD&wxfs=100&nw=1&jpm=27
#include <bits/stdc++.h>

using namespace std;
typedef long long ll;
typedef unsigned long long ull;


int main() {
    ll n, d;
    cin >> n >> d;
    vector<ll> v(n), a(n), f(n), g(n);
    
    for (int i=0; i <n-1; i++) {
        cin >> v[i];
        if (i > 0) {
            v[i] += v[i-1];
        }
    }
    for (int i=0; i < n; i++) {
        cin >> a[i];
    }

    f[0] = 0;
    int premin = a[0];
    ll ans = 0;
    
    for (int i=0; i < n-1; i ++) {
        if (a[i] < premin) {
            premin = a[i];
        }
        int x = (v[i]-f[i]-1)/d+1;
        f[i+1] = f[i]+x*d;
        // cout << x << " " << premin << endl;
        // 简化：累加x*premin 即可
        ans += x*premin;
    }
    
    cout << ans << endl;
    
}