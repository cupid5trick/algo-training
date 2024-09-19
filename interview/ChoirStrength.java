package interview;

import java.util.Scanner;

public class ChoirStrength {

    // 判断当前的 mid 是否可以通过 k 次训练，将所有人的最小值提升到 mid
    public static boolean canAchieve(int[] abilities, int n, int k, int l, int mid) {
        int[] dp = new int[n]; // 标记数组，表示在当前位置是否需要进行调整
        int count = 0;         // 记录训练的次数
        int windowSum = 0;     // 滑动窗口中的调整人数

        for (int i = 0; i < n; i++) {
            // 如果当前能力值小于 mid，表示需要调整
            if (abilities[i] < mid) {
                dp[i] = 1; // 标记需要调整
            }
        }

        // 滑动窗口遍历
        for (int i = 0; i < n; i++) {
            if (i >= l) {
                windowSum -= dp[i - l]; // 窗口移除
            }
            windowSum += dp[i]; // 窗口增加

            // 如果当前窗口中的需要调整的人数大于0，且剩余训练次数足够
            if (windowSum > 0) {
                count++;
                windowSum = 0; // 使用一次训练机会重置窗口
                if (count > k) {
                    return false; // 训练次数不足
                }
            }
        }

        return true;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        // 输入
        int n = scanner.nextInt(); // 人数
        int k = scanner.nextInt(); // 训练次数
        int l = scanner.nextInt(); // 每次最多调整连续人数
        int[] abilities = new int[n]; // 能力值数组
        for (int i = 0; i < n; i++) {
            abilities[i] = scanner.nextInt();
        }

        // 二分查找
        int low = 1; // 能力值最小为1
        int high = (int) 1e9; // 能力值最大为10^9
        int result = 0;

        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (canAchieve(abilities, n, k, l, mid)) {
                result = mid; // 更新结果
                low = mid + 1; // 尝试更大的 mid
            } else {
                high = mid - 1; // mid 太大，降低范围
            }
        }

        // 输出结果
        System.out.println(result);
    }
}
