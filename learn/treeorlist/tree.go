package main

import (
	"fmt"
	"math/bits"
)

// 1483. 树节点的第 K 个祖先 - 力扣（LeetCode）: https://leetcode.cn/problems/kth-ancestor-of-a-tree-node
// 学一个花式技巧：树上倍增
// 如果是暴力的做法，每次不断从parent数组求当前 node 的第 k 个祖先需要 O(k) 的复杂度
// 但是我们可以对查询 (node, k) 的 k 做一个数位分解的优化：
// 我们如果能在 O(1) 的复杂度求出 (node, 2^j)，那么就能用 O(logk)的复杂度求出 (node, k)
// 所以我们需要做一个预处理，求出每个 node 的第 1, 2, 4, ..., 2^log(n) 个祖先节点
// 这也就是我们定义一个 fa[n][m] 数组的目的
// 1483. 树节点的第 K 个祖先 - 力扣（LeetCode）: https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solutions/2305895/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw
type TreeAncestor struct {
	parent []int
	fa     [][]int
}

func Constructor(n int, parent []int) TreeAncestor {
	solver := TreeAncestor{parent, make([][]int, n)}
	m := bits.Len(uint(n))

	for i := range solver.fa {
		solver.fa[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			solver.fa[i][j] = -1
		}
		solver.fa[i][0] = parent[i]
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			if t := solver.fa[i][j-1]; t != -1 {
				solver.fa[i][j] = solver.fa[t][j-1]
			} else {
				solver.fa[i][j] = -1
			}
		}
	}

	fmt.Println(solver.fa)

	return solver
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	for k > 0 && node != -1 {
		idx := k & -k
		node = this.fa[node][bits.TrailingZeros(uint(idx))]
		k &= (k - 1)
	}
	return node
}
