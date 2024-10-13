package interview.intern_coding.xiaomi;

import java.util.ArrayList;
import java.util.List;

public class TryCatch {
    public static int testTryCatchReturn() {
        int res = 1;
        try {
            res++;
            System.out.println("try ======== res:" + res);
            int a = 1 / 0;
            return res;
        } catch (Exception e) {
            res++;
            System.out.println("catch ======== res:" + res);
            return res;
        } finally {
            res++;
            System.out.println("finally ======== res:" + res);
        }
    }

    public static List testTryCatchReturn1() {
        List res = new ArrayList();
        try {
            res.add(1);
            System.out.println("try ======== res:" + res);
            int a = 1 / 0;
            return res;
        } catch (Exception e) {
            res.add(2);
            System.out.println("catch ======== res:" + res);
            return res;
        } finally {
            res.add(3);
            System.out.println("finally ======== res:" + res);
        }
    }
}
