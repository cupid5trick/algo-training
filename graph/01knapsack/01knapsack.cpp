#include <iostream>
#include <cstdio>
#include <vector>
#include <iterator>	// for prev
#include <algorithm>	// for for_each
#include <cerrno>       // for errno

using namespace std;


void _01_knapsack() {
	
	const int volume = 100;
	int n;
	scanf("%d", &n);
	vector<int> v(n);
	vector<int> weights(n);
	
	for (int i = 0; i < n; ++ i) {
		scanf("%d", &v[i]);
	}
	for (int i = 0; i < n; ++ i) {
		scanf("%d", &weights[i]);
	}
	
	// dp[i,j]=0, dp2[i,j]=0
	vector<vector<int>> dp(n+1, vector<int>(volume+1));
	vector<vector<int>> dp2(n+1, vector<int>(volume+1));
	
	// dp
	/*
	** V[i,w] = max(V(i-1, w-weights[i])+v[i], V(i-1, w)), if i>0 && w>=weights[i]
	**          V(i-1, w), if i>0 && w<weights[i]
	**          0, otherwise (i=0 || w=0)
	*/
	for (int i = 1; i <= n; ++ i) {
		for (int w = 1; w <= volume; ++ w) {
			if (w >= weights[i]) {
				if (dp[i-1][w-weights[i]]+v[i] > dp[i-1][w]) {
					// item i is taken
					dp[i][w] = dp[i-1][w-weights[i]]+v[i];
					dp2[i][w] = 1;
				}
				else {
					dp[i][w] = dp[i-1][w];
					dp2[i][w] = 0;
				}
			}
		}
	}
	
	int i = n, w = volume;
	int value = 0, weight = 0;
	vector<int> items;
	while (i > 0) {
		if (dp2[i][w]) {
			// item i taken
			value += v[i-1];
			weight += weights[i-1];
			w -= weights[i-1];
			items.push_back(i);
		}
		-- i;
	}
	for (auto it = items.rbegin(); it != items.rend(); ++ it) {
		printf("%s%d", it == items.rbegin()? "":" ", *it);
	}
	printf("\nTotal value: %d, Total weight: %d\n", value, weight);
}

int main() {
//    FILE* f = freopen("fractional-knapsack/test.txt", "r", stdin);
	FILE* f = freopen("test.txt", "r", stdin);
    if (!f) {
        perror("Cannot open file!\n");
        return errno;
    }
    
    _01_knapsack();

	fclose(f);
	return 0;
}
