package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := `5 8
	1 2 3 4 10`
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

	solv1(out, nums, x)

}

func solv2(out io.Writer, nums []int, x int) {
	
}

func solv1(out io.Writer, nums []int, x int) {
	n := len(nums)
	dp := make([][2]int, x+1)
	for j := 0; j <= x; j++ {
		dp[j][0] = n + 1
		dp[j][1] = n + 1
	}

	dp[0][0] = 0

	for _, num := range nums {
		for j := num/2; j <= x; j ++ {
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
	Fprintln(out, ans)
}