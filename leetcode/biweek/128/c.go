package main

// Dijkstra 算法
// 3112. 访问消失节点的最少时间 - 力扣（LeetCode）: https://leetcode.cn/problems/minimum-time-to-visit-disappearing-nodes/description/
// 学习一下 Dijkstra 算法。Dijkstra 算法是维护一个没有计算出最短路径的节点集合，每次从中选出一个距离最短的节点作为当前更新的起点。
// 对于每一步选出来的起点，会遍历该节点的所有邻居节点，更新其余未计算出最短路的节点的距离。
// 当然初始条件就是 dist[0] = 0 (0 号节点作为起点)
// 解题思路：
// 本题实际上就是 Dijkstra 算法的考查，唯一要注意的就是如果消失时间之前还未到达或这一刻刚好到达，不应该更新这条路径。
// 只要在更新距离的部分增加一个条件即可 (dist[v] < t[v])

type pair struct{ u, w int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].w < h[j].w }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.(pair)) }
func (h *hp) Pop() (x any)      { old := *h; *h, x = old[:len(old)-1], old[len(old)-1]; return }

func minimumTime(n int, edges [][]int, t []int) []int {
	g := make([][]pair, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], pair{v, w})
		g[v] = append(g[v], pair{u, w})
	}

	dist := make([]int, n)
	for i := range dist {
		if i > 0 {
			dist[i] = -1
		}
	}

	h := hp{{}}

	for len(h) > 0 {
		top := heap.Pop(&h).(pair)
		x, dx := top.u, top.w
		if dx > dist[x] {
			continue
		}

		for _, adj := range g[x] {
			y, dy := adj.u, dx+adj.w
			if dy < t[y] && (dist[y] < 0 || dy < dist[y]) {
				heap.Push(&h, pair{y, dy})
				dist[y] = dy
			}
		}
	}

	return dist
}