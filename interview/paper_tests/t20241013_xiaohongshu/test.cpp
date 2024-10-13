#include<bits/stdc++.h>
using namespace std;

/**
背景：在“歪星”上，有一场运动会正在进行，吸引了许多选手参加。小红书歪星频道正在直播这场运动会，观众们在小红书上互动讨论。

评分系统：每位选手的表现由来自小红书的n位评委评分，评分使用字母表示，从a（最优）到z（最次）。评分可以是小写或大写字母，它们之间有一一对应的关系。

评分处理：对于每位选手，n位评委给出的评分中，会去掉一个最高分和一个最低分，然后计算剩余评分的平均值，并向上取整得到一个新的字母评分。

输出要求：需要按照新评分从高到低的顺序输出选手的编号。如果有选手的评分相同，则编号小的选手排在前面。同时，还需要输出每位选手的最终评分字母。

输入格式：

第一行包含两个整数n和m，分别代表评委数量和选手数量（1 ≤ n ≤ 10^6，1 ≤ m ≤ 10^6）。
接下来的n行，每行是一个长度为m的字符串，由大小写字母混合构成，代表第i个评委对第j个选手的评分。
输出格式：

第一行输出m个整数，代表按新评分排序后的选手编号。
第二行输出m个小写字母，代表每位选手的最终评分。
示例输入：
4 5
abaKm
acAfe
czaCd
DbAVa
 */
int main(int argc, char const *argv[]) {
    int n , m;
    cin >> n >> m ;
    string score;
    vector<vector<int>> scores(m);
    vector<pair<char, int>> ans;
    for (int i = 0; i < n; i++) {
        cin >> score;
        for (int j = 0; j < m; j ++) {
            int x = score[j] >= 'a' && score[j] <= 'z' ? score[j] : score[j] - 'A' + 'a';
            scores[j].push_back(x);
        }
    }
    
    for (int i = 0; i < m; i ++) {
        sort(scores[i].begin(), scores[i].end());
        int x {};
        for (int j = 1; j < n-1; j ++) {
            x += scores[i][j];
        }
        x = (x-1)/(n-2)+1;
        // cout << (char) x << x << endl;
        ans.push_back(pair<char, int>((char)x, i+1));
    }

    sort(ans.begin(), ans.end());
    for (int i = 0; i < m; i ++) {
        cout << (i > 0 ? " ": "") << ans[i].second;
    }
    cout << endl;
    for (int i = 0; i < m; i ++) {
        cout << (i > 0 ? " ": "") << ans[i].first;
    }
    return 0;
}
