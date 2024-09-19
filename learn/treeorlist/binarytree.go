package main

// 力扣 331. 验证二叉树的前序序列化 - 力扣（LeetCode）: https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/
// 题解地址: https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/solutions/650583/yan-zheng-er-cha-shu-de-qian-xu-xu-lie-h-jghn/
// 方法一：栈
// 我们可以定义一个概念，叫做槽位。一个槽位可以被看作「当前二叉树中正在等待被节点填充」的那些位置。
// 二叉树的建立也伴随着槽位数量的变化。每当遇到一个节点时：
// 如果遇到了空节点，则要消耗一个槽位；
// 如果遇到了非空节点，则除了消耗一个槽位外，还要再补充两个槽位。
// 此外，还需要将根节点作为特殊情况处理。
// 我们使用栈来维护槽位的变化。栈中的每个元素，代表了对应节点处剩余槽位的数量，而栈顶元素就对应着下一步可用的槽位数量。当遇到空节点时，仅将栈顶元素减 1；当遇到非空节点时，将栈顶元素减 1 后，再向栈中压入一个 2。无论何时，如果栈顶元素变为 0，就立刻将栈顶弹出。
// 遍历结束后，若栈为空，说明没有待填充的槽位，因此是一个合法序列；否则若栈不为空，则序列不合法。此外，在遍历的过程中，若槽位数量不足，则序列不合法。
// 复杂度分析
// 时间复杂度：O(n)，其中 n 为字符串的长度。我们每个字符只遍历一次，同时每个字符对应的操作都是常数时间的。
// 空间复杂度：O(n)。此为栈所需要使用的空间。
func isValidSerialization_stack(preorder string) bool {
	n := len(preorder)
	stk := []int{1}
	for i := 0; i < n; {
		if len(stk) == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			stk[len(stk)-1]--
			if stk[len(stk)-1] == 0 {
				stk = stk[:len(stk)-1]
			}
			i++
		} else {
			// 读一个数字
			for i < n && preorder[i] != ',' {
				i++
			}
			stk[len(stk)-1]--
			if stk[len(stk)-1] == 0 {
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, 2)
		}
	}
	return len(stk) == 0
}

// 方法二：计数
// 能否将方法一的空间复杂度优化至 O(1) 呢？
// 回顾方法一的逻辑，如果把栈中元素看成一个整体，即所有剩余槽位的数量，也能维护槽位的变化。
// 因此，我们可以只维护一个计数器，代表栈中所有元素之和，其余的操作逻辑均可以保持不变。

// 复杂度分析
// 时间复杂度：O(n)，其中 n 为字符串的长度。我们每个字符只遍历一次，同时每个字符对应的操作都是常数时间的。
// 空间复杂度：O(1)。
func isValidSerialization_counter(preorder string) bool {
	n := len(preorder)
	slots := 1
	for i := 0; i < n; {
		if slots == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			slots--
			i++
		} else {
			// 读一个数字
			for i < n && preorder[i] != ',' {
				i++
			}
			slots++ // slots = slots - 1 + 2
		}
	}
	return slots == 0
}


// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTargetCopy(original *TreeNode, cloned *TreeNode, target *TreeNode) *TreeNode {
	if original == nil || original == target {
		return cloned
	}
	left := getTargetCopy(original.Left, cloned.Left, target)
	if left != nil {
		return left
	}
	return getTargetCopy(original.Right, cloned.Right, target)
}

/**
 * 2. 两数相加 - 力扣（LeetCode）: https://leetcode.cn/problems/add-two-numbers/description/
 * 处理好进位就行了，解决链表问题常用的哨兵节点，能够写出优雅的解法。
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    a, b, c := 0, 0, 0
    pre := &ListNode{-1, nil}
    ans := pre
    for l1 != nil && l2 != nil {
        a, b = l1.Val, l2.Val
        pre.Next = &ListNode{(a+b+c)%10, nil}
        c = (a+b+c)/10
        pre, l1, l2 = pre.Next, l1.Next, l2.Next
    }

    head := l1
    if head == nil {
        head = l2
    }
    for head != nil {
        a = head.Val
        pre.Next = &ListNode{(a+c)%10, nil}
        c = (a+c)/10
        pre, head = pre.Next, head.Next
    }
    if c != 0 {
        pre.Next = &ListNode{c, nil}
    }
    return ans.Next
}