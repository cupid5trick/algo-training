package main

import (
	"math"
	"math/bits"
	"sort"
)

// 100267. 单面值组合的第 K 小金额 - 力扣（LeetCode）: https://leetcode.cn/problems/kth-smallest-amount-with-single-denomination-combination/description/
// 容斥原理 + 二分
// 题目比较难，涉及到容斥原理、最大公约数和最小公倍数、枚举子集（这里使用了二进制位运算的枚举子集）
// 分享｜从集合论到位运算，常见位运算技巧分类总结！ - 力扣（LeetCode）: https://leetcode.cn/circle/discuss/CaOJ45/

// 最小公倍数：lcm(a,b) = a*b/gcd(a,b)
func lcm(a, b int) int {
	ta, tb := a, b
	if b > a {
		a, b = b, a
	}
	for b > 0 {
		a, b = b, a%b
	}
	// 为了避免溢出可能需要先除再乘
	return ta / a * tb
}

func findKthSmallest(coins []int, k int) int64 {
    n := len(coins)

    check := func (mid int) bool {
        cnt := 0
        // 位运算枚举子集
        for i:=1; i < 1 << n; i ++ {
            lc := 1
            for j:=0; j < n; j ++ {
                if i >> j & 1 == 1 {
                    lc = lcm(lc, coins[j])
                }
            }
            c := mid / lc
            if bits.OnesCount(uint(i)) & 1 == 0 {
                c = -c
            }
            cnt += c
        }
        return cnt >= k
    }
    
    upper := math.MaxInt
    for i := range coins {
        upper = min(upper, coins[i])
    }
    upper *= k
    ans := sort.Search(upper, check)
    if ans < math.MaxInt64 {
        return int64(ans)
    }
    return -1
}