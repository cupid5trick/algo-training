package main

import (
	"fmt"
	"math"
)

// 155. 最小栈 - 力扣（LeetCode）: https://leetcode.cn/problems/min-stack/
//  设计一个支持 min 函数的栈数据结构，实现 min、push、pop方法

type MinStack struct {
	s    []int
	minv []int
}

func Constructor() MinStack {
	return MinStack{s: []int{}, minv: []int{math.MaxInt64}}
}

func (st MinStack) GetMin() (x any) {
	if len(st.minv) == 0 {
		return nil
	}
	return st.minv[len(st.minv)-1]
}

func (st *MinStack) Push(x any) {
	m := len(st.minv)
	st.minv = append(st.minv, min(st.minv[m-1], x.(int)))
	st.s = append(st.s, x.(int))
}

func (st *MinStack) Pop() (x any) {
	st.minv = st.minv[:len(st.minv)-1]
	x = st.s[len(st.s)-1]
	st.s = st.s[:len(st.s)-1]
	return
}

func (st *MinStack) Top() (x any) {
	if len(st.s) > 0 {
		x = st.s[len(st.s)-1]
	}
	return
}

func main() {
	st := Constructor()
	st.Push(3)
	st.Push(2)
	st.Push(1)
	st.Push(1)
	fmt.Println(st.GetMin())
	st.Pop()
	fmt.Println(st.GetMin())
	st.Pop()
	fmt.Println(st.GetMin())
	st.Pop()
	fmt.Println(st.GetMin())
	st.Pop()
	fmt.Println(st.GetMin())
	fmt.Println(st.Top())
}
