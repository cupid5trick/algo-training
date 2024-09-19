package main

/// 41. 缺失的第一个正数 - 力扣（LeetCode）: https://leetcode.cn/problems/first-missing-positive/description/
/// 是一个典型的构造法：因为要求时间复杂度 O(n)，所以不能排序。可以分析一下答案的范围：最大是 N+1，最小是 1。
/// 所以就可以考虑把原数组用来保存每个正整数是否出现过，如果出现过就设置为负数：nums[x-1] = -abs(nums[x-1])
/// 还要考虑 nums[x-1] 可能本来就是 0，所以这时要设置为 -(N-1)
func abs(x int) int{
    if x > 0 {
        return x
    }
    return -x
}

func firstMissingPositive(nums []int) int {
	/// 1. 不统计负数，予以删除
    for i, x := range nums {
        if x < 0 {
            nums[i] = 0
        }
    }
    n := len(nums)
	/// 2. 统计出现的正整数
    for _, x := range nums {
        x = abs(x)
        if x > 0 && x <= len(nums) {
            val := abs(nums[x-1])
            if val < 0 {
                val = 0
            }
            if val == 0 {
                nums[x-1] = -n-1
            } else {
                nums[x-1] = - val
            }
        }
    }

    if nums[0] >= 0 {
        return 1
    }
    ans := 1
    for i:=1; i < len(nums); i ++ {
        x := nums[i]
        if x < 0 && ans == i {
            ans ++
        }
    }
    return ans +1
}