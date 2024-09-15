package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

/*
 type ListNode struct{
   Val int
   Next *ListNode
 }
*/

/**
代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
*

@param head ListNode类
@param m int整型
@param n int整型
@return ListNode类
*/

func main() {
	// run(os.Stdin, os.Stdout)
	input := `2 4 1 2 3 4 5`
	run(strings.NewReader(input), os.Stderr)
}

func run(rd io.Reader, wt io.Writer) {
	in := bufio.NewReader(rd)
	out := bufio.NewWriter(wt)
	defer out.Flush()

	dummy := &ListNode{Val: -1, Next: nil}
	node := dummy
	var m, n int
	Fscan(in, &m, &n)
	for true {
		var num int
		flag, err := Fscan(in, &num)
		if flag != 0 || err != nil {
			break
		}
		node.Next = &ListNode{Val: num, Next: nil}
	}
	
	res := reverseBetween(dummy.Next, m, n)
	Fprintln(out, res)
}

type ListNode struct{
	Val int
	Next *ListNode
}

func reverseBetween( head *ListNode ,  m int ,  n int ) *ListNode {
	
	dummy := &ListNode{Val: -1, Next: head}

	pre := dummy
	/// 0-> 1->2->3->4->5
	for node, cnt:=head, 1; node != nil; cnt ++ {
		if cnt < m {
			pre = node
		} else if cnt == n {
			reverse(pre, node.Next)
		}
		node = node.Next
	}
	return dummy.Next
}

// ... -> pre ->head->...->pretail->tail
// ... -> pre ->pretail->...->head->tail
func reverse(pre, tail *ListNode) *ListNode {
	dummy, head := pre, pre.Next
	node, pre := head, tail
	for node != tail {
		node, node.Next, pre = node.Next, pre, node
	}
	dummy.Next = pre
	return head
}