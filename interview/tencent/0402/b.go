package main

// 二叉树的最大宽度：每层最左边非 null 的节点到最右边非 null 的叶子节点的距离
// 中间的 null 节点也计入
import (
	. "fmt"
)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelWidth(root *TreeNode) (ans int) {
	q := []*TreeNode{}

	q = append(q, root)

	for len(q) > 0 {
		n := len(q)
		next := []*TreeNode{}
		cur := n
		for node := range q {
			next = append(next, node.Left)
			next = append(next, node.Right)
		}

		
		for i:=0; i < n && q[i] == nil; i ++ {
			cur --
		}
		for i:=n-1; i >= 0 && q[i] == nil; i -- {
			cur --
		}
		ans = max(ans, cur)
	}

	return
}

func main_b() {
	// 构造测试二叉树
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 8,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 9,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
	}

	Println(maxLevelWidth(root))

}
