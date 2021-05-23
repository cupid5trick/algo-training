#include <iostream>
#include <cstdio>
#include <vector>
#include <iterator>	// for prev
#include <algorithm>	// for for_each
#include <cerrno>       // for errno

using namespace std;

class Item {
public:
	int value;
	int weight;
	int id;
	bool operator() (const Item& i, const Item& j) {
		return (double)i.value / i.weight > (double)j.value / j.weight;
	}
};

void fractional_knapsack() {
	int n;
	scanf("%d", &n);
	vector<Item> items(n);
	
	for (int i = 0; i < n; ++ i) {
		scanf("%d", &items[i].value);
		items[i].id = i;
	}
	for (int i = 0; i < n; ++ i) {
		scanf("%d", &items[i].weight);
	}
	
//	for_each(items.begin(), items.end(), [](const Item& i) {
//		printf("%d %d %d\n", i.id+1, i.value, i.weight);
//	});
	
	sort(items.begin(), items.end(), items[0]);
	int volume = 100;
	
	int weight = 0;
	double value = 0;
	// begin greedy method
	for (auto it = items.begin(); volume > 0 && it != items.end(); ++ it) {
		int w = volume >= it->weight? it->weight: volume;
		volume -= w;
		weight += w;
		value += (double)it->value / it->weight * w;
		printf("%s%d", it == items.begin()? "":" ", it->id+1);
	}
	printf("\n");
	printf("Total value: %lf, Total weight: %d\n", value, weight);
}

int main() {
//    FILE* f = freopen("fractional-knapsack/test.txt", "r", stdin);
	FILE* f = freopen("test.txt", "r", stdin);
    if (!f) {
        perror("Cannot open file!\n");
        return errno;
    }
    
    fractional_knapsack();

	fclose(f);
	return 0;
}
