package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)
/*
### **题目描述**

玩家得到了一个数组，允许从中取走一些数，使得最终所有取走的数乘积的末尾包含尽可能多的 `0`。要求 **取的数字必须不相邻**。
问最多能取出多少个数字，使得它们的乘积的末尾有最多的 `0`？

---

### **输入描述**

-   第一行输入一个整数 `n`，表示数组的长度。
-   第二行输入 `n` 个正整数，表示玩家手中的数组。

#### **输入限制**

-   1≤n≤1000
-   1≤ai≤109

---

### **输出描述**

-   输出一个整数，表示最多可以取得的数字个数，使得乘积的末尾包含最多的 `0`。

---

### **示例**

#### **输入**

```
4
25 30 125 64
```

#### **输出**

```
2
```
*/
func count(x, b int) int {
	ans := 0
	for x%b == 0 && x != 0 {
		ans++
		x /= b
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func solve(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	nums_2 := make([]int, n)
	nums_5 := make([]int, n)
	f1 := 0
	f2 := 0

	for i := 0; i < n; i ++ {
		var x int
		Fscan(in, &x)
		nums_2[i] = count(x, 2)
		nums_5[i] = count(x, 5)
	}

	ans := 0
    /// g(i) = max(g(i-1), min(f2+nums_2[i], f5+nums_5[i]))
	for i := 0; i < n; i++ {
		tmp := max(f1, f2)
		f2 = max(f2, f1+min(nums_2[i], nums_5[i]))
		f1 = tmp
		ans = max(ans, max(f1, f2))
	}
    Fprintln(out, ans)
    
}
func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	solve(in, out)
}
func main() {
	run(os.Stdin, os.Stdout)
}
