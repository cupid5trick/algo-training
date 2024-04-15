package interview.papertests.aliguoji20240415;

import java.util.Arrays;
import java.util.Scanner;

/**
 * 应该是个典型的动态规划，有点像 213. 打家劫舍 II - 力扣（LeetCode）: https://leetcode.cn/problems/house-robber-ii/description/
 * 
 * 小苯面前有n盏灯首尾相接排成一圈，每盏灯都有"蓝"和"黄"两种可选的颜色，分别用(xi, yi)对描述每盏灯。
 * 如果第i盏灯打开了"蓝色"，则可以获得 xi分，如果打开"黄色"，则可以获得 yi 分。
 * 如果不点亮该灯，则不得分。现在大白熊能想让小茶选择一些灯点亮，但需要满足，如果有两盏相邻的灯同时被点完，则其两者的颜色必须不同，
 * 同时大白能还限制小苯必须点亮尽可能多的灯。
 * 现在小苯想知道他在点灯数量尽可能多的情况下最多可以获得多少分，请你帮帮他
 * 
 * 输入包含n+1行。
 * 第一行一个正整数 n 表示灯的个数。(2 <n <10^5)接下来 n 行，每行两个正整数 xi, yi (1 <= xi, yi <=
 * 10^9)，描述每盖灯。
 * 输出描述
 * 输出一行一个整数表示小苯的最大得分
 */
public class B {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int n = sc.nextInt();

        long[][] lamb = new long[n][2];
        for (int i = 0; i < n; i++) {
            lamb[i][0] = sc.nextLong();
            lamb[i][1] = sc.nextLong();
        }

        long ans = 0;
        ans = Math.max(ans, solver(lamb, 0, 0));
        ans = Math.max(ans, solver(lamb, 0, 1));
        ans = Math.max(ans, solver(lamb, n - 1, 0));
        ans = Math.max(ans, solver(lamb, n - 1, 1));
        System.out.println(ans);

    }

    public static long solver(long[][] lamb, int ex, int st) {
        int n = lamb.length;
        long tmp = lamb[ex][st];
        lamb[ex][st] = 0;

        long[][] dp = new long[n][2];
        dp[0][0] = lamb[0][0];
        dp[0][1] = lamb[0][1];

        for (int i = 0; i < n; i++) {
            dp[i][0] += i >= 1 && lamb[i - 1][1] >= 0 ? lamb[i - 1][0] : 0;
            dp[i][1] += i >= 1 && lamb[i - 1][1] >= 0 ? lamb[i - 1][1] : 0;
        }
        lamb[ex][st] = tmp;

        return Math.max(dp[n - 1][0], dp[n - 1][1]);
    }

}
