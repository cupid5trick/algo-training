//
// Created by cupid5trick on 2021/5/27.
//

#include <iostream>
#include <cstdio>
#include <vector>
#include <list>
#include <iterator>		// for prev
#include <algorithm>	// for for_each
#include <cerrno>       // for errno

using namespace std;

class Item {
public:
    int id;
    int v;
    int w;
    double c;

public:
    Item() {}
    Item(int _id, int _v, int _w): id(_id), v(_v), w(_w) {
        this->c = (double)(this->v)/this->w;
    }

    bool operator() (const Item& one, const Item& other) {
        return one.v > other.v;
    }
};

class Solution {
	
public:
    int _n;
    int volume;
    vector<Item> items;
    vector<int> most_valuables;
    vector<bool> taken;

    int cur_volume;
    int cur_value;
    int max_value;
    vector<bool> opt_taken;

public:
    Solution(int n, int vol, const vector<int>& v, const vector<int>& w): _n(n), volume(vol) {
        items = vector<Item>(this->_n);
        for (int i = 0; i < items.size(); ++ i) {
            items[i] = Item(i+1, v[i], w[i]);
        }
    }

    // use as search(-1, false);
    void search(int i, bool takes) {
        if (i >= this->_n+1) {
            return ;
        }

        // 剪枝， 如果已经超出容量 或 已经不可能获得最大价值，则剪枝回溯
        if (this->cur_volume < 0 ||
            (i > 0 && i < this->_n && cur_value + most_valuables[this->_n-i-1] <= max_value)) {

            this->cur_value -= this->taken[i-1]*this->items[i-1].v;
            this->cur_volume += this->taken[i-1]*this->items[i-1].w;
            this->taken[i-1] = false;
            return ;
        }

        // 所有物品取舍确定，检查是否出现最大价值，然后回溯
        if (i == this->_n) {
            if (this->cur_value > this->max_value && this->cur_volume >= 0) {
                this->max_value = this->cur_value;
                this->opt_taken = this->taken;
            }

            this->cur_value -= this->taken[i-1]*this->items[i-1].v;
            this->cur_volume += this->taken[i-1]*this->items[i-1].w;
            this->taken[i-1] = false;
            return ;
        }

        if (i >= 0 && takes) {
            this->cur_volume -= this->items[i].w;
            cur_value += this->items[i].v;
            this->taken[i] = true;
        }

        log(i);
        search(i+1, false);
        search(i+1, true);

        // 遍历完所有子分支后回溯状态
        if (i >= 0) {
            this->cur_value -= this->taken[i]*this->items[i].v;
            this->cur_volume += this->taken[i]*this->items[i].w;
            this->taken[i] = false;
        }
    }

    void traceback_solve() {
        vector<int> items_cpy(items.size(), 0);
        for (int i = 0; i < items.size(); ++ i) {
            items_cpy[i] = items[i].v;
        }
        sort(items_cpy.begin(), items_cpy.end(), [](int x, int y) {
            return x > y;
        });
        // compute for cut branches
        most_valuables = vector<int>(this->_n, 0);

        for (int i = 0; i < items.size(); ++ i) {
            for (int j = 0; j <= i; ++ j) {
                most_valuables[i] += items_cpy[j];
            }
        }

        // for search
        this->taken = vector<bool>(this->_n, false);
        this->cur_volume = this->volume;
        this->max_value = this->cur_value = 0;

        // search
        search(-1, false);

        printf("The max value: %d, total weight: %d\n", this->max_value, this->volume-this->cur_volume);
        int n_take = 0;
        for (int i = 0; i < this->_n; ++ i) {
            if (this->opt_taken[i]) {
                printf("%s%d", n_take==0? "":" ", this->items[i].id);
                n_take ++ ;
            }
        }
        printf("\n");

    }

    void log(int i) {
        if (i < 0) {
            return ;
        }
        string s = string("    ");
        string b("");
        for (int j = 0; j < i; ++ j) {
            b+=s;
        }
        printf("%st%d=%d, value=%d weight=%d\n",
               b.c_str(), this->items[i].id, (int)this->taken[i], this->cur_value, this->volume-this->cur_volume);
    }
};

void solve() {
    int n, volume;
    scanf("%d %d", &n, &volume);
    vector<int> v(n, 0);
    vector<int> w(n, 0);
    for (int i = 0; i < n; ++ i) {
        scanf("%d", &v[i]);
    }
    for (int i = 0; i < n; ++ i) {
        scanf("%d", &w[i]);
    }
    Solution solver(n, volume, v, w);
    solver.traceback_solve();
}

int main() {
//    FILE* f = freopen("_01_knapsack/test.txt", "r", stdin);
    FILE* f = freopen("test.txt", "r", stdin);
    if (!f) {
        perror("Cannot open file!\n");
        return errno;
    }

    solve();

    fclose(f);
    return 0;
}
