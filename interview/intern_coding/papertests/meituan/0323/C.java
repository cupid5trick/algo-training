import java.util.Scanner;

/**
 * 小美拿到了一个排列，其中初始所有元素都是红色，但有此元素被类成了白色。
 * 小美每次操作可以选择交换任意两个红色元素的位置，她希望操作可能少的次数使得数组变成非降序，你能帮帮她吗?
 * 排列是指:一个长度为 n 的数组，其中1到n每个元素恰好出现了一次。
 * 输入描述：
 * 第一行输入一个正整数7代表数组的长度第二行输入几个正整数Q:，代表数组的元素,
 * 第二行输入一个长度为72的字符电，代表数组元素的染色情况，第1个字符为R代表第1个元素被染成红色，为W代表初始的白色
 * 输出描述
 * 如果无法完成排序，请输出-1否则输出一个整数，代表操作的最小次数。
 * 示例 1
 * 输入
 * 9
 * 1 3 2 4
 * WRRW
 * 输出
 * 1
 */
class C {
    public static void main(String[] args) {
        run();

    }

    public static void run() {
        try (Scanner sc = new Scanner(System.in)) {
            int n = sc.nextInt();
            int[] nums = new int[n];
            for (int i = 0; i < n; i++) {
                nums[i] = sc.nextInt();
            }
            sc.nextLine();
            String color = sc.nextLine();
            for (int i = 0; i < n; i++) {
                char c = color.charAt(i);
                if (c == 'W' && nums[i] != i + 1) {
                    System.out.println(-1);
                }
            }

            int swaps = n;
            for (int i = 0; i < n; i++) {
                if (nums[i] == i + 1 || (nums[i] > i + 1 && nums[nums[i] - 1] == i + 1)) {
                    swaps--;
                }
                if (nums[i] > i + 1) {
                    nums[nums[i] - 1] = i + 1;
                }
            }
            System.out.println(swaps);
        }
    }
}