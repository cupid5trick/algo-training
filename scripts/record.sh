#!/bin/bash

# 用法函数
usage() {
    echo "用法: $0 <tag> <problem_name> <problem_url> <file>"
    echo "示例:"
    echo "  $0 动态规划 '最长上升子序列' https://leetcode.cn/problems/longest-increasing-subsequence/ dp/LIS.cpp"
    exit 1
}

# 参数检查
if [ $# -lt 4 ]; then
    usage
fi

tag=${1:-""}
problem_name=$2
problem_url=$3
file=$4

f=Problems.md
id=$(grep -P "^\|\d+" "$f" | wc -l)
prefix="https://github.com/cupid5trick/algo-training/tree/main/"

# 输出 markdown 表格行
echo -ne "|$((id+1)) |$tag | |[$problem_name]($problem_url) |[$file]($prefix/$file) |\n"
