
import java.io.BufferedReader;
import java.io.IOException;
import java.io.StringReader;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.StringTokenizer;

/**
 * 小红有一个 n 个节点的树，一开始每条边都没有颜色。
 * 小红每次可以选择两条有公共点且都没有被染色的边，然后将这两条边染成红色。小红想知道，最多可以染色多少次，并且输出一种方案。
 * 输入描述
 * 第一行一个整数 n，表示树的节点数，保证 n是奇数。
 * 接下来 n-1行，每行两个整数 u, v，表示 u和v之间有一条边。
 * 1 ≤n≤ 10^5
 * 1<u,v<n
 */
public class C {
    static Map<Integer, List<Integer>> graph = new HashMap<>();
    static int[] uncoloredEdges;
    static List<int[]> colors = new ArrayList<>();

    public static void main(String[] args) throws IOException {
        String input = "7\n" +
                "1 2\n" +
                "2 3\n" +
                "2 4\n" +
                "1 5\n" +
                "5 6\n" +
                "5 7";
        run(input);
    }

    static void run(String input) throws IOException {
        BufferedReader reader = new BufferedReader(new StringReader(input));
        // Uncomment below line for user input
        // BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(reader.readLine());
        uncoloredEdges = new int[n];

        for (int i = 0; i < n; i++) {
            graph.put(i, new ArrayList<>());
        }

        for (int i = 0; i < n - 1; i++) {
            StringTokenizer st = new StringTokenizer(reader.readLine());
            int u = Integer.parseInt(st.nextToken()) - 1;
            int v = Integer.parseInt(st.nextToken()) - 1;
            graph.get(u).add(v);
            graph.get(v).add(u);
            uncoloredEdges[u]++;
            uncoloredEdges[v]++;
        }

        dfs(0, -1);

        int maxColorings = 0;
        for (int edge : uncoloredEdges) {
            maxColorings += edge / 2;
        }

        System.out.println(maxColorings);
        for (int[] color : colors) {
            System.out.println((color[0] + 1) + " " + (color[1] + 1));
        }
    }

    static int dfs(int node, int parent) {
        int childCount = 0;
        for (int child : graph.get(node)) {
            if (child != parent) {
                childCount += dfs(child, node);
            }
        }

        if (parent != -1) {
            uncoloredEdges[node] -= Math.min(2, childCount);
            if (uncoloredEdges[node] == 0) {
                colors.add(new int[] { parent, node });
                uncoloredEdges[parent]--;
                if (uncoloredEdges[parent] == 0 && parent != 0) {
                    colors.add(new int[] { 0, parent });
                }
                childCount++;
            }
        }

        return childCount;
    }
}
