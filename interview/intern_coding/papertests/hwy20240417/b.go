package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	// "strings"
)

// type info struct {n1, n2 int}


func main() {
	
	run(os.Stdin, os.Stdout)
}
/*
40 12
a * 0 2
a * 1 2
b a 0 3
b a 1 5
c a 1 3
d a 0 1
d a 1 3
e b 0 2
f * 0 8
f * 1 10
g f 1 2
h * 0 4
*/ 
func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()
	
	var m, n int
	Fscan(in, &m, &n)
	var t string
	Fscan(in, &t)
	children := map[string][]string{}
	stat1 := map[string]int{}
	stat2 := map[string]int{}
	// 无父节点为云服务
	cloud := map[string]bool{}
	for i:=0; i < n; i ++ {
		var a, b string
		var c, d int
		Fscanf(in, "%s %s %d %d\n", &a, &b, &c, &d)
		Fprintln(out, a, b, c, d)
		if b != "*" {
			children[b] = append(children[b], a)
		} else {
			cloud[a] = true
		}
		if c == 0 {
			stat1[a] = d
		} else {
			stat2[a] = d
		}
	}
	ans := 0
	var dfs func(string) [2]int
	dfs = func(i string) (res [2]int) {
		
		for _, j := range children[i] {
			sub := dfs(j)
			res[0] += sub[0]
			res[1] += sub[1]
		}
		res[0] += stat1[i]
		res[1] += stat2[i]
		if 5*res[0]+2*res[1] > m {
			ans ++
		}
		return res

	}
	for k, v := range cloud {
		if v {
			dfs(k)
		}
	}

	Fprintln(out, ans)


}