/**
 * 题目链接：P5689 [CSP-S2019 江西] 多叉堆 - 洛谷 | 计算机科学教育新生态: https://www.luogu.com.cn/problem/P5689
 * 题解：P5689 [CSP-S2019 江西] 多叉堆 - 洛谷 | 计算机科学教育新生态: https://www.luogu.com.cn/problem/solution/P5689
 */
#include <bits/stdc++.h>

#define N 300005
#define P 1000000007ll
using namespace std;

typedef long long ll;
int n, m, fa[N];

ll fac[N], f[N], sz[N];

ll fastpow(ll x, ll y) {
    ll mul = 1;
    for (; y; y >>= 1, x = (x * x) % P)
    {
        if (y & 1)
            mul = (mul * x) % P;
    }
    return mul;
}

void init() {
    for (int i = 1; i <= n; i++)
    {
        f[i] = sz[i] = 1, fa[i] = i;
    }
    fac[0] = fac[1] = 1;
    for (int i = 2; i <= n; i++)
    {
        fac[i] = (fac[i - 1] * i) % P;
    }
}

ll C(int x, int y) {
    return fac[x] * fastpow(fac[y], P - 2) % P * fastpow(fac[x - y], P - 2) % P;
}

// 通过并查集连接 x->y 并维护连通大小
void join(int x, int y) {
    int xx = get(x), yy = get(y);
    f[yy] = (f[yy] * C(sz[xx] + sz[yy] - 1, sz[xx]) % P * f[xx]) % P;
    fa[xx] = yy;
    sz[yy] += sz[xx];
}

int get(int x) {
    return fa[x] == x ? x : fa[x] = get(fa[x]);
}

int main() {
    scanf("%d%d", &n, &m);
    init(); // cout<<C(2,1)<<endl;
    ll ans = 0, x, y, op;
    for (int i = 1; i <= m; i++)
    {
        scanf("%lld%lld", &op, &x);
        if (op == 1)
        {
            scanf("%lld", &y);
            x = (x + ans) % n + 1, y = (y + ans) % n + 1;
            join(x, y);
            // cout<<x<<" "<<y<<" "<<f[get(y)]<<" "<<sz[get(y)]<<endl;
        }
        else
        {
            x = (x + ans) % n + 1;
            printf("%lld\n", ans = f[get(x)]);
        }
    }
    return 0;
}