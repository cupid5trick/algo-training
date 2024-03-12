#include <iostream>
#include <cstdio>
#include <cerrno>
#include <vector>
#include <string>
#include <stack>

using namespace std;


void _print_lcs(const vector<vector<int>>& dp, const vector<vector<int>>& dp2, const string& s1, const string& s2) {
	int i = s1.length() - 1;
	int j = s2.length() - 1;
	int l = dp[i+1][j+1];
	string s(l, ' ');
	
	while (i >= 0 && j >= 0) {
		if (dp2[i][j] == 1) {
			s[-- l] = s1[i];
			-- i;
			-- j;
		}
		else if (dp2[i][j] == 2) {
			-- i;
		}
		else if (dp2[i][j] == 3){
			-- j;
		}
	}
	cout << "one of the longest common subsequence: " << s << endl;
}

int _lcs(const string& s1, const string& s2) {
	int n = s1.length();
	int m = s2.length();
	
	vector<vector<int>> dp(n+1, vector<int>(m+1));
	vector<vector<int>> dp2(n, vector<int>(m));
	
	// dp[i][j] = dp[i-1][j-1] + 1, a_i = b_j
	for (int i = 1; i < n+1; ++ i) {
		for (int j = 1; j < m+1; ++ j) {
			if (s1[i] == s2[j]) {
				dp[i][j] = dp[i-1][j-1] + 1;
				dp2[i-1][j-1] = 1;
			}
			else if (dp[i-1][j] >= dp[i][j-1]) {
				dp[i][j] = dp[i-1][j];
				dp2[i-1][j-1] = 2;
			}
			else {
				dp[i][j] = dp[i][j-1];
				dp2[i-1][j-1] = 3;
			}
		}
	}
	_print_lcs(dp, dp2, s1, s2);
	return dp[n][m];
}

int lcs() {
	string s1, s2;
	cin >> s1;
	cin >> s2;
	cout << s1 <<endl;
	cout << s2 <<endl;
	return _lcs(s1, s2);
}


int main() {
	FILE* f = freopen("lcs_test.txt", "r", stdin);
	if (! f) {
		perror("Cannot open test.txt!\n");
		return errno;
	}
	
	char c;
	while ((c=getchar()) != EOF) {
		ungetc(c, stdin);
		int n = lcs();
		printf("Length = %lld\n", n);
	}
	
	
	fclose(f);
	return 0;
}
