package main

// 小红拿到了一个链表。她准备将这个链表断裂成两个链表，再连接到一起，使得链表从头节点到尾部升序。
// 你能帮小红达成目的吗？
// 给定的为一个链表数组，你需要对数组中每个链表进行一次 "是" 或 "否" 的回答，并返回布尔数组
// 每个链表的长度不小于2，且每个链表中不包含两个相等的元素。保证所有链表长度之和不超过 10^5

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func canSorted(lists []*ListNode) []bool {
	n := len(lists)
	res := make([]bool, n)
	for i := 0; i < n; i++ {
		head := lists[i]
		p := lists[i]
		count := 0
		for p != nil && p.Next != nil {
			if p.Next.Val < p.Val {
				count++
			}
			p = p.Next
		}
		if count > 1 || (count == 1 && p.Val > head.Val) {
			res[i] = false
		} else {
			res[i] = true
		}
	}
	return res
}

func main_b() {

	// Example usage
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	lists := []*ListNode{list1, list2}

	result := canSorted(lists)
	fmt.Println(result) // Output: [true true]
}
