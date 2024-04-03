#include <iostream>
#include <cstdio>
#include <vector>
#include <algorithm> // for for_each

using namespace std;

int _partition(vector<int> &v, int l, int u)
{
	int i = -1, j;
	int x = v[u - 1];
	for (j = 0; j < u - 1; ++j)
	{
		if (v[j] < x && j != ++i)
		{
			int t = v[i];
			v[i] = v[j];
			v[j] = t;
		}
	}
	v[u - 1] = v[++i];
	v[i] = x;

	return i;
}

void _quick_sort(vector<int> &v, int l, int u)
{
	if (l != u - 1)
	{
		int q = _partition(v, l, u);
		//		printf("q=%d A[q]=%d\n", q, v[q]);
		if (q > 0)
		{
			_quick_sort(v, l, q);
		}
		if (q + 1 < u)
		{
			_quick_sort(v, q + 1, u);
		}
	}
}

void quick_sort(vector<int> &v)
{
	_quick_sort(v, 0, v.size());
}

int main()
{
	FILE *f = freopen("../test.txt", "r", stdin);
	if (!f)
	{
		printf("Cannot open file!\n");
		return 1;
	}

	int n;
	scanf("%d\n", &n);
	while (n--)
	{
		int n;
		scanf("%d", &n);
		vector<int> v(n);
		for (int i = 0; i < n; ++i)
		{
			scanf("%d", &v[i]);
		}
		getchar();
		for_each(v.begin(), v.end(), [](const int &x)
				 { printf("%d ", x); });
		printf("\b\n");
		quick_sort(v);
		for_each(v.begin(), v.end(), [](const int &x)
				 { printf("%d ", x); });
		printf("\b\n");
		puts("=================================");
	}

	fclose(f);
	return 0;
}
