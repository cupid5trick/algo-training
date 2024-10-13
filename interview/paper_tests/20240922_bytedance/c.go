package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

/**
小红有一个长度为 n 的字符串 S，现在她可以执行以下的操作:
选择一个素引i(1 ≤i≤ n)，并将 S 按照字母表上的顺序、循环右移一位。例如:'a'右移-位为'b'，'z'右移-位为'a'。
小红想知道使得字符串 S 任意两个相邻的字符都不一致的最小操作次数为多少。

输入描述
第一行输入一个整数n(2 ≤ n ≤ 10^5)代表字符串的长度。第二行输入一个长度为n，且只包含小写字母的字符串S
输出描述
在一行上输出一个整数，代表使得任意两个相邻的字符都不一致的最小操作次数

# 题解
子序列动态规划方法有“枚举选哪个”和“枚举选与不选”两种方法。

- 相邻相关：枚举选哪个
- 相邻无关：枚举选与不选
*/

func main_c() {
	run(os.Stdin, os.Stdout)
}

func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var n int
	var s string
	Fscan(in, &n)
	Fscan(in, &s)

	// dp[i][c] 表示将第i个字符变为字母c的最小操作次数
	dp := make([][26]int, n)

	// 初始化第一个字符的操作代价
	for c := 0; c < 26; c++ {
		dp[0][c] = cost(s[0], byte(c+'a'))
	}

	// 动态规划状态转移
	for i := 1; i < n; i++ {
		for c := 0; c < 26; c++ {
			dp[i][c] = math.MaxInt32
			for prevC := 0; prevC < 26; prevC++ {
				if prevC != c {
					dp[i][c] = min(dp[i][c], dp[i-1][prevC]+cost(s[i], byte(c+'a')))
				}
			}
		}
	}

	// 找到最后一个字符的最小代价
	result := math.MaxInt32
	for c := 0; c < 26; c++ {
		result = min(result, dp[n-1][c])
	}

	Fprintln(out, result)
}

// 计算字符转换的代价
func cost(from, to byte) int {
	if from == to {
		return 0
	}
	return (int(to-from) + 26) % 26
}
