package interview.papertests.xiaohongshu.t20230329;

import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.util.Arrays;
import java.util.Scanner;

/**
 * 小苯是小红书APP的忠实用户，他有n个账号，每个账号粉丝数为a。
 * 这天他有创建了一个新账号，他希望新账号的粉丝数恰好等于x。为此他可以向自己的已有账号的粉丝推荐新账号，这样新账号就得到了之前粉丝的关注。
 * 他想知道他最少需要在多少个旧账号发“推荐新账号”的文章，可以使得他新账号的粉丝数恰好为x。除此之外，还可以最多从中选择一个旧账号多次发送“推荐新账号”的文章。
 * （我们假设所有旧账号的粉丝们没有重叠，并且如果在第i个旧账号的粉丝中间推荐了新账号，则新账号会直接涨粉
 * ai/2 下取整个。而如果选择在第i个旧账号中多次推荐新账号，那新账号就可以直接涨粉ai）
 */

public class B {
    static class Main {
        public static void main(String[] args) {
            String s = "100 100 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10 1 2 3 4 10";
            InputStream input = new ByteArrayInputStream(s.getBytes());
            Scanner sc = new Scanner(input);
            int n = sc.nextInt();
            int x = sc.nextInt();
            int[] arr = new int[n];
            for (int i = 0; i < n; i++) {
                arr[i] = sc.nextInt();
            }

            long startTime = System.nanoTime();
            int ans = solve(arr, -1, x);

            System.out.printf("%d, %d, %d, %s\n", ans, n, x, Arrays.toString(arr));
            for (int i = 0; i < arr.length; i++) {
                ans = Math.min(ans, solve(arr, i, x - arr[i]) + 1);
                System.out.println(ans);
            }
            if (ans > n) {
                ans = -1;
            }
            System.out.println(ans);
            long endTime = System.nanoTime();
            long duration = (endTime - startTime) / 1000000; // 转换为毫秒
            System.out.printf("Execution time: %d ms\n", duration);
        }

        public static int solve(int[] arr, int idx, int target) {
            if (target < 0) {
                return arr.length + 1;
            }
            int[] dp = new int[target + 1];
            Arrays.fill(dp, arr.length + 1);
            dp[0] = 0;
            System.out.println(Arrays.toString(dp));

            for (int ai : arr) {
                if (idx != -1 && ai == idx) {
                    continue;
                } else {
                    ai /= 2;
                }
                for (int x = 1; x <= target; x++) {
                    if (x >= ai) {
                        dp[x] = Math.min(dp[x - ai] + 1, dp[x]);
                    }
                }
            }
            return dp[target];
        }
    }
}
