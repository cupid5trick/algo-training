package main

import "sort"

// 56. 合并区间 - 力扣（LeetCode）: https://leetcode.cn/problems/merge-intervals
// 对区间按照左端点升序排序。当遍历过程发现前面有[aj, bj] 满足右端点和当前区间左端点有重叠时就可以合并！
// 当前面有区间且 bj >= ai，则合并后的区间是 [aj, max(bj, bi)]
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	ans := [][]int{}
	for _, x := range intervals {
		size := len(ans)
		if size == 0 || ans[size-1][1] < x[0] {
			ans = append(ans, x)
		} else {
			ans[size-1] = []int{ans[size-1][0], max(x[1], ans[size-1][1])}
		}
	}

	return ans
}
