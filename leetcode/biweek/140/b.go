package main

import "sort"

/**
3301. 高度互不相同的最大塔高和 - 力扣（LeetCode）: https://leetcode.cn/problems/maximize-the-total-height-of-unique-towers/description/
贪心尝试每个塔尽量取到最高值即可。
 */

func maximumTotalSum(nums []int) (ans int64) {
    sort.Ints(nums)
    n := len(nums)
    // fmt.Println(nums)
    mx := 0x7fff_ffff
    for i := n-1; i >= 0; i -- {
        x := nums[i]
        mx= min(x, mx-1)
        // fmt.Println(mx)
        if mx <= 0 {
            ans = -1
            break
        }
        ans += int64(mx)
    }
    return
}

func min(x, i int) int {
	if x > i {
		return x
	}
	return i
}