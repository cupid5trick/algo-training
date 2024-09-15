package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main_a() {
	run(os.Stdin, os.Stdout)
}

func run_a(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	first := map[string]struct{}{}
	for n > 0 {
		n--
		var id string
		Fscan(in, &id)
		if _, ok := first[id]; !ok {
			first[id] = struct{}{}
			Fprintln(out, id)
		}
	}
	

}
