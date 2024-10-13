package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

/*
题目描述：
米小游正在玩《绝区零》。在《绝区零》中有一些关卡，这些关卡形成了一棵以 1 为根的有根树。具体来说，对于第 i 个关卡，必须通过它的前置关卡 f_i，后才能通过第 i 个关卡，其中第 1 个关卡没有前置关卡。

每个关卡都有一个解密值 a_i 和一个操作值 b_i。一个关卡的趣味程度就是解密值与操作值之和。

米小游想知道她通过若干个关卡可以获得的趣味程度之和的最大值是多少。

输入描述：
第一行输入一个整数 n（1 ≤ n ≤ 10^5），表示关卡数量。
第二行输入 n-1 个整数 f_i（1 ≤ f_i ≤ i），表示第 i 个关卡的前置关卡。
第三行输入 n 个整数 a_i（-10^9 ≤ a_i ≤ 10^9），表示第 i 个关卡的解密值。
第四行输入 n 个整数 b_i（-10^9 ≤ b_i ≤ 10^9），表示第 i 个关卡的操作值。
输出描述：
输出一个整数，表示答案，即通过若干个关卡可以获得的趣味程度之和的最大值。

示例1：
输入：

5
1 1 2 2
1 -2 3 -4 5
-1 2 -3 4 -5

输出：

0

*/
func abs(x int) int {
    if x > 0 { return x }
    return -x
}

func max(a, b int) int {
    if a > b { return a}
    return b 
}

func run(rd io.Reader, wt io.Writer) {
    in := bufio.NewReader(rd)
    out := bufio.NewWriter(wt)
    defer out.Flush()

    var n int
    Fscan(in, &n)
    /// 有向树
    g := make([][]int, n)
    for i := 1; i < n; i ++ {
        var fa int
        Fscan(in, &fa)
        fa --
        g[fa] = append(g[fa], i)
    }
    p := make([]int, n )
    for i := 0; i < n; i ++ {
        Fscan(in, &p[i])
    }
    for i := 0; i < n; i ++ {
        var tmp int
        Fscan(in, &tmp)
        p[i] += tmp
    }
    // Fprintln(out, g)
    // Fprintln(out, p)

    /// dfs
    
    var dfs func(int, int) int
    dfs = func(i, fa int) int {
        // Fprintln(out, i, fa, ans, max(ans+p[i], 0))
        res := p[i]
        for _, j := range g[i] {
            if j != fa {
                tmp := dfs(j, i)
                res += max(tmp, 0)
            }
        }
        return res
    }
    res := dfs(0, -1)
    Fprintln(out, max(0, res))
}

func main() {
    run(os.Stdin, os.Stdout)
}