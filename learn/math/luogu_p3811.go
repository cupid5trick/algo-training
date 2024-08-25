package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 题目链接：P3811 【模板】模意义下的乘法逆元 - 洛谷 | 计算机科学教育新生态: https://www.luogu.com.cn/problem/P3811
// 题解：乘法逆元 - zjp_shadow - 博客园: https://www.cnblogs.com/zjp-shadow/p/7773566.html
// 求解逆元可能用到费马小定理。用这种方法可以在线性时间内求解组合数。
// func main() {
// 	run(os.Stdin, os.Stdout)
// }

func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var n, p int
	fmt.Fscan(in, &n, &p)
	f := make([]int, n+1)
	f[1] = 1
	fmt.Fprintln(out, 1)
	for i := 2; i <= n; i ++ {
		f[i] = (p-p/i)*f[p%i] % p
		fmt.Fprintln(out, f[i])
	}
}