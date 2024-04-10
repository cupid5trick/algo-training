package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	input := `1 7 3 2 2 2 1 1 1 1`
	run(strings.NewReader(input), os.Stdout)
}

func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var t int
	Fscan(in, &t)

	for t > 0 {
		var n, k int
		Fscan(in, &n, &k)
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &nums[i])
		}

		ans := solv(nums, k, out)
		Fprintln(out, ans)
		t--

	}
}

func solv(nums []int, k int, out io.Writer) int {
	counter := map[int]int{}
	for _, x := range nums {
		counter[x]++
	}
	cnt := []int{}
	for _, v := range counter {
		cnt = append(cnt, v)
	}

	sort.Ints(cnt)

	m := len(cnt)

	pre := make([]int, m+1)
	for i := 0; i < m; i++ {
		pre[i+1] = pre[i] + cnt[i]
	}

	ans := -1
	for value, c := range counter {
		r := sort.SearchInts(cnt, c)
		presum := c * (m - r)
		if r < m {
			presum += pre[r]
		}
		if presum > k {
			ans = max(ans, value)
		}
	}
	return ans
}
