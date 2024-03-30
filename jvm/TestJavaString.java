import java.lang.reflect.Field;

public class TestJavaString {
    static final String str0 = "HELLO".intern();
    static final String str = "HELLO";


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

        double f = Math.floor(1.4D);
        System.out.printf("%.10f\n", f);
    }

}
