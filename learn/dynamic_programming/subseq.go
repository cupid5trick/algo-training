package main

import (
	"math"
	"sort"
)

/*
子序列专题
*/

// 300. 最长递增子序列 - 力扣（LeetCode）: https://leetcode.cn/problems/longest-increasing-subsequence/
// 这是一个选择之间有顺序依赖的背包问题
// 维护小于当前数字的所有值 dp[x] = max(dp[x], dp[k]+1), k < x
// - 暴力方法：O(n^2)
// - O(nlogn): 构造法
// 0x3f 讲题链接，最长递增子序列【基础算法精讲 20】_哔哩哔哩_bilibili: https://www.bilibili.com/video/BV1ub411Q7sB
func lengthOfLIS_bforce(nums []int) int {
	dp := map[int]int{}
	n := len(nums)

	ans := 0
	for i := 0; i < n; i++ {
		dp[nums[i]] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[nums[i]] = max(dp[nums[i]], dp[nums[j]]+1)
			}
		}
		ans = max(ans, dp[nums[i]])
	}

	return ans
}

// 通过二分法查找当前构造的子序列如果存在第一个比 x 大的，替换为 x
//              否则把 x 加入子序列。答案就是子序列长度！
//              构造法，有点单调栈的意思。
func lengthOfLIS_constr(nums []int) int {
	g := []int{}

	for _, x := range nums {
		pos := sort.SearchInts(g, x)
		if pos >= len(g) {
			g = append(g, x)
		} else {
			g[pos] = x
		}
	}

	return len(g)
}

// 2915. 和为目标值的最长子序列的长度 - 力扣（LeetCode）: https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target
// 03xf 讲题地址，0-1背包 完全背包_哔哩哔哩_bilibili: https://www.bilibili.com/video/BV16Y411v7Y6/?vd_source=23639efd7feab5bca1547c67b23ba88f
// dp[x] = max(dp[x], dp[x-num]+1)
func lengthOfLongestSubsequence(nums []int, target int) int {
    dp := make([]int, target+1)
    for i:=1; i<= target; i ++ {
            dp[i] = math.MinInt;
    }

    for _, num := range nums {
		// 得研究下这里，为什么倒着遍历可以，正着走就不对???
        for x:=target; x >= num; x -- {
            dp[x] = max(dp[x], dp[x-num]+1)
        }
    }

    if dp[target] < 0 {
        return -1
    }
    return dp[target]
}