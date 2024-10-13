package interview.intern_coding.taotian20240416;

import java.util.Deque;
import java.util.LinkedList;
import java.util.Stack;


/**
 * 一道不带优先级的括号匹配
 */
public class B {

    public static boolean check(char a, char b) {
        return a == '(' && b == ')' || a == '[' && b == ']' || a == '{' && b == '}';
    }
    
    public static void main(String[] args) {
        Deque<Character> stk = new LinkedList<>();
        String input = "(}";
        boolean ans = true;
        for (char c : input.toCharArray()) {

            if (c == '(' || c == '[' || c == '{') {
                stk.offerLast(c);
            } else if (!check(stk.getLast(), c)) {
                // System.out.printf("%c %c\n", c, stk.getLast());
                stk.pollLast();
                ans = false;
            } 
        }

        System.out.println(ans);
    }
}
