package main

import (
	"bufio"
	. "fmt"
	"io"
    "os"
)

/*
长度为 n 的字符串删除 m 个字符，使得剩余字符串的字典序最小，输出字典序最小的字符串。
每个测试点有 t 个样例。
*/
func solve_b(in io.Reader, out io.Writer) {
    var n, m int
    Fscan(in, &n, &m)
    var s string
    Fscan(in, &s)
    stk := []rune{}
    for _, c := range s {
        for len(stk) > 0 && m > 0 && stk[len(stk)-1] > c {
            stk = stk[:len(stk)-1]
            m --
        }
        stk = append(stk, c)
    }
    if m > 0{ stk = stk[:len(stk)-m] }
    Fprintln(out, string(stk))
}
func run_b(rd io.Reader, wt io.Writer) {
    in := bufio.NewReader(rd)
    out := bufio.NewWriter(wt)
    defer out.Flush()

    var t int
    Fscan(in, &t)
    for t > 0 {
        t --
        solve(in, out)
    }
}
// func main() {
//     run(os.Stdin, os.Stdout)
// }