#include <bits/stdc++.h>

using namespace std;

// 方法一：哈希表加有序数组
class NumberContainers {
    unordered_map<int, int> nums {};
    unordered_map<int, set<int>> occur {};
public:    
    void change(int index, int number) {
        // erase old value (if exists)
        auto& x = nums[index];
        auto& exists = occur[x];
        if (exists.find(index) != exists.end()) {
            exists.erase(index);
        }
        // insert new
        occur[number].insert(index);
        nums[index] = number;
    }
    
    int find(int number) {
        return !occur[number].empty()? *occur[number].begin() : -1;
    }
};

// 方法二：哈希表+懒删除堆
