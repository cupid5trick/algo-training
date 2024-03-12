#include <iostream>
#include <cstdio>
#include <cerrno>
#include <vector>
#include <string>
#include <stack>

using namespace std;


int _max_sum(const vector<int>& A) {
	int n = A.size();
	int l, u, m = 0x80000000;
	vector<vector<int>> dp(n, vector<int>(n));
	
	for (int i = 0; i < n; ++ i) {
		dp[i][i] = A[i];
	}
	
	// dp
	for (int i = 1; i < n-1; ++ i) {
		for (int j = i+1; j < n; ++ j) {
			dp[i][j] = dp[i][j-1] + A[j];
			if (dp[i][j] > m) {
				m = dp[i][j];
				l = i;
				u = j+1;
			}
		}
	}
	
	// output result
	printf("Max sum: %d =  %d", m, A[l]);
	for (int i = l+1; i < u; ++ i) {
		printf(" + %d", A[i]);
	}
	printf("\n");
	return m;
}

int max_sum() {
	int n;
	scanf("%d", &n);
	vector<int> v(n);
	for (int i = 0; i < n; ++ i) {
		scanf("%d", &v[i]);
	}
	
	return _max_sum(v);
}


int main() {
	FILE* f = freopen("maxsum_test.txt", "r", stdin);
	if (! f) {
		perror("Cannot open test.txt!\n");
		return errno;
	}
	
	char c;
	while ((c=getchar()) != EOF) {
		ungetc(c, stdin);
		int n = max_sum();
		printf("maxsum = %lld\n", n);
	}
	
	
	fclose(f);
	return 0;
}
