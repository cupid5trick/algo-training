package interview.i20240925_pdd;


import java.util.*;
/**
 * ChatGPT 解法
 */
public class Pdd {
    
    public static class Info {
        public int st;
        public int step;

        public Info(int a, int b) {
            st = a;
            step = b;
        }
    }

    public static int state(int[] st) {
        int res  = 0;
        for (int i = 0; i < st.length; i++) {
            res += st[i] << (i * 8);
        }
        return res;
    }

    public static Integer minSteps(int[] start, int[] end, List<int[]> deadlocks) {
        int LIMIT = 10;
        Deque<Info> q = new LinkedList<>();
        Set<Integer> unreachable = new HashSet<>();
        Set<Integer> visited = new HashSet<>();

        for (int[] lk : deadlocks) {
            unreachable.add(state(lk));
        }
        
        int st = state(start);
        int en = state(end);
        q.offer(new Info(st, 0));
        visited.add(st);

        while (!q.isEmpty()) {
            Info cur = q.poll();

            if (cur.st == en) {
                return cur.step;
            }

            for (int i = 0; i < end.length; i++) {
                for (int k = -1; k <= 1; k += 2) {
                    int newVal = (cur.st >> (i * 8) & 0xff) + k;
                    newVal = (newVal + LIMIT) % LIMIT; // Wrap around
                    int nextSt = cur.st & ~(0xff << (i * 8)); // Clear the current position
                    nextSt |= (newVal << (i * 8)); // Set new value

                    if (!unreachable.contains(nextSt) && visited.add(nextSt)) {
                        q.offer(new Info(nextSt, cur.step + 1));
                    }
                }
            }
        }

        return -1; // 如果无法到达目标状态
    }

    public static void main(String[] args) {
        int[] initial = {0, 0, 0, 0};
        int[] target = {3, 4, 5, 6};
        List<int[]> invalidStates = Arrays.asList(
                new int[]{1, 0, 0, 0},
                new int[]{3, 4, 5, 5}
        );

        int minOperations = minSteps(initial, target, invalidStates);
        System.out.println("最少需要的操作次数: " + minOperations);
    }
}
