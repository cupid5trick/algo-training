package main

/**
3300. 替换为数位和以后的最小元素 - 力扣（LeetCode）: https://leetcode.cn/problems/minimum-element-after-replacement-with-digit-sum/description/
打卡题
*/

func minElement(nums []int) (ans int) {
    ans = 0x7fff_ffff
    for _, x := range nums {
        s := 0
        for x > 0 {
            s += x%10
            x /= 10
        }
        ans = min(ans, s)
    }
    return
}