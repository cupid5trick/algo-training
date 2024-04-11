package main

import "fmt"

// 并查集的路径压缩实现
// 并查集是一个树形结构算法，用来解决连通性问题：判断两个元素是否在同一个集合。
// 用一个一维数组 father 来表示连通关系，例如将三个节点 ABC 连通为 A=>B=>C。
// 即：father[A] = B，father[B] = C 这样就表述 A 与 B 与 C连通了（有向连通图）
// 代码随想录: https://programmercarl.com/%E5%9B%BE%E8%AE%BA%E5%B9%B6%E6%9F%A5%E9%9B%86%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%80.htm
type unionfind []int

func (uf *unionfind) union(u, v int) {
	u = uf.find(u)
	v = uf.find(v)
	if u != v {
		(*uf)[v] = u
	}
}

func (uf *unionfind) find(a int) int {
	fa := *uf
	if a == fa[a] {
		return a
	} else {
		// 路径压缩
		fa[a] = uf.find(fa[a])
	}
	return fa[a]
}

func main() {
	uf := make(unionfind, 10)
	for i := range uf {
		uf[i] = i
	}
	uf.union(2, 1) // 1 => 2
	uf.union(4, 3) // 3 => 4
	uf.union(6, 5) // 5 => 6
	uf.union(8, 7) // 7 => 8

	fmt.Println("Root of 1:", uf.find(1)) // Should print 2
	fmt.Println("Root of 3:", uf.find(3)) // Should print 4

	uf.union(4, 2)                        // 2 => 4
	fmt.Println("Root of 1:", uf.find(1)) // Should print 4
	fmt.Println("Root of 5:", uf.find(5)) // Should print 6
	fmt.Println("Root of 7:", uf.find(7)) // Should print 8
	fmt.Println(uf)
}
