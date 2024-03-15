package leetcode.io;

import java.io.*;
import java.util.Scanner;

// 注意类名必须为 Main, 不要有任何 package xxx 信息
public class Main {
    public static void main(String[] args) throws IOException {
        // java
        StreamTokenizer token = new StreamTokenizer(new BufferedReader(new InputStreamReader(System.in)));

        // 注意 hasNext 和 hasNextLine 的区别
        Scanner in = new Scanner(System.in);
        token.nextToken();
        int n = (int) token.nval;
        token.nextToken();
        int q = (int) token.nval;
        long cnt = 0;
        long sum = 0;
        while (n-- > 0) { // 注意 while 处理多个 case
            token.nextToken();
            int a = (int) token.nval;
            if (a == 0) {
                cnt++;
            } else {
                sum += a;
            }
        }
        PrintWriter p = new PrintWriter(new OutputStreamWriter(System.out));
        while (q-- > 0) {
            token.nextToken();
            long low = (long) token.nval;
            token.nextToken();
            long high = (long) token.nval;
            long a = sum + low * cnt;
            long b = sum + high * cnt;
            p.println(a + " " + b);
        }
        p.flush();
    }
}