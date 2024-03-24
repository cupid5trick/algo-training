package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

/**
输入一个数组 nums，1<=nums[i]<=10^9，还有 n, m, k, d 四个整数 (n, k [1,10^6], m, d [0,10^6])
Alice, Bob 两个人依次操作一次，Alice 操作目的是为了让数组和最大，Bob是为了让数组和最小：
- Alice 先操作，从 nums 删除最多 d 个数，
- Bob 从剩余的数中选择最多 m 个乘上 -k。
现在假设两个人都足够聪明，请求出操作结束后数组的和。

给了4个用例：
4
3 1 1 1
4 1 1
3 1 1 1
4 3 2
5 4 2 0
3 5 1 4 1
10 4 1 6
1 8 2 9 3 3 4 5 3 200

0
1
-25
-9
*/

// func main() {
//     run(os.Stdin, os.Stdout)
// }

func run_a(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()
	
	var n, m, k, d int
	Fscan(in, &n, &m, &k, &d)

	ans := int64(math.MinInt64)
	
    Fprintln(out, ans)
}
