package interview.intern_coding.papertests.aliguoji20240415;

import java.util.ArrayList;
import java.util.Scanner;

/**
 * 小红拿到了一棵有根树，根节点为1号节点。小红定义每个节点的权值为:该节点每个儿子的子树大小的极差。特殊的，叶子的权值为0
 * 请你求出所有节点的权值之和。
 * 输入描述
 * 第一行输入一个正整数n，代表节点的数量接下来的n-1行，每行输入两个正整数 u,v，代表节点v和节点u有一条边连接1<u,v<n, 1 ≤n ≤10^5)
 * 输出描述
 * 一个整数，代表所有节点的权值之和。
 */
class C {
    static int ans = 0;

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        ArrayList<Integer>[] adj = new ArrayList[n];
        for (int i = 0; i < n; i++) {
            adj[i] = new ArrayList<>();
        }

        for (int i = 1; i < n; i++) {
            int u = scanner.nextInt() - 1;
            int v = scanner.nextInt() - 1;
            adj[u].add(v);
        }

        postOrder(adj, 0);
        System.out.println(ans);
    }

    public static int postOrder(ArrayList<Integer>[] adj, int root) {
        if (adj[root].size() == 0) {
            return 1;
        }

        int sumOfSubtrees = 0;
        int min = 0x7fffffff;
        int max = 0;
        for (int child : adj[root]) {
            int subSize = postOrder(adj, child);
            sumOfSubtrees += subSize;
            min = Math.min(min, subSize);
            max = Math.max(max, subSize);
        }
        ans += max - min;

        return sumOfSubtrees + 1;
    }

}