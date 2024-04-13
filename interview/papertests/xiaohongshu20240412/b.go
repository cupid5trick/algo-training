package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// 小苯是“小红书app”的忠实用户，他有n个账号，每个账号粉丝数为。
// 这天他又创建了一个新账号，他希望新账号的粉丝数恰好等于x。为此他可以向自己已有账号的粉丝们推荐自己的新账号，这样以来新账号就得到了之前粉丝的关注。
// 他想知道，他最少需要在几个旧账号发“推荐新账号”的文章，可以使得他的新账号粉丝数恰好为x，除此以外，他可以最多从中选择一个账号多次发“推荐新账号”的文章。
// (我们假设所有旧账号的粉丝们没有重叠，并且如果在第i个旧账号的粉丝们推荐了新账号，则新账号会直接涨粉 ai/2 下取整个，而如果小苯选择在第 i 个旧账号中多次推荐新账号，那么新账号就可以直接涨粉 ai。)
// 输入描述:
// 输入包含2行。
// 第一行两个正整数 n, x(1 ≤ n, k ≤ 100)，分别表示小苯的旧账号个数，和新账号想要的粉丝数。
// 第二行n个正整数 (1 ≤ ai ≤ 100)，表示小苯每个旧账号的粉丝数。
// 输出描述:
// 输出包含一行一个整数，表示小苯最少需要向多少个旧帐号推荐新账号，如果无法做到，输出-1。

// 这道题是 "和恰好为 x 的最短子序列" 的进阶版，对动态规划递推公式、初始条件、递推更新方向的掌握都有比较高的要求。值得收藏！

func main() {
	input := `5 8 1 2 3 4 10`
	run(strings.NewReader(input), os.Stdout)
	// run(os.Stdin, os.Stdout)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var n, x int
	Fscan(in, &n, &x)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &nums[i])
	}

	ans := solv1(out, nums, x)
	Fprintln(out, ans)
}

func solv2(out io.Writer, nums []int, x int) int {
	n := len(nums)
	mem := make([][][2]int, n+1)
	for i := 0; i <= n; i++ {
		mem[i] = make([][2]int, x+1)
		for j := 0; j <= x; j++ {
			mem[i][j] = [2]int{-1, -1}
		}
	}
	// Fprintln(out, mem)
	var dfs func(int, int, int) int
	dfs = func(i, k, flag int) (ans int) {
		if k == 0 {
			return 0
		}
		if i >= n {
			return n + 1
		}
		if mem[i][k][flag] != -1 {
			return mem[i][k][flag]
		}

		ai := nums[i]
		t1 := n + 1
		if k >= ai/2 {
			t1 = dfs(i+1, k-ai/2, flag) + 1
		}
		t2 := n + 1
		if k >= ai && flag == 0 {
			t2 = dfs(i+1, k-ai, 1) + 1
		}
		ans = min(t1, t2)
		ans = min(ans, dfs(i+1, k, flag))
		mem[i][k][flag] = ans
		return ans
	}
	ans := dfs(0, x, 0)
	// Fprintln(out, mem)
	if ans == n+1 {
		ans = -1
	}
	return ans
}

func solv1(out io.Writer, nums []int, x int) int {
	n := len(nums)
	dp := make([][2]int, x+1)
	for j := 0; j <= x; j++ {
		dp[j][0] = n + 1
		dp[j][1] = n + 1
	}

	dp[0][0] = 0

	for _, num := range nums {
		// for j := num/2; j <= x; j ++ {
		// 因为只能取一次所以这里必须倒着遍历！
		for j := x; j >= num/2; j-- {
			dp[j][0] = min(dp[j][0], dp[j-num/2][0]+1)
			dp[j][1] = min(dp[j][1], dp[j-num/2][1]+1)
			if j >= num {
				dp[j][1] = min(dp[j][1], dp[j-num][0]+1)
			}
		}
	}

	ans := min(dp[x][0], dp[x][1])
	if ans == n+1 {
		ans = -1
	}
	return ans
}
