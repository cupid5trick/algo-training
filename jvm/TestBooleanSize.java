import java.lang.reflect.Field;
import sun.misc.Unsafe;

/**
 * JVM 伪共享学习
 */
public class TestBooleanSize {
    public static void main(String[] args)
            throws NoSuchFieldException, SecurityException, IllegalArgumentException, IllegalAccessException {
        boolean[] a = new boolean[1024];

        // 使用反射获取 Unsafe 实例
        Field unsafeField = Unsafe.class.getDeclaredField("theUnsafe");
        unsafeField.setAccessible(true);
        Unsafe unsafe = (Unsafe) unsafeField.get(null);

        // 获取基本类型数组中单个元素的大小
        long elementSize = unsafe.arrayIndexScale(a.getClass());

        // 计算整个数组的大小
        long arraySize = elementSize * a.length;

        System.out.println("Size of int array: " + arraySize + " bytes");
    }
}
