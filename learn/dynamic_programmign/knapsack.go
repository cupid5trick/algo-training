package main

/*
【题单】动态规划（入门/背包/状态机/划分/区间/状压/数位/树形/数据结构优化） - 力扣（LeetCode）: https://leetcode.cn/circle/discuss/tXLS3i/
讲解：[0-1 背包 完全背包](https://leetcode.cn/link/?target=https://www.bilibili.com/video/BV16Y411v7Y6/)

#### [](https://leetcode.cn/circle/discuss/tXLS3i//#§31-01-背包)§3.1 0-1 背包

每个物品只能选一次。

-   [2915. 和为目标值的最长子序列的长度](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/) 1659
-   [顺丰 02. 小哥派件装载问题](https://leetcode.cn/contest/sf-tech/problems/cINqyA/)
-   [416. 分割等和子集](https://leetcode.cn/problems/partition-equal-subset-sum/)
-   [494. 目标和](https://leetcode.cn/problems/target-sum/)
-   [2787. 将一个数字表示成幂的和的方案数](https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/) 1818
-   [474. 一和零](https://leetcode.cn/problems/ones-and-zeroes/)（二维）
-   [1049. 最后一块石头的重量 II](https://leetcode.cn/problems/last-stone-weight-ii/) 2092
-   [1774. 最接近目标价格的甜点成本](https://leetcode.cn/problems/closest-dessert-cost/)
-   [879. 盈利计划](https://leetcode.cn/problems/profitable-schemes/) 2204
-   [3082. 求出所有子序列的能量和](https://leetcode.cn/problems/find-the-sum-of-the-power-of-all-subsequences/) ~2300
-   [956. 最高的广告牌](https://leetcode.cn/problems/tallest-billboard/) 2381
-   [2518. 好分区的数目](https://leetcode.cn/problems/number-of-great-partitions/) 2415
-   [2742. 给墙壁刷油漆](https://leetcode.cn/problems/painting-the-walls/) 2425
-   [LCP 47. 入场安检](https://leetcode.cn/problems/oPs9Bm/)
-   [2291. 最大股票收益](https://leetcode.cn/problems/maximum-profit-from-trading-stocks/)（会员题）
-   [2431. 最大限度地提高购买水果的口味](https://leetcode.cn/problems/maximize-total-tastiness-of-purchased-fruits/)（会员题）

#### [](https://leetcode.cn/circle/discuss/tXLS3i//#§32-完全背包)§3.2 完全背包

物品可以重复选，无个数限制。

-   [322. 零钱兑换](https://leetcode.cn/problems/coin-change/)
-   [518. 零钱兑换 II](https://leetcode.cn/problems/coin-change-ii/)
-   [279. 完全平方数](https://leetcode.cn/problems/perfect-squares/)
-   [1449. 数位成本和为目标值的最大数字](https://leetcode.cn/problems/form-largest-integer-with-digits-that-add-up-to-target/) 1927 打印方案

#### [](https://leetcode.cn/circle/discuss/tXLS3i//#§33-多重背包)§3.3 多重背包

物品可以重复选，有个数限制。

> 注：力扣上只有求方案数的题目。

-   [1155. 掷骰子等于目标和的方法数](https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/) 1654
-   [2585. 获得分数的方法数](https://leetcode.cn/problems/number-of-ways-to-earn-points/) 1910
-   [2902. 和带限制的子多重集合的数目](https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/) 2759

#### [](https://leetcode.cn/circle/discuss/tXLS3i//#§34-分组背包)§3.4 分组背包

同一组内的物品至多/恰好选一个。

-   [1981. 最小化目标值与所选元素的差](https://leetcode.cn/problems/minimize-the-difference-between-target-and-chosen-elements/) 2010
-   [2218. 从栈中取出 K 个硬币的最大面值和](https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles/) 2158
*/

// leetcode 322. 零钱兑换 - 力扣（LeetCode）: https://leetcode.cn/problems/coin-change
// dfs(i, j) = min(dfs(i-1, j), dfs(i-1, j-coins[i])+1), j >= coins[i]
// dfs(0, 0) = 0, 其他 amount+1

func coinChange_322(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for _, coin := range coins {
		for j := 1; j <= amount; j++ {
			if j >= coin {
				dp[j] = min(dp[j], dp[j-coin]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 518. 零钱兑换 II - 力扣（LeetCode）: https://leetcode.cn/problems/coin-change-ii/
// 原本的零钱兑换：每种面值数量无限，求组成 amount 金额需要的最少硬币数量
// dfs(i,j) = min(dfs(i-1,j), dfs(i-1, j-coins[i])+1), j >= coins[i] (dfs(0,0)=0)
// 零钱兑换II：求可以凑出总金额的硬币组合方案数
// dfs(i,j) = sum(dfs(i-1,j), dfs(i-1,j-k*coins[i])), j >= k*coins[i]
// 初始值 dfs(i,j) = 0
// 优化：
// - 和零钱兑换I 一样动态规划数组可以删去 硬币面值维度
// - 由于 j-k*coins[i] 一样是从小到大计算的，可以像前缀和的思想去优化成一层循环
func change_518_original(amount int, coins []int) int {
    // dp := make([]int, amount+1)
    // dp[0] = 1
    // for _, coin := range coins {
    //     for j := coin; j <= amount; j ++ {
    //         dp[j] += dp[j-coin]
    //     }
    // }

    // return dp[amount]

    dp := make([][]int, len(coins)+1)
    for i := 0; i <= len(coins); i ++ {
        dp[i] = make([]int, amount+1)
    }
    dp[0][0] = 1
    
    for i:=1; i <= len(coins); i ++ {
        coin := coins[i-1]
        for j := 0; j <= amount; j ++ {
            dp[i][j] = dp[i-1][j]
            for k := 1; k*coin <= j ; k ++ {
                dp[i][j] += dp[i-1][j-k*coin]
            }
        }
    }

    return dp[len(coins)][amount]
}

// 空间优化和前缀和优化版本
func change_518_optim(amount int, coins []int) int {
    dp := make([]int, amount+1)
    dp[0] = 1
    for _, coin := range coins {
        for j := coin; j <= amount; j ++ {
            dp[j] += dp[j-coin]
        }
    }

    return dp[amount]
}