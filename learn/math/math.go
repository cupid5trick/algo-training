package main

import "fmt"

// 1. 最大公约数算法：辗转相除法
// 辗转相除法 递归版本
func gcd_recursive(a, b int) int {
    if b > a {
        return gcd_recursive(b, a)
    }
    if a % b == 0 {
        return b
    }
    
    return gcd_recursive(b, a%b)
}

// 辗转相除法 迭代版本 (a>b)
// a, b := b, a%b
func gcd_iteration(a, b int) int {
    if b > a {
        a, b = b, a
    }
    for b > 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
	// f1, f2 := 1, 1
	// f1, f2 = f2, f1+f2
	var dfs func(int) int
	dfs = func(n int) int {
		if n == 1 || n == 2{
			return 1
		}
		return dfs(n-2) + dfs(n-1)
	}

	fmt.Println(dfs(100))
}