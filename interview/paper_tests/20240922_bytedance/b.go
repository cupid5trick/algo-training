package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
小红有一个长度为 n的数组 a1,a2,…,an，她每次会询问一个区间[l,r]她想知道数组a的所有长度大于等于l且小于等于r的子数组之和的最大值是多?

输入描述
第一行输入两个整数几和q(1≤n≤2000;1≤q≤ 10^6)代表数组中的元素数量和询问次数。
第二行输入 几 个整数 a1,a2,...,an(-10^9< ai < 10^9)代表数组元素此后q行，每行输入两个整数l,r(1 <=l<=r<= n)代表询问的长度区间。
输出描述
对于每个询问，在一行上输出一个整数代表答案
*/
var DEBUG bool = true

func debug(msg string, args ...interface{}) {
	if ! DEBUG {
		return
	}
	fmt.Fprintln(os.Stderr, msg, args)
}

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	// 读入数组
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	// 创建一个长度为 n 的切片来存储每个长度的最大子数组和
	maxSum := make([]int, n+1)

	// 计算所有长度子数组的和
	for start := 0; start < n; start++ {
		cur := 0
		for end := start; end < n; end++ {
			if cur > 0 {
				cur += a[end]
			} else {
				cur = a[end]
			}
			maxSum[end-start+1] = max(maxSum[end-start+1], cur)
		}
	}

	debug("maxsum: ", maxSum)

	// 处理每个询问
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)

		maxValue := maxSum[l]
		for length := l + 1; length <= r; length++ {
			if maxSum[length] > maxValue {
				maxValue = maxSum[length]
			}
		}
		fmt.Fprintln(output, maxValue)
	}
}
