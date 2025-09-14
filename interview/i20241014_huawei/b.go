package main

import "fmt"

// - 946. 验证栈序列 - 力扣（LeetCode）: <https://leetcode.cn/problems/validate-stack-sequences/description/>

func main() {
	nums := []byte{'a', 'b', 'c', 'd'}
	output := []byte{'a', 'c', 'b', 'd'}

	stk := []byte{}
	j := 0
	for _, x := range nums {
		/// 入栈
		stk = append(stk, x)
		/// 出栈
		for len(stk) > 0 && stk[len(stk)-1] == output[j] && j < len(output) {
			j++
			stk = stk[:len(stk)-1]
		}
	}

	/// 合法
	if len(stk) == 0 && j == len(output) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}