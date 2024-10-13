package main

/*
- 3320. 统计能获胜的出招序列数: <https://leetcode.cn/problems/count-the-number-of-winning-sequences/description/>
# 解题思路

每一局的分数情况有三种，B赢：1、A赢：-1，平局：0。
1. 利用三维DP f(i,x,j) 表示第i轮出了j，并且B-A的得分之差为 x。
2. 由于 x 的范围是 [-n,n]，需要通过 f[i][x+n][j] 来访问 f(i,x,j)。
3. 分数差的计算：
	为了优化分数差计算效率，我们把出招序列 FWE 映射到 012
	对于出招方案 (A,B) B赢的情况有 (1,0), (2,1), (0, 2)，
	因此可以通过 (b+1+3)%3 == a 快速判断B赢的情况，而平局就是 a == b。
4. 递推关系式：f(i,x,j) = (f(i,x,j)+f(i-1,x-d,k)) % M
	d = 分数差
	j = 当前第i轮B的出招
	k = 上一轮B的出招
5. 可能的剪枝：
	1. 必胜：
	2. 必败
# 比赛时的 Python 代码

```python
class Solution:
    def countWinningSequences(self, s: str) -> int:
        role = "FWE"
        mp = dict(zip(role, range(3)))
        s = list(map(lambda x: mp[x], s))
    
        M = int(1e9+7)
        n = len(s)
        losenum = 2*n

        def score_diff(i: int, j: int) -> int:
            # 赢 +1分
            if (i+1)%3 == j:
                return 1
            elif (j+1)%3 == i:    # 输 -1分
                return -1
            return 0
            
        f = [[[0]*3 for _ in range(losenum+1)] for _ in range(n)]
        # 初值
        for j in range(3):
            f[0][n+score_diff(s[0], j)][j] = 1
        # dp
        for i in range(1,n):
            for x in range(-i-1,i+2):
                for j in range(3):
                    for k in range(3):
                        if j == k: continue
                        d = score_diff(s[i], j)
                        if 0 <= x+n-d <= losenum:
                            # if any(x < 0 for x in [i, x+n, j, i-1, k, x+n-d]):
                            #     print(f"0000{[i, x+n, j, i-1, k, x+n-d]}")
                            f[i][x+n][j] += f[i-1][x+n-d][k]
                            f[i][x+n][j] %= M
        res = 0
        for k in range(n+1, losenum+1):
            for x in f[n-1][k]:
                res = (res+x)%M
                
        return res
```

# 总结
本题是多维的状态机 DP，练习题目参考动态规划题单的【状态机DP】和【多维DP】。
- 分享丨【题单】动态规划（入门/背包/状态机/划分/区间/状压/数位/树形/数据结构优化）: <https://leetcode.cn/circle/discuss/tXLS3i/>
*/

//lint:file-ignore U1000 unused function

func countWinningSequences(s string) int {
	// 映射字符到数字
	mp := map[byte]int{
		'F': 0,
		'W': 1,
		'E': 2,
	}

	// 将字符串转换为对应的数字索引
	sIndices := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		sIndices[i] = mp[s[i]]
	}

	M := int(1e9 + 7)
	n := len(s)
	losenum := 2 * n

	// 定义得分差异函数
	scoreDiff := func(i, j int) int {
		if (i+1)%3 == j {
			return 1 // 赢 +1 分
		} else if (j+1)%3 == i {
			return -1 // 输 -1 分
		}
		return 0 // 平局 0 分
	}

	// 三维 dp 数组 f[i][x+n][j] 表示前 i 个字符的状态
	f := make([][][]int, n)
	for i := range f {
		f[i] = make([][]int, losenum+1)
		for j := range f[i] {
			f[i][j] = make([]int, 3)
		}
	}

	// 初始状态
	for j := 0; j < 3; j++ {
		f[0][n+scoreDiff(sIndices[0], j)][j] = 1
	}

	// 动态规划求解
	for i := 1; i < n; i++ {
		for x := -i - 1; x <= i+1; x++ {
			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					if j == k {
						continue
					}
					d := scoreDiff(sIndices[i], j)
					if x+n-d >= 0 && x+n-d <= losenum {
						f[i][x+n][j] = (f[i][x+n][j] + f[i-1][x+n-d][k]) % M
					}
				}
			}
		}
	}

	// 计算结果
	res := 0
	for k := n + 1; k <= 2*n; k++ {
		for x := 0; x < 3; x++ {
			res = (res + f[n-1][k][x]) % M
		}
	}

	return res
}