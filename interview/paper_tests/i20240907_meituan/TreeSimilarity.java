package interview.paper_tests.i20240907_meituan;

import java.util.*;

public class TreeSimilarity {
    private static List<List<Integer>> tree;      // 邻接表表示的树
    private static Map<Integer, Integer> childCountMap; // 用于记录直接子节点数量出现次数的哈希表
    private static int totalSimilarPairs; // 总的相似节点对数

    // DFS 函数
    private static void dfs(int node, int parent) {
        int childCount = 0;

        for (int neighbor : tree.get(node)) {
            if (neighbor == parent) {
                continue; // 跳过父节点
            }
            childCount++;
            dfs(neighbor, node);
        }

        // 记录每个节点的直接子节点数量的出现次数
        childCountMap.put(childCount, childCountMap.getOrDefault(childCount, 0) + 1);
    }

    // 计算相似节点对数
    private static int calculateSimilarNodePairs(int root) {
        dfs(root, -1);

        // 根据哈希表计算相似节点对数
        for (int count : childCountMap.values()) {
            if (count > 1) {
                totalSimilarPairs += count * (count - 1) / 2; // C(k, 2)
            }
        }

        return totalSimilarPairs;
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        // 读取测试用例数量
        int t = sc.nextInt();

        while (t-- > 0) {
            // 读取节点数
            int n = sc.nextInt();

            // 初始化树和相关变量
            tree = new ArrayList<>(n);
            childCountMap = new HashMap<>();
            totalSimilarPairs = 0;

            for (int i = 0; i < n; i++) {
                tree.add(new ArrayList<>());
            }

            // 读取每条边并构建邻接表
            for (int i = 0; i < n - 1; i++) {
                int u = sc.nextInt() - 1; // 输入的节点编号从1开始，所以需要减1
                int v = sc.nextInt() - 1;
                tree.get(u).add(v);
                tree.get(v).add(u); // 双向图
            }

            // 计算相似节点对数
            calculateSimilarNodePairs(0);

            // 输出哈希表内容和相似节点对数
            System.out.println(childCountMap);
            System.out.println(totalSimilarPairs);
        }

        sc.close();
    }
}
