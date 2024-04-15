package main

// 单调栈
// 100273. 边界元素是最大值的子数组数目 - 力扣（LeetCode）: https://leetcode.cn/problems/find-the-number-of-subarrays-where-boundary-elements-are-maximum
// 解题思路：
// 根据题目要求一旦发现前面有数字比当前数字大，就无法构成一个子数组使得最左边和最右边数字都是子数组的最大值。
// 我们可以枚举每个数字作为子数组的右边界，如果当前数字比前面遍历到的数字大，
// 前面的数字就无法构成一个长度大于一且满足条件的子数组，可以删除。
// 根据这种及时删除无用数据的单调性，可以采用单调栈方法：发现栈顶数字比当前数字小，就弹出栈顶元素。直到栈顶元素大于等于当前数字。
// 另外单独一个数字也是一个满足题目要求的子数组，我们可以把答案初始化为 len(nums)。

type info struct{ v, cnt int }

func numberOfSubarrays(nums []int) (ans int64) {
	stk := []info{}

	ans = int64(len(nums))
	for _, x := range nums {
		// m := len(stk)
		for len(stk) > 0 && stk[len(stk)-1].v < x {
			stk = stk[:len(stk)-1]
		}
		m := len(stk)
		if m == 0 || stk[m-1].v > x {
			stk = append(stk, info{x, 1})
		} else {
			ans += int64(stk[m-1].cnt)
			stk[m-1].cnt++
		}
	}

	return
}
