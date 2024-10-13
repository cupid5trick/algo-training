package main

import "sort"

/*
- 3319. 第 K 大的完美二叉子树的大小: <https://leetcode.cn/problems/k-th-largest-perfect-subtree-size-in-binary-tree/description/>

> 完美二叉树 是指所有叶子节点都在同一层级的树，且每个父节点恰有两个子节点。
上面完美二叉树的定义可以转化为左右子树的大小相等，并且各自的左右子树也满足左右子树大小相等的条件。
可以通过后序遍历来求解。

```python
class Solution:
    def kthLargestPerfectSubtree(self, root: Optional[TreeNode], k: int) -> int:
        q = []
        def dfs(root) -> int:
            if root == None:
                return 0
            left = dfs(root.left)
            right = dfs(root.right)
            if left == right :
                if len(q) < k and left+right+1>0:
                    heapq.heappush(q, left+right+1)
                elif len(q) == k and left+right+1 >= q[0]:
                    heapq.heappop(q)
                    heapq.heappush(q, left+right+1)
                    # print(q)
                return left+right+1
            return -1
        dfs(root)
        return q[0] if len(q) == k else -1
```
*/

//lint:file-ignore U1000 unused function
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthLargestPerfectSubtree(root *TreeNode, k int) int {
    var dfs func(*TreeNode) int
    ans := []int{}
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        left := dfs(root.Left)
        right := dfs(root.Right)
        x := left+right+1
        if left == right && x > 0 {
            ans = append(ans, x)
            return x
        }
        return -1
    }
    dfs(root)
    if len(ans) < k {
        return -1
    }
    sort.Ints(ans)
    return ans[len(ans)-k]
}