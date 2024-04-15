package main

import (
	"io"
	"math/rand"
	"strings"
	"testing"
	. "fmt"
	. "github.com/cupid5trick/algotrain/learn/testutil"
)


func inputGenerator() string {
	// 生成两个100以内的随机数 n 和 x
	n := rand.Intn(100) + 1
	x := rand.Intn(100) + 1

	// 生成 n 个随机整数并构建输出字符串
	var sb strings.Builder
	sb.WriteString(Sprintf("%d %d\n", n, x)) // 第一行是 n 和 x
	for i := 0; i < n; i++ {
		num := rand.Intn(100) + 1 // 生成一个100以内的随机整数
		sb.WriteString(Sprintf("%d ", num)) // 添加到输出字符串中
	}
	sb.WriteString("\n") // 添加换行符
	return sb.String()
}

func TestProblemB(t *testing.T) {
	
	var run1 IoFunc
	var run2 IoFunc

	var runner func(func(io.Writer, []int, int) int) IoFunc
	runner = func(solver func(io.Writer, []int, int) int) IoFunc {
		return func(in io.Reader, out io.Writer) {
			var n, x int
			Fscan(in, &n, &x)
			nums := make([]int, n)
			for i := 0; i < n; i++ {
				Fscan(in, &nums[i])
			}

			ans := solver(out, nums, x)
			Fprintln(out, ans)
			}
	}

	run1 = runner(solv1)
	run2 = runner(solv2)

	AssertEqualRunResultsInf(t, inputGenerator, run2, run1)
}

