package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// 小红非常喜欢一首叫做剪切线的歌，所以小红现在非常喜欢剪切各种东西。
// 现在小红有一个数组，她可以按任意顺序排序这个数组，然后她把这个数组剪切成两个部分，她想知道两个部分的和的GCD 最大是多少。
// GCD:最大公约数(Greatest Common Divisor)。
// 输入描述
// 第一行输入一个整数 n(1 < n < 1000) 表示数组长度。
// 第二行输入n 个正整数 a(1 ≤ ai ≤ 1000)表示数组。
// 输出描述
// 输出一个整数表示答案。
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func solver1(out io.Writer, arr []int) {
	n := len(arr)
    maxSum := 0
    for _, v := range arr {
        maxSum += v
    }

    dp := make([][2]int, n+1)
    dp[0][0] = -1

    for i, num := range arr {
        for j := i-1; j > 0; j-- {
			g1 := gcd(dp[j][1], maxSum - dp[j][1])
			g2 := gcd(dp[j-1][1] + num, maxSum - dp[j-1][1] - num)
            dp[j][0] = max(g1, g2)
			Fprintln(out, dp, g1, g2)
			if g2 > g1 {
				dp[j][1] = dp[j-1][1] + num
			}
        }
    }

    res := dp[n][0]
    
    Fprintf(out, "%d\n", dp)
    Fprintf(out, "%d\n", res)
}

func run(rd io.Reader, wt io.Writer) {
    in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

    var n int
    Fscan(in, &n)

    arr := make([]int, n)
    for i := 0; i < n; i++ {
        Fscan(in, &arr[i])
    }

    solver1(out, arr)
}

func main() {
	input := `3 1 2 3`
	run(strings.NewReader(input), os.Stdout)
    run(os.Stdin, os.Stdout)
}
