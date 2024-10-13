package interview.intern_coding.taotian20240416;

/**
 * 反转单链表
 */
public class A {
    static class ListNode {
        public ListNode(int i, ListNode head) {
            val = i;
            next = head;
        }
        int val;
        ListNode next;
    }

    public static ListNode reverse(ListNode head) {
        // ListNode dummy = new ListNode(-1, head);
        ListNode pre = null, node = head;

        while (node != null) {
            ListNode tmp = node.next;
            node.next = pre;
            pre = node;
            node = tmp;
        }

        return pre;
        
    }

    public static void main(String[] args) {
        ListNode n1 = new ListNode(1, null);
        ListNode n2 = new ListNode(2, null);
        ListNode n3 = new ListNode(3, null);
        ListNode n4 = new ListNode(4, null);
        ListNode n5 = new ListNode(5, null);
        n1.next = n2;
        n2.next = n3;
        n3.next = n4;
        n4.next = n5;

        ListNode tmp = n1;
        while (tmp != null) {
            System.out.printf(" %d", tmp.val);
            tmp = tmp.next;
        }
        ListNode res = reverse(n1);
        while (res != null) {
            System.out.printf(" %d", res.val);
            res = res.next;
        }

    }
}

