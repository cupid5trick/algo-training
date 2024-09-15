package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

/**
输入两个只包含0和1的字符串A,B，长度分别是m,n（m>n）。有一个字符串构造规则如下，根据这个规则生成一个长度为n的字符串 S：
- Si = A[x+i] ^ B[i], x [0,m-n]
- S0, S1, ..., Sn-1 的异或值为 0

问：能构造出多少个不同的S？(n [1,10^6])
*/
type set map[string]struct{}

func (st *set) add(str string) {
	(*st)[str] = struct{}{}
}

func (st set) exists(str string) bool {
	_, ok := st[str]
	return ok
}

// func main() {
//     run(os.Stdin, os.Stdout)
// }

func run_b(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var m, n int
	var str, target string
	Fscan(in, &m, &n)
	Fscan(in, &str, &target)
    
    ans := 0
	hset := set{}
	for i := 0; i < m-n+1; i++ {
		s := str[i : i+n]
		sg := make([]byte, n)
		for j, ch := range []byte(s) {
			sg[j] = (ch - '0') ^ (target[j] - '0') + '0'
		}
		s_new := string(sg)
		if !hset.exists(s_new) {
			hset.add(s_new)
		} else {
            continue
        }
        
        xor := byte(0)
        for _, ch := range []byte(s_new) {
            xor ^= (ch-'0')
        }
        if xor == 0 {
            ans ++
        }
	}

    Fprintln(out, ans)
}
