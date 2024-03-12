#include <iostream>
#include <cstdio>
#include <vector>
#include <list>
#include <climits>
#include <algorithm>

using namespace std;

const int INF = INT_MAX;


pair<vector<list<list<int>>>, vector<int>> dijistra(vector<vector<int>>& aj_matrix, int src) {
	using Path = list<int>;
	
	int n = aj_matrix.size();
	vector<bool> visits(n, false);
	
	vector<list<Path>> path(n);
	vector<int> pathc(n);
	vector<int> d(n, INF);
	
	d[src] = 0;
	path[src] = list<Path>(1, Path(1, src));
	pathc[src] = 1;
	
	int cur;
	while (n) {
		int min = INF, argmin=-1;
		for (int i=0; i<visits.size(); ++ i) {
			if (!visits[i] && d[i] < min) {
				min = d[i];
				argmin = i;
			}
		}
		if (argmin == -1) {
			break;
		}
		cur = argmin;
		printf("cur=%d, min=%d\n", cur, min);
		visits[cur] = true;
		-- n;
		for (int i = 0; i < visits.size(); ++ i) {
			if (visits[i] || aj_matrix[cur][i] == INF) {
				continue;
			}
			if (d[cur]+aj_matrix[cur][i] < d[i]) {
				d[i] = d[cur] + aj_matrix[cur][i];
				pathc[i] = pathc[cur];
				path[i] = path[cur];
				for_each(path[i].begin(), path[i].end(), [i](Path& p) {
					p.push_back(i);
				});
			}
			else if (d[cur]+aj_matrix[cur][i] == d[i]) {
				pathc[i] += pathc[cur];
				list<Path> pl = path[cur];
				for_each(pl.begin(), pl.end(), [i](Path& p) {
					p.push_back(i);
				});
				path[i].insert(path[i].end(), pl.begin(), pl.end());
			}
		}
		for_each(d.begin(), d.end(), [](int& x) {
			printf(" %d", x);
		});
		printf("\n");
	}
	return make_pair(path, pathc);
}

void run_dijistra() {
	int n, m, src;
	scanf("%d %d %d\n", &n, &m, &src);
	vector<vector<int>> aj(n, vector<int>(n, INF));
	for (int i = 0; i < m; ++ i) {
		int u, v, w;
		scanf("%d %d %d\n", &u, &v, &w);
		aj[u][v] = w < aj[u][v]? w: aj[u][v];
	}
		
	pair<vector<list<list<int>>>, vector<int>> p = dijistra(aj, src);
	for (int i = 0; i < p.second.size(); ++ i) {
		printf("Node %d: %d pathes\n", i, p.second[i]);
		for (auto it = p.first[i].begin(); it != p.first[i].end(); ++ it) {
			for (auto j = it->begin(); j != it->end(); ++ j) {
				printf("%s%d", j == it->begin()? "":"->", *j);
			}
			printf("\n");
		}
	}
}


int main() {
	FILE* f = freopen("test.txt", "r", stdin);
	if (! f) {
		printf("Cannot open file!\n");
		return 1;
	}
	
	run_dijistra();
	
	fclose(f);
	
	return 0;
}
