#include <iostream>
#include <cstdio>
#include <cerrno>
#include <vector>
#include <string>
#include <stack>
#include <list>

using namespace std;

const int INF = 0x7fffffff;

int shortest_path(const vector<vector<int>>& aj, const vector<list<int>>& stages) {
	int k = stages.size();
	int n = aj.size();
	vector<int> dp(n, INF);
	vector<int> dp2(n);
	dp[0] = 0;
	
	
	// dp
	// dp[i] = min{dp[j] + w_{ji} | j \in V_{k_i+1}}
	for (int ki = 1; ki < k; ++ ki) {
		// the previous layer
		for (auto p1 = stages[ki].begin(); p1 != stages[ki].end(); ++ p1) {
			int i = *p1;
			// next layer
			for (auto p2 = stages[ki-1].begin(); p2 != stages[ki-1].end(); ++ p2) {
				int j =*p2;
				if (aj[j][i] < 0) {
					continue;
				}
				if (dp[j] + aj[j][i] < dp[i]) {
					dp[i] = dp[j] + aj[j][i];
					dp2[i] = j;
				}
			}
		}
		
	}
	
	// reconstruct the path
	stack<int> p;
	int v = n-1;
	while (v != 0) {
		p.push(v);
		v = dp2[v];
	}
	printf("The shortest path: %d", 0);
	while (!p.empty()) {
		printf(" -> %d", p.top());
		p.pop();
	}
	printf("\n");
	return dp[n-1];
}

int multistages_graph() {
	int n, m, k;
	scanf("%d %d %d\n", &n, &m, &k);
	
	
	vector<list<int>> stages(k);
	vector<vector<int>> aj(n, vector<int>(n, -1));
	
	for (int i = 0; i < n; ++ i) {
		aj[i][i] = 0;
	}
	
	// input
	for (int i = 0; i < m; ++ i) {
		int k1, k2, i1, i2, w12;
		scanf("%d %d %d %d %d\n", &k1, &k2, &i1, &i2, &w12);
		aj[i1][i2] = w12;
		stages[k1].push_back(i1);
		stages[k2].push_back(i2);
	}
	
	return shortest_path(aj, stages);
}


int main() {
	FILE* f = freopen("multistages_graph_test.txt", "r", stdin);
	if (! f) {
		perror("Cannot open test.txt!\n");
		return errno;
	}
	
	char c;
	while ((c=getchar()) != EOF) {
		ungetc(c, stdin);
		int n = multistages_graph();
		printf("Length = %lld\n", n);
	}
	
	
	fclose(f);
	return 0;
}
