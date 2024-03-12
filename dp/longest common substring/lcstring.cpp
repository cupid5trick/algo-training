#include <iostream>
#include <cstdio>
#include <cerrno>
#include <vector>
#include <string>
#include <stack>

using namespace std;


class Index {
public:
	int x;
	int y;
	int l;
	
	Index() {}
	Index(int i, int j, int len): x(i), y(j), l(len) {}
};

void _print_lcstring(const vector<vector<int>>& dp, const string& s1, stack<Index>& s) {
	if (s.empty()) {
		return;
	}
	
	while(!s.empty()) {
		Index ind = s.top();
//		printf("x=%d y=%d length=%d\n", ind.x, ind.y, ind.l); 
		string str(ind.l, ' ');
		while (dp[ind.x][ind.y] > 0) {
			str[dp[ind.x][ind.y] - 1] = s1[ind.x-1];
			-- ind.x;
			-- ind.y;
		}
		printf("%s\n", str.c_str());
		s.pop();
	}
}

int _compute_lcstring(const string& s1, const string& s2) {
	stack<Index> s;
	int m = s1.length();
	int n = s2.length();
//	printf("length1=%d, length2=%d\n", s1.length(), s2.length());
	vector<vector<int>> dp(m+1, vector<int>(n+1, 0));
	
	for (int i = 1; i <= m; ++ i) {
		for (int j = 1; j <= n; ++ j) {
			if (s1[i-1] == s2[j-1]) {
				dp[i][j] = dp[i-1][j-1] + 1;
				// take record of current longest length
				while (!s.empty() && dp[i][j] > s.top().l) {
//					Index ind = s.top();
//					printf("x=%d y=%d length=%d\n", ind.x, ind.y, ind.l); 
					s.pop();
				}
				if (s.empty() || dp[i][j] == s.top().l) {
					s.push(Index(i, j, dp[i][j]));
				}
			}
		}
	}
	
//	for (int i = 0; i <= m; ++ i) {
//		for (int j = 0; j <= n; ++ j) {
//			printf("%s%d", j==0?"":" ", dp[i][j]);
//		}
//		printf("\n");
//	}
	
	int l = s.top().l;
	_print_lcstring(dp, s1, s);
	return l;
}

int lcstring() {
	string s1, s2;
	cin >> s1;
	cin >> s2;
	cout << s1 << endl;
	cout << s2 << endl;
	return _compute_lcstring(s1, s2);
}


int main() {
	FILE* f = freopen("lcstring_test.txt", "r", stdin);
	if (! f) {
		perror("Cannot open test.txt!\n");
		return errno;
	}
	
	char c;
	while ((c=getchar()) != EOF) {
		ungetc(c, stdin);
		int n = lcstring();
		printf("Length = %lld\n", n);
	}
	
	
	fclose(f);
	return 0;
}
