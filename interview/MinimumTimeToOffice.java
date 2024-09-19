package interview;

import java.util.Arrays;
import java.util.Scanner;

public class MinimumTimeToOffice {

    // 检查在时间T内是否能够让所有员工都拿到通行证并到达办公室
    public static boolean canAssign(int[] employees, int[] passes, int n, int k, int p, long T) {
        int passIndex = 0; // 通行证的索引

        for (int i = 0; i < n; i++) {
            // 找到一个能满足当前员工i在时间T内拿到通行证并到达办公室的通行证
            while (passIndex < k
                    && (Math.abs(employees[i] - passes[passIndex]) + Math.abs(passes[passIndex] - p)) > T) {
                passIndex++;
            }

            // 如果已经没有通行证可以分配了，说明无法在时间T内完成
            if (passIndex == k) {
                return false;
            }

            // 当前通行证分配给第i个员工
            passIndex++;
        }

        return true;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        // 输入数据
        int n = scanner.nextInt(); // 员工人数
        int k = scanner.nextInt(); // 通行证数量
        int p = scanner.nextInt(); // 办公室位置
        int[] employees = new int[n];
        int[] passes = new int[k];

        for (int i = 0; i < n; i++) {
            employees[i] = scanner.nextInt();
        }

        for (int i = 0; i < k; i++) {
            passes[i] = scanner.nextInt();
        }

        // 对员工和通行证位置进行排序
        Arrays.sort(employees);
        Arrays.sort(passes);

        // 二分查找最小的最大时间
        long low = 0;
        long high = 2L * (long) 1e9;
        long result = high;

        while (low <= high) {
            long mid = low + (high - low) / 2;

            if (canAssign(employees, passes, n, k, p, mid)) {
                result = mid; // 尝试更小的时间
                high = mid - 1;
            } else {
                low = mid + 1; // 增大时间
            }
        }

        // 输出结果
        System.out.println(result);
    }
}
