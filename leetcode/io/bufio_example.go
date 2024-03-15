package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	n, q := 0, 0
	Fscan(in, &n, &q)
	sum := int64(0)
	cnt := 0
	for i := 0; i < n; i++ {
		num := 0
		Fscan(in, &num)
		if num != 0 {
			sum += int64(num)
		} else {
			cnt++
		}
	}

	for i := 0; i < q; i++ {
		var low int64
		var high int64
		Fscan(in, &low, &high)
		Fprintf(out, "%d %d\n", sum+low*int64(cnt), sum+high*int64(cnt))
	}
}
