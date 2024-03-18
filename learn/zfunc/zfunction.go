package main

import (
	. "fmt"
	"strings"
	"io"
	"bufio"
	"os"
)

func main() {
	input := `abacaba 4`
	run(strings.NewReader(input), os.Stdout)

}

func solution(str string, k int) int {
	n := len(str)
	z := make([]int, n)
	
	l, r := 0, 0
	for i:=1; i<n; i ++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && str[z[i]] == str[i+z[i]] {
			l, r = i, i+z[i]
			z[i] ++
		}
		if i % k == 0 && z[i] >= n-i {
			return i/k
		}
	}
	
	return (n-1)/k +1
}

func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var str string
	var k int
	Fscan(in, &str, &k)
	times := solution(str, k)
	Fprintln(out, times)
}
