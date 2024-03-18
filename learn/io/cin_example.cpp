#include <iostream>
using namespace std;

int main()
{
    int n, q;
    std::ios::sync_with_stdio(false);
    std::cin.tie(0);
    cin >> n >> q;
    long long sum = 0;
    int cnt = 0;
    while (n-- > 0)
    {
        int a;
        cin >> a;
        if (a > 0)
        {
            sum += a;
        }
        else
        {
            cnt++;
        }
    }

    while (q-- > 0)
    {
        int low, high;
        cin >> low >> high;
        cout << sum + low * cnt << " " << sum + high * cnt << endl;
    }
}
// 64 位输出请用 printf("%lld")