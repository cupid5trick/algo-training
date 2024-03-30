package main

import (
	. "fmt"
	"io"
	"os"
	"strings"
)

/**
判别回文串
*/

func main() {
	// input := `abazaba`
	input := `ababc`
	run(strings.NewReader(input), os.Stdout)
}

func check(s string) bool {
	n := len(s)

	for i:=0; i < n/2; i ++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func run(rd io.Reader, wt io.Writer) {
	var str string
	Fscan(rd, &str)
	
	ok := check(str)

	Fprintln(wt, ok)

}