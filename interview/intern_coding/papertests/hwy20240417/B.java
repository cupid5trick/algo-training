package interview.papertests.hwy20240417;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.io.PrintWriter;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Scanner;

public class B {
    static int ans = 0;
    static int m;
    static Map<String, List<String>> children = new HashMap<>();
    static Map<String, Integer> stat1 = new HashMap<>();
    static Map<String, Integer> stat2 = new HashMap<>();
    static Map<String, Boolean> cloud = new HashMap<>();

    public static void main(String[] args) {
        String input = "40 12\n" +
                "a * 0 2\n" +
                "a * 1 2\n" +
                "b a 0 3\n" +
                "b a 1 5\n" +
                "c a 1 3\n" +
                "d a 0 1\n" +
                "d a 1 3\n" +
                "e b 0 2\n" +
                "f * 0 8\n" +
                "f * 1 10\n" +
                "g f 1 2\n" +
                "h * 0 4";

        run(System.in, System.out, input);
    }

    public static void run(InputStream inputStream, OutputStream outputStream, String input) {
        Scanner scanner = new Scanner(input);
        PrintWriter out = new PrintWriter(outputStream);
        m = scanner.nextInt();
        int n = scanner.nextInt();
        scanner.nextLine(); // Consume newline


        for (int i = 0; i < n; i++) {
            String line = scanner.nextLine();
            String[] parts = line.split("\\s+");
            String a = parts[0];
            String b = parts[1];
            int c = Integer.parseInt(parts[2]);
            int d = Integer.parseInt(parts[3]);

            if (!b.equals("*")) {
                if (!children.containsKey(b)) {
                    children.put(b, new ArrayList<>());
                }
                children.get(b).add(a);
            } else {
                cloud.put(a, true);
            }

            if (c == 0) {
                stat1.put(a, d);
            } else {
                stat2.put(a, d);
            }
        }

        for (String k : cloud.keySet()) {
            if (cloud.get(k)) {
                dfs(k);
            }
        }

        out.println(ans);
        out.flush();
    }

    private static int[] dfs(String i) {
        int[] res = new int[2];

        for (String j : children.getOrDefault(i, new ArrayList<>())) {
            int[] sub = dfs(j);
            res[0] += sub[0];
            res[1] += sub[1];
        }

        res[0] += stat1.getOrDefault(i, 0);
        res[1] += stat2.getOrDefault(i, 0);
        if (5*res[0]+2*res[1] > m && cloud.getOrDefault(i, false)) {
            ans ++;
        }
        return res;
    }
}
