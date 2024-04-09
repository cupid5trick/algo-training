package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 160. 相交链表 - 力扣（LeetCode）: https://leetcode.cn/problems/intersection-of-two-linked-lists
// 脑洞题。有点像莫比乌斯环。。。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    n1, n2 := headA, headB

    if n1 == nil || n2 == nil {
        return nil
    }
    
    for n1 != n2 {
        if n1 != nil {
            n1 = n1.Next
        } else {
            n1 = headB
        }
        if n2 != nil {
            n2 = n2.Next
        } else {
            n2 = headA
        }
    }

    return n1

}