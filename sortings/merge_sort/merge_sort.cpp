#include <iostream>
#include <cstdio>
#include <vector>
#include <iterator>	// for prev
#include <algorithm>	// for for_each

using namespace std;


void _merge(vector<int>& v, int l, int q, int u) {
	int i = l, j = q, k;
	int n1 = q-l, n2 = u-q;
	vector<int> lv(n1), rv(n2);
	// copy
	for (k=0; i<q && j < u; ++ k) {
		lv[k] = v[i ++];
		rv[k] = v[j ++];
	}
	if (j != u) {
		rv[k ++] = v[j ++];
	}
	
	// merge
	i = 0;
	j = 0;
	for (int k = l; k < u; ++ k) {
		if ((i < n1 && j < n2 && lv[i] < rv[j]) || j >= n2) {
			v[k] = lv[i ++];
		}
		else {
			v[k] = rv[j ++];
		}
	}
}

void _merge_sort(vector<int>& v, int l, int u) {
	if (l != u-1) {
		int q = (l+u)/2;
		_merge_sort(v, l, q);
		_merge_sort(v, q, u);
		_merge(v, l, q, u);
	}
}

void merge_sort(vector<int>& v) {
	_merge_sort(v, 0, v.size());
}


int main() {
	FILE* f = freopen("../test.txt", "r", stdin);
	if (!f) {
		printf("Cannot open file!\n");
		return 1;
	}
	
    int n;
    scanf("%d\n", &n);
    while (n --) {
        int n;
		scanf("%d", &n);
		vector<int> v(n);
		for (int i = 0; i < n; ++ i) {
			scanf("%d", &v[i]);
		}
		getchar();
		for_each(v.begin(), v.end(), [](const int& x) {
			printf("%d ", x);
		});
		printf("\b\n");
		merge_sort(v);
		for_each(v.begin(), v.end(), [](const int& x) {
			printf("%d ", x);
		});
		printf("\b\n");
		puts("=================================");
    }
	
	fclose(f);
	return 0;
}
