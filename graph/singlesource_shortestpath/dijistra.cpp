#include <iostream>
#include <cstdio>
#include <vector>
#include <list>
#include <iterator>	// for prev
#include <algorithm>	// for for_each
#include <cerrno>       // for errno

using namespace std;


using Path=list<int>;

void dijistra() {
	int n, m;
	scanf("%d %d\n", &n, &m);
	
	const int INF = 0x7fffffff;
	vector<vector<int>> aj(n, vector<int>(n, INF));
	vector<string> nodes(n);
	for (int i = 0; i < n; ++ i) {
		string s;
		cin >> s;
		nodes[i] = s;
	}
	
//	for_each(nodes.begin(), nodes.end(), [](const string& s) {
//		cout << s <<endl;
//	});
	
	for (int i = 0; i < m; ++ i) {
		int u, v, w;
		scanf("%d %d %d\n", &u, &v, &w);
		aj[u][v] = w;
	}
	
	vector<int> d(n, INF);
	vector<Path> p(n);
	vector<bool> visited(n, false);
	d[0] = 0;
	
	// dijistra
	for (int i = 0; visited[i] != true;) {
//		printf("i=%s\n", nodes[i].c_str());
		visited[i] = true;
		for (int j = 0; j < n; ++ j) {
			if (aj[i][j] == INF) {
				continue;
			}
			if (d[j] > d[i] + aj[i][j]) {
				d[j] = d[i] + aj[i][j];
				p[j] = p[i];
				p[j].push_back(j);
			}
		}
		
//		for_each(d.begin(), d.end(), [](const int x) {
//			printf(" %d", x);
//		});
//		printf("\n");
		
		// i = argmin{d[i]}
		int min = INF;
		for (int j = 0; j < n; ++ j) {
			if (d[j] == INF || visited[j]) {
				continue;
			}
			if (d[j] < min) {
				min = d[j];
				i = j;
			}
		}
	}
	
	// output
	for (int i = 0; i < n; ++ i) {
		printf("distance from %s to %s: %d\n%s", nodes[0].c_str(), nodes[i].c_str(), d[i], nodes[0].c_str());
		for (auto it = p[i].begin(); it != p[i].end(); ++ it) {
			printf("->%s", nodes[*it].c_str());
		}
		printf("\n");
	}
}


int main() {
//    FILE* f = freopen("fractional-knapsack/test.txt", "r", stdin);
	FILE* f = freopen("test.txt", "r", stdin);
    if (!f) {
        perror("Cannot open file!\n");
        return errno;
    }
    
    dijistra();

	fclose(f);
	return 0;
}
