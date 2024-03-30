package main

import "sort"

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
