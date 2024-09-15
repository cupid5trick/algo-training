import java.io.BufferedReader;
import java.io.IOException;
import java.io.StringReader;
import java.util.ArrayList;

/**
 * 小美定义一个字符串的权值为:字符串长度乘以字符的种类数。例如，"arcaea"的权值为6*4=24
 * 现在小美拿到了一个字符串，她希望你将该字符串切割成若干个连续了串，使得每个了串的权值不小于k。请你求出最终最多可以切割出的子串数量。
 * 请注意，由于字符串过长，给出的字符串将是以连续段长度形式给出，例如:aabbaaa将描述为 a(2)b(2)a(3)，aaaaaaaaaaaab 将描述为
 * a(12)b(1)。
 * 输入描述
 * 第一行输入一个两个正整数n,k，代表原字符串长度和每个子串至少应取的权值。第二行是一个仅包含小写字母、数字和括号的字符串。长度不超过10^6
 * 保证所有括号内的数字之和恰好等于 n，给定的每个字母后面必然包含一个括号加数字。l< k,n< 10^18
 * 输出描述
 * 如果整个字符串的权值小于k，请直接输出-1.
 * 否则输出一个正整教，代表可以切割的最多子串数量。
 */
public class D {
    public static void main(String[] args) {
        String input = "a(1234)b(12345)c(678)";
        StringReader rawReader = new StringReader(input);
        try {
            // 使用BufferedReader读取输入
            BufferedReader reader = new BufferedReader(rawReader);
            StringBuilder stringBuilder = new StringBuilder();
            int c;
            // 逐字符读取并构建一个字符串
            while ((c = reader.read()) != -1) {
                stringBuilder.append((char) c);
            }
            // 关闭Reader
            reader.close();

            // 解析字符串
            ArrayList<Character> charList = new ArrayList<>();
            ArrayList<Integer> numList = new ArrayList<>();
            StringBuilder numStringBuilder = new StringBuilder();
            boolean insideNumber = false;
            for (char ch : stringBuilder.toString().toCharArray()) {
                if (ch == '(') {
                    insideNumber = true;
                } else if (ch == ')') {
                    insideNumber = false;
                    numList.add(Integer.parseInt(numStringBuilder.toString()));
                    numStringBuilder.setLength(0);
                } else if (insideNumber) {
                    numStringBuilder.append(ch);
                } else {
                    charList.add(ch);
                }
            }

        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
