package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

type pair struct {
	x, y int
}

func solve_dfs(grid [][]byte, m, n int) int {
	// 定义目标字符串
	target := "tencent"
	// 初始化路径数量
	count := 0

	// 定义DFS函数，用于搜索路径
	var dfs func(x, y, index int)
	dfs = func(x, y, index int) {
		// 如果当前位置超出了网格范围，直接返回
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		// 如果当前字符不匹配，直接返回
		if grid[x][y] != target[index] {
			return
		}
		// 如果已经匹配到了目标字符串的最后一个字符，增加路径数量
		if index == len(target)-1 {
			count++
			return
		}
		// // 将当前位置的字符标记为已访问
		// temp := grid[x][y]
		// grid[x][y] = ' '
		// 向四个方向继续搜索
		dfs(x+1, y, index+1)
		dfs(x-1, y, index+1)
		dfs(x, y+1, index+1)
		dfs(x, y-1, index+1)
		// // 恢复当前位置的字符
		// grid[x][y] = temp
	}

	// 遍历整个网格，以每个位置为起点进行DFS搜索
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, 0)
		}
	}

	return count
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	Printf("Alloc = %v MiB", bToMb(m.Alloc))
	Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	Printf("\tSys = %v MiB", bToMb(m.Sys))
	Printf("\tNumGC = %v\n", m.NumGC)
    // Println(m)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	printMemUsage()
	// m, n := 3, 3
	// grid := [][]byte{
	//     {'t', 'e', 'n'},
	//     {'n', 't', 'c'},
	//     {'e', 'n', 'e'},
	// }
	// // 设置随机数种子
	// rand.Seed(time.Now().UnixNano())
	randx := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成一个1000*1000的网格，并填充随机字符
    // Go语言 12000*12000 的回溯都不爆栈啊?
	size := 200
	m, n := size, size
	grid := make([][]byte, m)
	table := "tencentzzzzzzz"
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := range grid[i] {
			grid[i][j] = table[randx.Intn(len(table))] // 填充随机字符
		}
	}

	var input string
	input += Sprintf("%d %d\n", m, n)
	for _, row := range grid {
		input += string(row) + "\n"
	}

	printMemUsage()
	start2 := time.Now()
	run_e(strings.NewReader(input), os.Stderr, solve_dfs)
	duration2 := time.Since(start2)

	Fprintln(os.Stdout, "dfs method: ", duration2)
	printMemUsage()
}

func run_e(rd io.Reader, wt io.Writer, solver func([][]byte, int, int) int) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	var m, n int
	Fscan(in, &m, &n)
	grid := make([][]byte, m)
	for i := 0; i < m; i++ {
		var row string
		Fscan(in, &row)
		grid[i] = []byte(row)
	}

	// 调用函数计算路径数量并输出结果
	Fprintln(out, "路径数量:", solver(grid, m, n))

}
