package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"strings"
)

func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	s := 0
	for i := 0; i < n; i++ {
		Fscan(in, &a[i], &b[i])
		s += a[i]
	}
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	// 初值
	for i := 1; i < n; i ++ {
		f[0][i] = math.MinInt/2
	}

	/// dp
	ans := make([]int, n+1)
	sum := 0
	for i := 1; i <= n; i ++ {
		sum += a[i-1]
		for j := 1; j <= i; j ++ {
			f[i][j] = max(f[i-1][j]+a[i-1], f[i-1][j-1]+b[i-1]-(j^a[i-1]))
			ans[j] = max(ans[j], f[i][j]-sum+s)
		}
	}

	Fprintln(out, s)
	for j := 1; j <= n; j ++ {
		Fprintln(out, ans[j])
	}

}

func main() {
	in := strings.NewReader(`
4
2 4
3 5
2 3
6 4
	`)
	run(in, os.Stdout)
}
