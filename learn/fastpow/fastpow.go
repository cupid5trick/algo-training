package main

import (
	"bufio"
	. "fmt"
	"math/big"
	"io"
	"os"
	"strings"
)

const mod = 1000000007

func main() {
	input := `2 1024`
	run(strings.NewReader(input), os.Stdout)
}

func fastpow(base, p, mod int) int {
	ans := 1
	for p > 0 {
		if p & 1 == 1 {
			ans = (ans*base) % mod
		}
		base = (base*base) % mod
		p >>= 1
	}
	return ans
}

func run(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()
	var base, p int
	Fscan(in, &base, &p)

	/*
fastpow:        812734592
math.Pow:       812734592 
	 */
	Fprintf(out, "fastpow:\t%d\n", fastpow(base, p, mod))
	ans := new(big.Int)
	Fprintf(out, "math.Pow:\t%d\n", ans.Exp(big.NewInt(int64(base)), big.NewInt(int64(p)), big.NewInt(mod)))
}