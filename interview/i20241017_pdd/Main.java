package interview.i20241017_pdd;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Stream;

public class Main {
    public static void main(String[] args) {
        int[] nums = new int[] {5,4,3,2,1};
        int n = nums.length;
        int[][] p = new int[n][2];
        for (int i = 0; i < n; i ++) {
            p[i][0] = nums[i];
            p[i][1] = i;
        }
        Arrays.sort(p, (i, j) -> i[0] - j[0]);
        p = Arrays.copyOf(p, 2*n);
        System.out.println(Arrays.toString(Stream.of(p).map(x -> Arrays.toString(x)).toArray()));

        int sz = n;
        int total = 0;
        List<int[]> ans = new ArrayList();
        for (int i = 1; i < 2*n; i += 2) {
            if (p[i] == null) {
                break;
            }
            int x =  p[i][0] + p[i-1][0];
            p[sz] = new int[] {x, sz};
            for (int j = sz-1; j > i; j --) {
                int[] tmp = p[j];
                if (p[j+1][0] < p[j][0]) {
                    p[j] = p[j+1];
                    p[j+1] = tmp;
                }
            }
            System.out.println(Arrays.toString(Stream.of(p).map(xx -> Arrays.toString(xx)).toArray()));

            /// 记录答案
            total += x;
            ans.add(new int[] {p[i][0], p[i][1], p[i-1][0], p[i-1][1], x, sz});
            sz ++;
        }
        for (int[] res: ans) {
            System.out.println(Arrays.toString(res));
        }
        System.out.println(total);
    }
}