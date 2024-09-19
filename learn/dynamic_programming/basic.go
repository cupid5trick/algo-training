package main

import (
	"fmt"
	"math"
)

/// 1186. 删除一次得到子数组最大和 - 力扣（LeetCode）: https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/description/
func maximumSum1(arr []int) (ans int) {
    n := len(arr)
    left, right := make([]int, n+1), make([]int, n+1)
    for i := 0; i < n; i ++ {
        left[i+1] = arr[i]
        if left[i] > 0 {
            left[i+1] += left[i]
        }
        right[n-i-1] = arr[n-i-1]
        if right[n-i] > 0 {
            right[n-i-1] += right[n-i]
        }
    }

    fmt.Println(left)
    fmt.Println(right)

    left[0] = math.MinInt/2
    right[n] = math.MinInt/2
    ans = math.MinInt
    for i := range arr {
        // left[i+1], right[i], left[i]+right[i]
        if left[i+1] > ans {
            ans = left[i+1]
        }
        if right[i] > ans {
            ans = right[i]
        }
        if left[i]+right[i+1] > ans {
            ans = left[i]+right[i+1]
        }
    }

    return
}

func maximumSum2(arr []int) int {
    n := len(arr)
    f := make([][2]int, n+1)
    f[0] = [2]int{math.MinInt/2, math.MinInt/2}
    ans := math.MinInt
    for i, x := range arr {
        f[i+1][0] = max(f[i][0], 0) + x
        f[i+1][1] = max(f[i][1]+x, f[i][0])
        ans = max(ans, max(f[i+1][1], f[i+1][0]))
    }

    return ans
}


/// 53. 最大子数组和 - 力扣（LeetCode）: https://leetcode.cn/problems/maximum-subarray/description/
func maxSubArray(nums []int) (ans int) {
    
    sum := 0
    ans = math.MinInt
    for _, x := range nums {
        if sum > 0 {
            sum += x
        } else {
            sum = x
        }

        ans = max(ans, sum)
    }
    return
}
