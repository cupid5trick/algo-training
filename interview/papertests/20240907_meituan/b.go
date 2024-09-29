package main

import (
	"bufio"
	. "fmt"
	"os"
	// "strconv"
)

var (
	tree             [][]int          // 邻接表表示的树
	childCountMap    map[int]int      // 用于记录直接子节点数量出现次数的哈希表
	totalSimilarPairs int             // 总的相似节点对数
)

func dfs(node, parent int) {
	childCount := 0

	for _, neighbor := range tree[node] {
		if neighbor == parent {
			continue // 跳过父节点
		}
		childCount ++
        dfs(neighbor, node)
	}

	// 记录每个节点的直接子节点数量的出现次数
	childCountMap[childCount] = childCountMap[childCount] + 1

}

func calculateSimilarNodePairs(root int) int {
	dfs(root, -1)

	// 根据哈希表计算相似节点对数
	for _, count := range childCountMap {
		if count > 1 {
			totalSimilarPairs += count * (count - 1) / 2 // C(k, 2)
		}
	}

	return totalSimilarPairs
}

/**
1 7 1 2 1 3 3 5 3 7 2 4 2 6

*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

    var t int 
    Fscan(reader, &t)

	for t > 0 {
		t--

		// 读取节点数
		var n int
        Fscan(reader, &n)
		// 初始化树和相关变量
		tree = make([][]int, n)
		childCountMap = make(map[int]int)
		totalSimilarPairs = 0

		for i := range tree {
			tree[i] = []int{}
		}

		// 读取每条边并构建邻接表
		for i := 0; i < n-1; i++ {
			edges := make([]int, 2)
			Fscan(reader, &edges[0], &edges[1])
			u, v := edges[0]-1, edges[1]-1 // 输入的节点编号从1开始，所以需要减1

			tree[u] = append(tree[u], v)
			tree[v] = append(tree[v], u) // 双向图
		}

		// 计算相似节点对数
		calculateSimilarNodePairs(0)

		// 输出结果
        Fprintln(writer, childCountMap)
		Fprintln(writer, totalSimilarPairs)
	}
}
