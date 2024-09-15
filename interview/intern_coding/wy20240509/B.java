package interview.wy20240509;

import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class B {

    static Map<String, Integer> memo;

    /**
     * 
     * @param args
     */
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        
        // 循环处理多组输入
        while (sc.hasNext()) {
            int n = sc.nextInt(); // 二进制字符串的长度
            int k = sc.nextInt(); // 最多可以交换的次数
            sc.next();
            String binaryString = sc.nextLine(); // 二进制字符串
            
            // 初始化记忆数组
            memo = new HashMap<>();
            
            // 调用函数计算结果并输出
            int result = calculateMinimumDecimalValue(n, k, binaryString);
            System.out.println(result);
        }
        
        sc.close();
    }
    
    // 计算字符串的十进制值的最小可能值
    private static int calculateMinimumDecimalValue(int n, int k, String binaryString) {
        return dfs(n, k, binaryString, 0);
    }
    
    private static int dfs(int n, int k, String binaryString, int swaps) {
        // 如果已经搜索过该状态，则直接返回之前计算的结果
        if (memo.containsKey(binaryString)) {
            return memo.get(binaryString);
        }
        
        // 计算当前状态下的十进制值
        int result = 0;
        for (int i = 0; i < n - 1; i++) {
            if (binaryString.charAt(i) == '1') {
                result += Math.pow(2, n - i - 2);
            }
        }
        
        // 如果剩余交换次数为 0 或者已经达到目标值，则返回当前结果
        if (swaps == k || result == 0) {
            memo.put(binaryString, result);
            return result;
        }
        
        // 尝试交换相邻的两个字符，更新最小值
        for (int i = 0; i < n - 1; i++) {
            if (binaryString.charAt(i) != binaryString.charAt(i + 1)) {
                StringBuilder sb = new StringBuilder(binaryString);
                char temp = sb.charAt(i);
                sb.setCharAt(i, sb.charAt(i + 1));
                sb.setCharAt(i + 1, temp);
                result = Math.min(result, dfs(n, k, sb.toString(), swaps + 1));
            }
        }
        
        // 将当前状态及其对应的结果保存到记忆数组中
        memo.put(binaryString, result);
        
        return result;
    }
}
