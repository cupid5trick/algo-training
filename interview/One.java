package interview;

import java.util.Scanner;

public class One {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int t = sc.nextInt();
        while (t -- > 0) {
            int n = sc.nextInt();
            int m = sc.nextInt();
            long k = sc.nextInt();
            long ans = 0;
            if (k > 0) {
                int l = Math.min((int)k, n-1);
                ans += m*1L*l*(l+1)/2;
                k -= n-1;
                // System.out.println("debug:" + ans + "," + k);
            }
            if (k > 0) {
                int l = Math.min((int)k, m-1);
                ans += m*1L*(n-1)*l + l*(l+1)/2;
                k -= m-1;
                // System.out.println("debug:" + ans + "," + k);
            }
            if (k > 0) {
                ans += (2L*n*m-3)*(k/2) + (1L*n*m-2)*(k%2);
            }
            // System.out.println("debug:" + ans + "," + k);
            System.out.println(ans);
        }
    }
}
