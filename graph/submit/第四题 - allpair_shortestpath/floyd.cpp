#include <iostream>
#include <cstdio>
#include <vector>
#include <list>
#include <iterator>	// for prev
#include <algorithm>	// for for_each
#include <cerrno>       // for errno

using namespace std;


using Path=list<int>;

void floyd() {
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
	
	// floyd
	// each row of d are singlesource shortest pathes
	vector<vector<int>> d(n, vector<int>(n, INF));
	for (int i = 0; i < n; ++ i) {
		d[i][i] = 0;
	}
	vector<vector<int>> prev(n, vector<int>(n, INF));
	
	for (int k = 0; k < n; ++ k) {
		for (int i = 0; i < n; ++ i) {
			for (int j = 0; j < n; ++ j) {
				if (aj[k][j] == INF || d[i][k] == INF || i == j) {
					continue;
				}
				if (d[i][k]+aj[k][j] < d[i][j]) {
					d[i][j] = d[i][k]+aj[k][j];
					prev[i][j] = k;
				}
			}
		}
	}
	
//	for_each(aj.begin(), aj.end(), [](const vector<int>& v){
//		for_each(v.begin(), v.end(), [](const int x) {
//			printf(" %d", x);
//		});
//		printf("\n");
//	});
	
	// output
	for (int i = 0; i < n; ++ i) {
		for (int j = 0; j < n; ++ j) {
			if (i == j) {
				continue;
			}
			if (d[i][j] == INF) {
				printf("No path between %s and %s\n", nodes[i].c_str(), nodes[j].c_str());
				continue;
			}
			printf("Distance between %s and %s: %d\n", nodes[i].c_str(), nodes[j].c_str(), d[i][j]);
			vector<int> path;
			int k = prev[i][j];
			while (k != INF) {
				path.push_back(k);
				k = prev[i][k];
			}
			for (auto it = path.rbegin(); it != path.rend(); ++ it) {
				printf("%s->", nodes[*it].c_str());
			}
			printf("%s\n", nodes[j].c_str());
		}
	}
}


int main() {
//    FILE* f = freopen("fractional-knapsack/test.txt", "r", stdin);
	FILE* f = freopen("test.txt", "r", stdin);
    if (!f) {
        perror("Cannot open file!\n");
        return errno;
    }
    
    floyd();

	fclose(f);
	return 0;
}
