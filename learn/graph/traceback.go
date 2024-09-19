package main

import "math"

/// 2850. 将石头分散到网格图的最少移动次数 - 力扣（LeetCode）: https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/description/
/// 题目要求可以总结为：从石头数量多的格子移动到石头数量为0格子的步数，所以只需要暴力枚举每个石头移动到哪个格子计算步数取最小值即可。
/// 考查全排列计算：回溯技巧。
/// 03xf 给出了另一种“最小费用最大流”的解法，时间复杂度为 O((mn)^4)
/// 下面是一些涉及到「匹配」的题目：
// -   [1947. 最大兼容性评分和](https://leetcode.cn/problems/maximum-compatibility-score-sum/)
// 	-   [1349. 参加考试的最大学生数](https://leetcode.cn/problems/maximum-students-taking-exam/)
// 	-   [LCP 04. 覆盖](https://leetcode.cn/problems/broken-board-dominoes/)
// 	-   [1879. 两个数组最小的异或值之和](https://leetcode.cn/problems/minimum-xor-sum-of-two-arrays/)
// 	-   [2172. 数组的最大与和](https://leetcode.cn/problems/maximum-and-sum-of-array/)
type pair struct{ x, y int }

func minimumMoves(grid [][]int) int {
	from := []pair{}
	to := []pair{}

	for i, row := range grid {
		for j, cnt := range row {
			if cnt > 1 {
				for k := 1; k < cnt; k++ {
					from = append(from, pair{i, j})
				}
			} else if cnt == 0 {
				to = append(to, pair{i, j})
			}
		}
	}

	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}

	cnt := len(from)
	minop := math.MaxInt
	dist := func(to []pair) (d int) {
		for i, point := range to {
			d += abs(point.x-from[i].x) + abs(point.y-from[i].y)
		}
		return
	}
	/// 全排列
	var permute func(int)
	permute = func(i int) {

		if i == cnt {
			// res := append([][]int{}, to...)
			minop = min(minop, dist(to))
			// fmt.Println(cnt, from, to, dist(to))
			return
		}
		for j := i; j < cnt; j++ {
			to[i], to[j] = to[j], to[i]
			permute(i + 1)
			to[i], to[j] = to[j], to[i]
		}
	}

	permute(0)
	return minop
}