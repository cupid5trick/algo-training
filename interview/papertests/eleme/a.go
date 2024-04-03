package main

import "fmt"

// 题目：我有个字符串，abc, 我要获取它下面所有字符组成的子窜，"" a b c ab ac bc abc


func main() {
	str := "abc"
	n := len(str)
	
	fmt.Println("")
	for l:=1; l <= n ; l ++ {
	  for i:=0; i+l <= n; i ++ {
		t := str[i:i+l]
		fmt.Println(t)
	  }
	}
  }