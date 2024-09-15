package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 小红拿到一个数组，准备把它分割成 k 段，使得每一段内部按位异或之后，在全部求和。
// 小红希望这个和尽可能大，你能帮帮她吗？
// 输入包含两个整数 n, k (1<=k<=n<=400) 分别表示数组长度和划分的段数。第二行是 n 个整数 ai 表示数组元素 (0<=ai<=10^9)。

func main_d() {
	run(os.Stdin, os.Stdout)
}

// dfs(i,k) = max(dfs(j, k-1) + (a[j+1]^...^a[i])), k-1<=j<=i-1
// 表示把当前位置的数字纳入第 k 段的最大异或值
func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	// nums := make([]int, n)
	pre := make([]int64, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &pre[i+1])
		pre[i+1] ^= pre[i]
	}
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, m+1)
	}
	ans := int64(0)
	for i := 1; i <= n; i++ {
		for k := 1; k <= m; k++ {
			// 到底正向更新还是逆向更新还是需要手算分析一下，好像是如果从前往后更新会出现数组没用完但是已经最大了？
			// 同时倒着更新可以把前缀和拆分到循环中去，实现滚动前缀和
			// 位运算一定要加括号保证优先级！！！
			// for j := k - 1; j <= i-1; j++ {
			// 	dp[i][k] = max(dp[i][k], dp[j][k-1]+(pre[i]^pre[j]))
			// }
			for j := i - 1; j >= k-1; j -- {
				dp[i][k] = max(dp[i][k], dp[j][k-1]+(pre[i]^pre[j]))
			}
		}
	}
	ans = dp[n][m]

	Fprintln(out, ans)
}
