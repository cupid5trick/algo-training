package interview.papertests.aliguoji20240415;

import java.util.Scanner;

/**
 * 小红拿到了一个数组，她希望选择两个不相邻的数，使得它们的和为偶数，小红想知道有多少种不同的取数方案?
 * 输入描述
 * 第一行输入一个正整数n，代表数组的大小。
 * 第二行输入几个正整数ai，代表数组的元素1 <n,ai < 200000
 * 输出描述
 * 一个整数，代表取数方案数
 */
public class A {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int[] a = new int[n];
        for (int i = 0; i < n; i++) {
            a[i] = in.nextInt() % 2;
        }
        int oneCount = 0;
        int zeroCount = 0;
        int l = -1;
        int h = 1;
        long res = 0;
        while (h < n) {
            if (l >= 0) {
                if (a[l] == 1)
                    oneCount++;
                else
                    zeroCount++;
            }
            if (a[h] == 1)
                res += oneCount;
            else
                res += zeroCount;
            h++;
            l++;
        }
        System.out.println(res);
    }
}
