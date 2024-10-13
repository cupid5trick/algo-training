package interview.intern_coding.wy20240509;

import java.util.Scanner;

public class A {
    /**
     * 3 
100 0
20 18
5 3
     * @param args
     */
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int t = sc.nextInt();
        for (int i = 0; i < t; i++) {
            int c = sc.nextInt();
            int m = sc.nextInt();

            int a = m/5;  // 3
            int n = Math.min(a, c/5);
            int tmp = n;
            c = c - n*5 + m%5;
            n += c/10;
            int total = tmp*5 + Math.max((n-tmp)*10 - m % 5, 0);
            System.out.printf("%d %d\n", n, total);

        }
    }
}
