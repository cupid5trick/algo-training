package main
import (
	"bufio"
	. "fmt"
	"io"
	"os"
)
// 小红拿到一个无向图，其中一些边被染成了红色。
// 小红定义一个点是“好点”，当且仅当这个点的所有邻边都是红边。请你求出这个无向图上“好点”的数量。
// 如果一个点没有任何邻边，那么它也是好点

func main_a() {
	run_a(os.Stdin, os.Stdout)
}

func run_a(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	ok := make([]bool, n)

	for i:=0; i < m; i++ {
		var u, v int
		var c string
		Fscan(in, &u, &v, &c)
		ok[u] = true
		ok[v] = true
	}

	ans := n
	for _, flag := range ok {
		if flag {
			ans --
		}
	}
	
	Fprintln(out, ans)
}
