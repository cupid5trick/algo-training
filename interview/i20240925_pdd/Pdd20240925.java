package interview.i20240925_pdd;

import java.util.Arrays;
import java.util.Deque;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Set;

/**
 * 密码锁，初始[0,0,0,0]，密码是[3,4,5,6]。每一位每次可以+1,-1, 要求不能经过 [[1,0,0,0],[3,4,5,5]]，最少需要多少次操作？？？
 */
public class Pdd20240925 {

    public static class Info {
        public int st;
        public int step;

        public Info(int a, int b) {
            st = a;
            step = b;
        }
    }

    public static int state(int[] st) {
        int res = 0;
        for (int i = 0; i < st.length; i++) {
            res += st[i] << (i * 8);
        }
        return res;
    }

    public static void debugState(int st, int i) {
        int[] state = new int[] { ((st >> 0) & 0xff), ((st >> 8) & 0xff), ((st >> 16) & 0xff), ((st>>24) & 0xff) };
        System.err.println("debug: " + Arrays.toString(state) + ", " + i);
    }

    public static Integer minSteps(int[] start, int[] end, List<int[]> deadlocks) {
        int LIMIT = 10;
        Deque<Info> q = new LinkedList<>();
        Set<Integer> unreachable = new HashSet<>();
        Set<Integer> vis = new HashSet<>();
        for (int[] lk : deadlocks) {
            unreachable.add(state(lk));
        }
        int st = state(start);
        int en = state(end);
        debugState(st, -1);
        debugState(en, -1);
        q.offer(new Info(st, 0));
        vis.add(st);

        while (!q.isEmpty()) {
            Info cur = q.pollFirst();
            debugState(cur.st, cur.step);
            for (int i = 0; i < end.length; i++) {
                for (int k = -1; k < 2; k += 2) {
                    // clear bits
                    int newst = cur.st & ~(0xff << (i * 8));
                    // set bits
                    int nxt = (((cur.st >> (i * 8))&0xff) + LIMIT + k) % LIMIT;
                    newst |=  nxt << (i * 8);
                    debugState(newst, cur.step+1);
                    if (!unreachable.contains(newst) && !vis.contains(newst)) {
                        q.offer(new Info(newst, cur.step + 1));
                        vis.add(newst);
                    }
                }
            }
            if (cur.st == en) {
                return cur.step;
            }
        }

        return -1;
    }

    public static void main(String[] args) {
        int[] initial = { 0, 0, 0, 0 };
        int[] target = { 3, 4, 5, 6 };
        List<int[]> invalidStates = Arrays.asList(
                new int[] { 1, 0, 0, 0 },
                new int[] { 3, 4, 5, 5 });

        int minOperations = minSteps(initial, target, invalidStates);
        System.out.println("最少需要的操作次数: " + minOperations);
    }
}
