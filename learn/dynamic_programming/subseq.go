package main

import (
	"math"
	"slices"
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
//	否则把 x 加入子序列。答案就是子序列长度！
//	构造法，有点单调栈的意思。
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
// 03xf 讲题地址，0-1背包 完全背包_哔哩哔哩_bilibili: https://w.bilibili.com/video/BV16Y411v7Y6/?vd_source=23639efd7feab5bca1547c67b23ba88f
// dp[x] = max(dp[x], dp[x-num]+1)
func lengthOfLongestSubsequence(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = math.MinInt
	}

	for _, num := range nums {
		// 得研究下这里，为什么倒着遍历可以，正着走就不对?
		// 因为构成子序列满足和为 x 的过程中只能用一次当前数字 num，如果是正向遍历的话前一个 dp[x-num] 有可能已经用到了 num !
		for x := target; x >= num; x-- {
			dp[x] = max(dp[x], dp[x-num]+1)
		}
	}

	if dp[target] < 0 {
		return -1
	}
	return dp[target]
}

// 2952. 需要添加的硬币的最小数量 - 力扣（LeetCode）: https://leetcode.cn/problems/minimum-number-of-coins-to-be-added/
// 构造法！！！
// 为方便描述，把 0 也算作可以得到的数。
// 假设现在得到了区间 [0,s-1] 中的所有整数，如果此时遍历到整数 x=coins[i]，那么把 [0,s−1] 中的每个整数都增加 x，我们就得到了区间 [x,s+x−1] 中的所有整数。
// 思考
// 把 coins 从小到大排序，遍历 x=coins[i]。分类讨论，看是否要添加数字：
// 如果 x≤s，那么合并 [0,s−1] 和 [x,s+x−1] 这两个区间，我们可以得到 [0,s+x−1] 中的所有整数。
// 如果 x>s，或者遍历完了 coins 数组，这意味着我们无法得到 s，那么就一定要把 s 加到数组中
// （加一个比 s 还小的数字就没法得到更大的数，不够贪），这样就可以得到了 [s,2s−1] 中的所有整数，
// 再与 [0,s−1] 合并，可以得到 [0,2s−1] 中的所有整数。然后再考虑 x 和 2s 的大小关系，继续分类讨论。
// 当 s>target 时，我们就得到了 [1,target] 中的所有整数，退出循环。
// 题解链接：https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/solutions/2502839/mo-ban-qia-hao-zhuang-man-xing-0-1-bei-b-0nca
func minimumAddedCoins(coins []int, target int) (ans int) {
	slices.Sort(coins)
	s, i := 1, 0
	for s <= target {
		if i < len(coins) && coins[i] <= s {
			s += coins[i]
			i++
		} else {
			s *= 2 // 必须添加 s
			ans++
		}
	}
	return
}
