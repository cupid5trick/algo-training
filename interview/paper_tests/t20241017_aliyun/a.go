package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*

*/
func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var n, k int
		fmt.Fscan(in, &n, &k)
		nums := make([]int, n)
		diff := make([]int, n+1)
		queries := make([][2]int, k)

		// 读取并存储 nums[i]&1 的值
		for i := 0; i < n; i++ {
			var x int
			fmt.Fscan(in, &x)
			nums[i] = x & 1
		}

		// 读取操作并使用差分数组记录修改位置
		for i := 0; i < k; i++ {
			var l, r int
			fmt.Fscan(in, &l, &r)
			queries[i] = [2]int{l - 1, r - 1}
			diff[l-1]++
			diff[r]--
		}

		// 更新 nums 数组并计算前缀和
		s := 0
		pre := make([]int, n+1)
		for i, x := range diff[:n] {
			s += x
			nums[i] ^= s & 1
			if i < n-1 {
				pre[i+1] = pre[i] + (nums[i] ^ nums[i+1])
			}
		}

		// 检查函数
		check := func(i, j, x int) bool {
			if i < 0 || j < 0 {
				return true
			}
			return pre[j+1]-pre[i] == x
		}

		// 查询结果
		ans := false
		for _, q := range queries {
			l, r := q[0], q[1]
			cur := check(0, l-2, l-1) && check(l, r-1, r-l) && check(r+1, n-2, n-r-2)
			if l > 0 {
				cur = cur && (nums[l-1]^nums[l] == 0)
			}
			if r < n-1 {
				cur = cur && (nums[r]^nums[r+1] == 0)
			}
			if cur {
				ans = true
				break
			}
		}
		if ans {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func main() {
	in := strings.NewReader(`
2
5 2
2 1 3 4 5
2 4
1 3
3 2
1 3 4
2 2
1 1
	`)
	run(in, os.Stdout)
}
