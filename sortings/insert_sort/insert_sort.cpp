#include <iostream>
#include <cstdio>
#include <vector>
#include <iterator>	// for prev
#include <algorithm>	// for for_each
#include <cerrno>       // for errno

using namespace std;

void insert_sort(vector<int>& v) {
	for (auto it = next(v.begin()); it != v.end(); ++ it) {
		auto j = prev(it);
		int key = *it;
		while (j != prev(v.begin()) && *j > key) {
			*next(j) = *j;
			-- j;
		}
		*next(j) = key;
	}
}


int main() {
    FILE* f = freopen("../test.txt", "r", stdin);
    if (!f) {
        perror("Cannot open file!\n");
        return errno;
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
		insert_sort(v);
		for_each(v.begin(), v.end(), [](const int& x) {
			printf("%d ", x);
		});
		printf("\b\n");
		puts("=================================");
	}

	fclose(f);
	return 0;
}
