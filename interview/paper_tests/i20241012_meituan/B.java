package interview.paper_tests.i20241012_meituan;

import java.math.BigInteger;
import java.util.*;
import java.util.Scanner;
/*
从给定的混合字符串中提取出所有的非负整数，并将这些整数按照降序排列，然后输出第k个数。如果不存在第k个数，则输出"N"。以下是题目的具体要求：

输入描述：

第一行输入一个正整数k（1≤k≤10^5），表示需要输出的数的位置。
第二行输入一个长度不超过10^5，且由小写字母和数字混合构成的字符串s。
输出描述：

输出第k个数，不包含前导0；如果不存在第k个数，则输出"N"。

# 总结

在遍历的同时找到分割位置，不要忘记最后一段的分割、以及BigInteger高精度计算的掌握。
 */

public class B {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int k = sc.nextInt();
        sc.nextLine();
        String s = sc.nextLine();
        StringBuilder b = new StringBuilder();
        List<BigInteger> q = new ArrayList<>();
        int n = s.length();
        for (int i = 0; i <= s.length(); i ++) {
            if (i < n && s.charAt(i) >= '0' && s.charAt(i) <= '9') {
                b.append(s.charAt(i));
            } else if (b.length() > 0) {
                q.add(new BigInteger(b.toString()));
                b = new StringBuilder();
            }
            
        }
        Collections.sort(q, Collections.reverseOrder());
        System.out.println(Arrays.toString(q.stream().toArray()));
        if (k > q.size()) {
            System.out.println("N");
        }
        else {
            System.out.println(q.get(k-1).toString());
        }
    }
}
