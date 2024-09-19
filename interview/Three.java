package interview;

import java.util.Scanner;

public class Three {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int k = sc.nextInt();
        int l = sc.nextInt();
        int[] nums = new int[n];
        for (int i =0; i < n; i++) {
            nums[i] = sc.nextInt();
        }
        int ans = 0;
        // 二分值域
        int left = 0, right = 0x7fff_ffff;
        while (left+1 < right) {
            int mid = (left+right)/2;
            
            // k 次训练够不够？
            int pre = -1;
            int cnt = 0;
            for (int i = 0; i < n; i ++) {
                if (nums[i] >= mid) {
                    continue;
                }
                if (pre < 0 || pre >=0 && i - pre + 1 > l) {
                    pre = i;
                    cnt ++;
                }
            }

            if (!(cnt <= k)) {
                right = mid;
            } else {
                left = mid;
                // 合法的能力值
                ans = mid;
            }
        }

        System.out.println(ans);
    }
}
