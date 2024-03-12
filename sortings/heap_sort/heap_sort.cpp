#include <iostream>
#include <cstdio>
#include <vector>
#include <iterator>	// for prev
#include <algorithm>	// for for_each

using namespace std;


void _maxify_heap(vector<int>& v, int i, int n) {
	int l=2*i+1, r=2*i+2, largest=i;
	if (l < n && v[l] > v[i]) {
		largest = l;
	}
	if (r < n && v[r] > v[largest]) {
		largest = r;
	}
	if (i != largest) {
		int t = v[i];
		v[i] = v[largest];
		v[largest] = t;
		_maxify_heap(v, largest, n);
	}
}

void build_max_heap(vector<int>& v, const int n) {
	for (int i=n/2-1; i>=0; -- i) {
		_maxify_heap(v, i, n);
	}
}

void heap_sort(vector<int>& v) {
	int n = v.size();
	build_max_heap(v, n);
	for (int i=n-1; i>=0; -- i) {
		int t = v[i];
		v[i] = v[0];
		v[0] = t;
		_maxify_heap(v, 0, -- n);
	}
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
		heap_sort(v);
		for_each(v.begin(), v.end(), [](const int& x) {
			printf("%d ", x);
		});
		printf("\b\n");
		puts("=================================");
    }
	
	fclose(f);
	return 0;
}
