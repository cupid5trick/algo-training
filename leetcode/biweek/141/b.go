package main

//lint:file-ignore SA4009 overwritten before first use. 

import (
	"fmt"
)

/*
- 3315. 构造最小位运算数组 II: <https://leetcode.cn/problems/construct-the-minimum-bitwise-array-ii/description/>

第二题和第一题只有数据范围不一样，就只写找出O(n)算法的过程了。

解题思路：
对于 x|(x+1) = nums[i] 这个关系式，观察一下左边的位运算。

1. 	任何数的或运算只会比自己大或者不变
2. 	x+1 的二进制位改变的部分只有，x 从最低位开始连续的 1 变成了 0、下一个 0 变成 1.
	所以 x|(x+1) 的值和【从最低位开始连续的 1 】的个数有关
3.	结合 1 和 2，求 x 时可以用 nums[i] 做初值，也不用依次遍历，逐个消除最低位连续的 1 即可
	消除的方向：从高位开始向右，也可以打表观察一下。

x|(x+1)	 x
2,		 -1,		 00000010,	 -0000001
3,		 1,			 00000011,	 00000001
5,		 4,			 00000101,	 00000100
7,		 3,			 00000111,	 00000011
11,		 9,			 00001011,	 00001001
13,		 12,		 00001101,	 00001100
17,		 16,		 00010001,	 00010000
19,		 17,		 00010011,	 00010001
23,		 19,		 00010111,	 00010011
29,		 28,		 00011101,	 00011100
31,		 15,		 00011111,	 00001111
37,		 36,		 00100101,	 00100100
41,		 40,		 00101001,	 00101000
43,		 41,		 00101011,	 00101001
47,		 39,		 00101111,	 00100111
53,		 52,		 00110101,	 00110100
59,		 57,		 00111011,	 00111001
61,		 60,		 00111101,	 00111100
67,		 65,		 01000011,	 01000001
71,		 67,		 01000111,	 01000011
73,		 72,		 01001001,	 01001000
79,		 71,		 01001111,	 01000111
83,		 81,		 01010011,	 01010001
89,		 88,		 01011001,	 01011000
97,		 96,		 01100001,	 01100000
*/

func minBitwiseArray(nums []int) []int {
	return bitwise(nums)
}

func bitwise(nums []int) []int {
	ans := []int{}
	for _, x := range nums {
		tmp := len(ans)
		bitlen := 0
		for ; (x>>bitlen)&1 == 1; bitlen++ {
		}
		fmt.Printf("x: %d, bitlen: %d\n", x, bitlen)
		for k := bitlen - 1; k >= 0; k-- {
			mask := ^(1 << k)
			res := x & mask
			fmt.Printf("%032b, %032b, %08b, %d\n", x, res, mask, k)
			if res|(res+1) == x {
				ans = append(ans, res)
				break
			}
		}
		if len(ans) == tmp {
			ans = append(ans, -1)
		}
	}
	return ans
}

// 暴力打表找规律
// go-
func bf(primes []int) []int {
	ans := []int{}
	check := func(x int) bool {
		for i := 3; i *i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	primes = []int{}
	for i := 2; i < 101; i++ {
		if check(i) {
			primes = append(primes, i)
		}
	}

	for _, x := range primes {
		tmp := len(ans)
		for i := 0; i < 1001; i++ {
			if i|(i+1) == x {
				ans = append(ans, i)
				break
			}
		}
		if len(ans) == tmp {
			ans = append(ans, -1)
		}
	}
	for i, res := range ans {
		x := primes[i]
		fmt.Printf("%d, %d, %08b, %08b\n", x, res, x, res)
	}
	return ans
}

func main() {
	nums := []int{3, 10, 6}
	fmt.Println("Bitwise Result:", minBitwiseArray(nums))
	fmt.Println("BF Result:", bf(nums))
}