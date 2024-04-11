package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 一个有向无环图，想要加一条边使得整个图成为连通图。计算加边的方案数。
// 保证输入是非连通图。

// 这是一个路径优化的并查集实现，有机会学学
// 代码随想录: https://programmercarl.com/%E5%9B%BE%E8%AE%BA%E5%B9%B6%E6%9F%A5%E9%9B%86%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%80.htm
func join(fa []int, a, b int) {
	c := find(fa, a)
	d := find(fa, b)
	fa[d] = c
}

func find(fa []int, a int) int {
	if fa[a] == a {
		return a
	} else {
		// 路径压缩
		fa[a] = find(fa, fa[a])
	}
	return fa[a]
}

func main_c() {
	run_c(os.Stdin, os.Stdout)
}

func run_c(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)

	fa := make([]int, n+1)
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fa[i] = i
	}

	for i := 0; i < m; i++ {
		var a, b int
		Fscan(in, &a, &b)
		join(fa, a, b)
	}

	for i := 1; i <= n; i++ {
		res := find(fa, i)
		ans[res]++
	}

	var list []int
	for i := 1; i <= n; i++ {
		if ans[i] != 0 {
			list = append(list, ans[i])
		}
	}

	if len(list) == 2 {
		Fprintln(out, list[0]*list[1])
	} else {
		Fprintln(out, 0)
	}

}
