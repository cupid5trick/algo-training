package main

// import (
// 	"bufio"
// 	. "fmt"
// 	"io"
// 	"os"
// )

// /*
// 图中描述的是一个关于《原神》游戏抢票的问题，具体内容如下：

// 1. **背景**：
//    - 米小游是一位《原神》的全勤玩家，但未能抢到《原神FES》的门票，因此感到非常伤心。
//    - 她决定开发一款有利于全勤玩家的抢票系统。

// 2. **抢票系统规则**：
//    - 票分为两个档位，每个档位的票数都为 m。
//    - 游戏运营了 n 天。
//    - 设置了一个抢票参数 t。
//    - 如果抢票玩家的游戏登录天数 x \geq t，则优先分配第1档位的票。如果第1档位没有票了，则分配第2档位的票。如果第2档位也没有票，则该玩家没有抢到票。
//    - 如果 x < t，则直接分配第2档位的票。如果第2档位没有票，则该玩家没有抢到票。

// 3. **米小游的情况**：
//    - 米小游是全勤玩家，登录天数为 n。
//    - 现在有 q 个玩家在和米小游抢票。
//    - 第 j 个玩家的登录天数为 d_i。
//    - 抢票的先后顺序可以看作是一个长度为 q+1 的排列，但具体的排列未知。

// 4. **问题**：
//    - 想知道有多少种排列可以使得米小游至少抢到一张票。

// 5. **输入描述**：
//    - 第一行输入四个整数 n, m, q（1≤ n, m, q ≤ 10^3），t（1≤ t ≤ n），表示游戏运营天数，每个档位的票数、抢票玩家数，抢票参数。
//    - 第二行输入 q 个整数 d_i（1≤ d_i ≤ n），表示玩家的游戏登录天数。

// 6. **输出描述**：
//    - 输出一个整数，表示满足条件的排列数量。

// */
// const M int = 1000000007

// var inv []int

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func fastpow(x, b int) int {
// 	ans := 1
// 	for b > 0 {
// 		if b&1 == 1 {
// 			ans = (ans * x) % M
// 		}
// 		b >>= 1
// 		x = x * x % M
// 	}
// 	return ans
// }

// ///
// func invm(n int) {
// 	nfac := 1
// 	for i := 2; i <= n; i++ {
// 		nfac = nfac * i % M
// 	}
// 	inv[n] = fastpow(nfac, M-2)
// 	for i := n - 1; i >= 1; i-- {
// 		inv[i] = inv[i+1] * (i + 1) % M
// 	}
// }

// func A(s, t int) int {
// 	res := 1
// 	inv[0] = 1
// 	for x := 2; x <= s; x++ {
// 		res = (res * x) % M
// 	}
// 	res = res * inv[s-t] % M
// 	Fprintf(os.Stdout, "A(%d,%d):%d\n", s, t, res)
// 	return res
// }

// func run(rd io.Reader, wt io.Writer) {
// 	in := bufio.NewReader(rd)
// 	out := bufio.NewWriter(wt)
// 	defer out.Flush()

// 	var n, m, q, t int
// 	Fscan(in, &n, &m, &q, &t)
// 	inv = make([]int, n+1)
// 	invm(n)
// 	l := 1
// 	for i := 0; i < q; i++ {
// 		var x int
// 		Fscan(in, &x)
// 		if x >= t {
// 			l++
// 		}
// 	}
// 	q++
// 	o := q - l
// 	// rk <= 2m
// 	res := A(q-1, q-2*m) * A(2*m, 2*m) % M
// 	// rk > 2m
// 	for rk := 2*m + 1; rk <= q && l >= m && o >= rk-m; rk++ {
// 		tmp := A(l-1, l-m) * A(m, m) % M
// 		tmp = tmp * A(o, o-(rk-m)) % M
// 		tmp = tmp * A(rk-m, rk-m) % M
// 		res = (res + tmp)%M
// 	}
// 	Fprintln(out, res)
// }

// func main() {
// 	run(os.Stdin, os.Stdout)
// }
