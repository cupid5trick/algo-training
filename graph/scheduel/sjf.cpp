#include <iostream>
#include <cstdio>
#include <vector>
#include <list>
#include <iterator>	// for prev
#include <algorithm>	// for for_each
#include <cerrno>       // for errno

using namespace std;

class Job{
public:
	int cost;
	int id;
};

void sjf() {
	int n;
	scanf("%d", &n);
	vector<Job> C(n);
	for (int i = 0; i < n; ++ i) {
		int t;
		scanf("%d", &C[i].cost);
		C[i].id = i + 1;
	}
	
	sort(C.begin(), C.end(), [](const Job& x, const Job& y){
		return x.cost < y.cost;
	});
	
	double avg = 0;
	printf("Schedueled tasks ");
	for (int i = 0; i < n; ++ i) {
		printf("%d ", C[i].id);
		for (int j = 0; j <= i; ++ j) {
			avg += C[j].cost;
		}
	}
	printf("with average completion time = %lf\n", avg/n);
}


int main() {
//    FILE* f = freopen("fractional-knapsack/test.txt", "r", stdin);
	FILE* f = freopen("test.txt", "r", stdin);
    if (!f) {
        perror("Cannot open file!\n");
        return errno;
    }
    
    sjf();

	fclose(f);
	return 0;
}
