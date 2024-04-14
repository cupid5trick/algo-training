package main

import "math"

// 100259. 划分数组得到最小的值之和 - 力扣（LeetCode）: https://leetcode.cn/problems/minimum-sum-of-values-by-dividing-array/description/
// 这一题需要了解按位与、按位或、GCD 等运算的单调性。除此之外就是一个划分型DP的技巧。

// 解题思路：
// 用 dfs(i, j, v) 表示 nums[i] 划入第 j 段，这一段的按位与是 v 子数组"值" 是 dfs(i,j,v)
// dfs(i,j,v) = min(dfs(i+1,j,v & ai), dfs(i+1,j+1,-1)+ai), ai = nums[i]
// 初始条件: dfs(0, 0, -1) 因为 -1 = 0xffffffff
// 1. j==m 时划分结束 如果i==n 说明找到一个合法结果返回0，否则返回表示非法的无穷大
// 2. 如果 m-j < n-i 说明数字不够了，返回非法结果
// 3. 按位与有单调性，如果当前 v < andValues[i] 返回非法结果
// 4. Go语言中记忆化数组的实现也是一个难点。and_表示一些数字按位与的结果，它是一个32位的int。
// 但是我们无法在内存中开辟这么大的数组，所以需要做一些位运算优化，用map[int64]int{}代替mem数组。
// 可以用int64的低32位表示and_的值，而高32位中可以用这部分的4个低位表示j，剩下的表示i。因为j和i最大分别是10和10^4。

var mem map[int64]int

func minimumValueSum(nums []int, andValues []int) int {
	n, m := len(nums), len(andValues)
	mem = make(map[int64]int)

	var dfs func(i, j, v int) int
	dfs = func(i, j, and_ int) int {
        // 2. 如果 m-j < n-i 说明数字不够了，返回非法结果
		if m-j > n-i {
			return math.MaxInt / 2
		}
        // 1. j==m 时划分结束 如果i==n 说明找到一个合法结果返回0，否则返回表示非法的无穷大
		if j == m {
			if i == n {
				return 0
			}
			return math.MaxInt / 2
		}
		key := int64(i)<<36 + int64(j)<<32 + int64(and_)
		if val, ok := mem[key]; ok {
			return val
		}

		and_ &= nums[i]
		if and_ < andValues[j] {
            // 3. 剪枝：按位与有单调性，如果当前 v < andValues[i] 返回非法结果
			mem[key] = math.MaxInt / 2
		} else {
			res := dfs(i+1, j, and_)
			if and_ == andValues[j] {
				res = min(res, dfs(i+1, j+1, -1)+nums[i])
			}
			mem[key] = res
		}
		return mem[key]
	}

	ans := dfs(0, 0, -1)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}
