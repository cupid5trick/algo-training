#include <iostream>
#include <cstdio>
#include <cerrno>
#include <vector>
#include <string>
#include <stack>

using namespace std;


void _print_mcp() {
	
}

// actually compute the max times of multiply
int _mcp(const vector<int>& p) {
	int n = p.size() - 1;
	int INF = 0x7fffffff;
	// dp[i][i] = 0
	vector<vector<int>> dp(n, vector<int>(n, INF));
	vector<vector<int>> dp2(n, vector<int>(n, 0));
	for (int i = 0; i < n; ++ i) {
		dp[i][i] = 0;
	}
	
	// dp
	for (int k = 1; k < n; ++ k) {
		for (int i = 0; i < n - k; ++ i) {
			// j \in [i, i+k)
			for (int j = i; j < i+k; ++ j) {
				if (dp[i][j] + dp[j+1][i+k] + p[i]*p[j+1]*p[i+k+1] < dp[i][i+k]) {
					dp[i][i+k] = dp[i][j] + dp[j+1][i+k] + p[i]*p[j+1]*p[i+k+1];
					dp2[i][i+k] = j;
				}
			}
		}
	}
	
	printf("N(1,n)=%d\n", dp[0][n-1]);
	return dp[0][n-1];
}

int mcp() {
	int n;
	// the number of matrices
	scanf("%d", &n);
	vector<int> p(n+1);
	// the dimensions
	for (int i = 0; i <= n; ++ i) {
		scanf("%d", &p[i]);
	}
	
	return _mcp(p);
}


int main() {
	FILE* f = freopen("mcp_test.txt", "r", stdin);
	if (! f) {
		perror("Cannot open test.txt!\n");
		return errno;
	}
	
	char c;
	while ((c=getchar()) != EOF) {
		ungetc(c, stdin);
		int n = mcp();
		printf("Length = %lld\n", n);
	}
	
	
	fclose(f);
	return 0;
}
