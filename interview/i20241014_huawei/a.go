package main

import "fmt"
/*
- 209. 长度最小的子数组: <https://leetcode.cn/problems/minimum-size-subarray-sum/description/>

# 解题思路

1. 滑动窗口：维护所有“和大于等于target”的窗口，同时更新答案的最小值。
2. 前缀和:	遍历计算前缀和，同时维护之前计算过的前缀和中，满足 `pre + target <= sum` 的下标。
			因为是长度最小子数组，只维护最近的一个即可。
			如何维护：因为前缀和序列一定是递增的，可以使用二分法。
*/
func main() {

	nums := []int{2, 3, 1, 2, 4, 3}
	sum := 7
	left, right := 0, 0
	s := 0
	n := len(nums)
	ans := n + 1
	/// 滑动窗口
	for ; right < n; right++ {
		s += nums[right]

		for s >= sum {
			ans = min(ans, right-left+1)
			s -= nums[left]
			left++
		}
	}
	/// 边界处理
	if ans > n {
		fmt.Println(-1)
	}
	fmt.Println(ans)
}