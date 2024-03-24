
import java.lang.String;
import java.util.Objects;

public class HashCodeDemo {

    static class T {
        Integer a;
        String b;
        long c;

        public int hashCode() {
            /// 某种对 a, b, c 的哈希计算
            
            return Objects.hash(a, b, c);
        }

        public boolean equals(Object other) {
            if (this == other) {
                return true;
            }
            if (!(other instanceof T) || other == null) {
                return false;
            }

            boolean flag = true;
            T another = (T) other;
            flag = (this.a != null && this.a.equals(another.a)) || another.a == null;
            flag = (this.b != null && this.b.equals(another.b)) || another.b == null;
            return flag && this.hashCode() == another.hashCode() && c == another.c;
        }
    }

    static class T2 {
        Integer a;
        String b;
        long c;

        @Override
        public boolean equals(Object o) {
            if (this == o)
                return true;
            if (o == null || getClass() != o.getClass())
                return false;
            T2 t = (T2) o;
            return c == t.c &&
                    Objects.equals(a, t.a) &&
                    Objects.equals(b, t.b);
        }

        @Override
        public int hashCode() {
            return Objects.hash(a, b, c);
        }
    }

    public static void main(String[] args) {

    }
}