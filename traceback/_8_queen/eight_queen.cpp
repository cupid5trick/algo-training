//
// Created by cupid5trick on 2021/5/28.
//

#include <iostream>
#include <cstdio>
#include <vector>
#include <list>
#include <stack>
#include <iterator>		// for prev
#include <algorithm>	// for for_each
#include <cerrno>       // for errno

using namespace std;


class Solution {
public:
    using Pos=vector<int>;
    Pos pos;
    stack<Pos> valid;
    vector<bool> allow;

public:
    Solution(int n): pos(vector<int>(n, -1)) {}

    // Solution::legal_area(int r);
    // 根据已经确定的皇后位置（pos[i]!=-1）返回第r行皇后的可行位置
    const vector<bool>& legal_area(int r) {
        allow = vector<bool>((int)pos.size(), true);
        for (int i = 0; i < r; ++ i) {
            if (pos[i] >= 0) {
                // 不同能处在同一列、同一对角线
                allow[pos[i]] = false;
                if (pos[i]+(r-i) >= 0 && pos[i]+(r-i) < (int)pos.size()) {
                    allow[pos[i]+(r-i)] = false;
                }
                if (pos[i]-(r-i) >= 0 && pos[i]-(r-i) < (int)pos.size()) {
                    allow[pos[i]-(r-i)] = false;
                }
            }
        }
        return allow;
    }

    // Solution::search(int r);
    // use as search(0)
    void search(int r) {
        // 如果搜索到了r=pos.size()说明1...r个皇后都是互相不冲突的
        if (r == (int)pos.size()) {
            valid.push(pos);
            return ;
        }

        vector<bool> allow = legal_area(r);

        for (int i = 0; i < (int)allow.size(); ++i) {
            if (allow[i]) {
                pos[r] = i;
                search(r+1);
                // 遍历完所有皇后的可行位置后，回溯子分支的状态
                pos[r] = -1;
            }
        }
    }

    void traceback_solve() {
        search(0);

        int i = 0;
        int n = valid.size();
        FILE* f = fopen("res.md", "w");
        fprintf(f, "|ID");
        for (int i = 0; i < (int)this->pos.size(); ++ i) {
            fprintf(f, "|queen%d", i+1);
        }
        fprintf(f, "|\n");
        for (int i = 0; i < (int)this->pos.size()+1; ++ i) {
            fprintf(f, "|:--:");
        }
        fprintf(f, "|\n");
        while (!valid.empty()) {
            printf("method %d\n", i);
            ++ i;
            fprintf(f, "|%d", i);
            for (int j = 0; j != (int)valid.top().size(); ++ j) {
                printf("%s%d", j==0? "":" ", valid.top()[j]+1);
                fprintf(f, "|%d", valid.top()[j]+1);
            }
            printf("\n");
            fprintf(f, "|\n");

            valid.pop();
        }
        fclose(f);
        printf("%d methods\n", n);
    }
};

void solve() {
    int n;
    scanf("%d", &n);
    Solution solver(n);
    printf("solver ok;\n");
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
