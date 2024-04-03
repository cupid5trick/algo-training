import java.lang.reflect.Field;
import sun.misc.Unsafe;
public class TestJavaString {
    static final String str = "HELLO".intern();
    // static final String str = "HELLO";

    static final Unsafe unsafe = getUnsafe();
    static final Unsafe getUnsafe() {
        {
            try {
                Field field = Unsafe.class.getDeclaredField("theUnsafe");
                field.setAccessible(true);
                return (Unsafe )field.get(null);
            } catch (NoSuchFieldException | SecurityException | IllegalArgumentException | IllegalAccessException e) {
                // TODO Auto-generated catch block
                e.printStackTrace();
            }
            
        }
        return null;
    }


    public static void main(String[] args) throws Exception {
        // String str = "Hello";
        String tmp = str;
        System.out.println("Original String: " + str);

        // 使用反射获取 String 中的 value 字段（char 数组）
        Field valueField = String.class.getDeclaredField("value");
        valueField.setAccessible(true); // 设置可访问

        // 获取 String 对象中的 char 数组
        char[] charArray = (char[]) valueField.get(str);

        // 尝试修改 char 数组中的内容
        charArray[0] = 'C';

        // 打印修改后的 String
        System.out.println("Modified String: " + str);
        System.out.println(str == tmp);
        
        
        long strOffset = unsafe.staticFieldOffset(TestJavaString.class.getDeclaredField("str"));
        System.out.println(String.format("0x%16s", Long.toHexString(strOffset)).replace(' ', '0'));

        System.out.printf("page size: %d, address size: %d\n", unsafe.pageSize(), unsafe.addressSize());
    }
}
