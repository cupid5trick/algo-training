
import java.util.*;
/**
 * From lyt.
 */
public class LeetcodeHots {

    static class ListNode {
        int val;
        ListNode next;

        public ListNode(int v) {this.val = v;}
    }

    static class TreeNode {
        TreeNode left;
        TreeNode right;
        int val;
    }
    /**
     * 一、数组
     */
    // 1.全排列 -- 回溯
    public List<List<Integer>> permute(int[] nums) {
        int len = nums.length;
        List<List<Integer>> res = new ArrayList<>();
        if (len == 0) {
            return res;
        }
        List<Integer> path = new ArrayList<>();
        boolean[] used = new boolean[len];
        dfs(nums, len, 0, used, path, res);
        return res;
    }

    public static void dfs(int[] nums, int len, int index, boolean[] used, List<Integer> path,
            List<List<Integer>> res) {
        // 递归结束条件
        if (index == len) {
            res.add(new ArrayList<>(path));
            return;
        }

        for (int i = 0; i < len; i++) {
            if (!used[i]) {
                used[i] = true;
                path.add(nums[i]);
                dfs(nums, len, index + 1, used, path, res);
                path.remove(path.size() - 1);
                used[i] = false;
            }
        }
    }

    // 2.前k个高频元素 -- 最小堆
    public int[] topKFrequent(int[] nums, int k) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int num : nums) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }
        // 维护一个元素数目为k的最小堆
        PriorityQueue<Map.Entry<Integer, Integer>> priorityQueue = new PriorityQueue<>(
                new Comparator<Map.Entry<Integer, Integer>>() {
                    @Override
                    public int compare(Map.Entry<Integer, Integer> o1, Map.Entry<Integer, Integer> o2) {
                        return o1.getValue() - o2.getValue();
                    }
                });

        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            if (priorityQueue.size() == k) {
                if (priorityQueue.peek().getValue() < entry.getValue()) {
                    priorityQueue.poll();
                    priorityQueue.offer(entry);
                }
            } else {
                priorityQueue.offer(entry);
            }
        }
        int[] res = new int[k];
        for (int i = 0; i < k; i++) {
            res[i] = priorityQueue.poll().getKey();
        }
        return res;
    }

    // 3.删除有序数组中的重复项
    public int removeDuplicates(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        int left = 0;
        int right = 1;
        while (right < nums.length) {
            if (nums[left] != nums[right]) {
                nums[left + 1] = nums[right];
                left++;
            }
            right++;
        }
        return left + 1;
    }

    /**
     * 二、链表
     */
    // 0.反转链表
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode nextHead = reverseList(head.next);
        head.next.next = head;
        head.next = null;
        return nextHead;
    }

    // 1.部分反转单向链表
    public ListNode reverseBetween(ListNode head, int left, int right) {
        ListNode dummy = new ListNode(-1);
        dummy.next = head;
        ListNode pre = dummy;
        // 头插法，pre位置不变，永远指向left
        for (int i = 0; i < left - 1; i++) {
            pre = pre.next;
        }
        // cur永远指向待反转区域的第一个节点
        ListNode cur = pre.next;
        // next永远执行cur的下一个节点
        ListNode next;
        for (int i = 0; i < right - left; i++) {
            next = cur.next;
            cur.next = next.next;
            next.next = pre.next;
            pre.next = next;
        }
        return dummy.next;
    }

    // 2.删除链表的倒数第N个节点
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode dummy = new ListNode(-1);
        dummy.next = head;
        ListNode fast = head;
        ListNode slow = dummy;
        while (n > 0) {
            fast = fast.next;
            n--;
        }
        while (fast != null) {
            fast = fast.next;
            slow = slow.next;
        }
        slow.next = slow.next.next;
        return dummy.next;
    }

    // 3.合并k个升序链表

    /**
     * 分治法：
     * k个链表配对并将同一对中的链表合并
     * 一直重复直到得出答案
     */
    public ListNode mergeKLists(ListNode[] lists) {
        return merge(lists, 0, lists.length - 1);
    }

    public ListNode merge(ListNode[] lists, int l, int r) {
        if (l == r) {
            return lists[l];
        }
        if (l > r) {
            return null;
        }
        int mid = l + (r - l) / 2;
        return mergeTwoLists(merge(lists, l, mid), merge(lists, mid + 1, r));
    }

    public ListNode mergeTwoLists(ListNode a, ListNode b) {
        if (a == null || b == null) {
            return a != null ? a : b;
        }
        ListNode head = new ListNode(0);
        ListNode tail = head;
        ListNode aPtr = a, bPtr = b;
        while (aPtr != null && bPtr != null) {
            if (aPtr.val < bPtr.val) {
                tail.next = aPtr;
                aPtr = aPtr.next;
            } else {
                tail.next = bPtr;
                bPtr = bPtr.next;
            }
            tail = tail.next;
        }
        tail.next = (aPtr != null ? aPtr : bPtr);
        return head.next;
    }

    /**
     * 三、树
     */
    // 0.求二叉树的最大距离
    /*
     * a.两个节点都是叶子节点 -- 最远距离就是左边最深的深度加上右边最深的深度
     */
    int maxDistance = 0;

    public int maxDistanceOfChildren(TreeNode root) {
        if (root == null) {
            return 0;
        }
        if (root.left == null && root.right == null) {
            return 0;
        }
        // 递归求左右子树高度
        int leftHeight = maxDistanceOfChildren(root.left) + 1;

        int rightHeight = maxDistanceOfChildren(root.right) + 1;

        int distance = leftHeight + rightHeight;

        maxDistance = Math.max(maxDistance, distance);
        // 该节点最大深度等于它左右孩子最大深度中较大的那个
        return Math.max(leftHeight, rightHeight);
    }

    // 1.判断是否是平衡二叉树（一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 ）
    public boolean isBalanced(TreeNode root) {
        if (root == null) {
            return true;
        }
        return maxDepth(root) >= 0;
    }

    public int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int left = maxDepth(root.left);
        int right = maxDepth(root.right);
        if (left >= 0 && right >= 0 && Math.abs(left - right) <= 1) {
            return Math.max(left, right) + 1;
        } else {
            return -1;
        }
    }

    // 2.二叉树的最小深度
    public int minDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        if (root.left == null && root.right == null) {
            return 1;
        }
        int left = minDepth(root.left);
        int right = minDepth(root.right);

        if (root.left == null || root.right == null) {
            return left + right + 1;
        }
        return Math.min(left, right) + 1;
    }

    // 3.组合
    // 77 组合
    List<List<Integer>> res = new ArrayList<>();
    List<Integer> path = new ArrayList<>();

    public List<List<Integer>> combine(int n, int k) {
        dfs(n, k, 1);
        return res;
    }

    public void dfs(int n, int k, int index) {
        if (path.size() == k) {
            res.add(new ArrayList<>(path));
            return;
        }
        for (int i = index; i <= n - (k - path.size()) + 1; i++) {
            path.add(i);
            dfs(n, k, i + 1);
            path.remove(path.size() - 1);
        }
    }

    // 216 组合总合3
    public List<List<Integer>> combinationSum3(int k, int n) {
        dfs2(k, 0, n, 1);
        return res;
    }

    public void dfs2(int k, int sum, int target, int index) {
        if (path.size() == k) {
            if (sum == target) {
                res.add(new ArrayList<>(path));
            }
            return;
        }
        for (int i = index; i <= 9 - (k - path.size()) + 1; i++) {
            path.add(i);
            sum += i;
            dfs2(k, sum, target, i + 1);
            sum -= i;
            path.remove(path.size() - 1);
        }
    }

    // 39 组合总和
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        Arrays.sort(candidates);
        dfs3(candidates, target, 0);
        return res;
    }

    public void dfs3(int[] candidates, int rest, int index) {
        if (rest == 0) {
            res.add(new ArrayList<>(path));
            return;
        }
        if (rest < 0) {
            return;
        }
        for (int i = index; i < candidates.length && rest - candidates[i] >= 0; i++) {
            rest -= candidates[i];
            path.add(candidates[i]);
            dfs3(candidates, rest, i);
            rest += candidates[i];
            path.remove(path.size() - 1);
        }
    }

    // 40 组合总和2
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        Arrays.sort(candidates);
        boolean[] used = new boolean[candidates.length];
        dfs4(candidates, target, 0, used);
        return res;
    }

    public void dfs4(int[] candidates, int rest, int index, boolean[] used) {
        if (rest < 0) {
            return;
        }
        if (rest == 0) {
            res.add(new ArrayList<>(path));
            return;
        }
        for (int i = index; i < candidates.length && rest - candidates[i] >= 0; i++) {
            if (i > 0 && candidates[i - 1] == candidates[i] && !used[i - 1]) {
                continue;
            }
            used[i] = true;
            path.add(candidates[i]);
            rest -= candidates[i];
            dfs4(candidates, rest, i + 1, used);
            rest += candidates[i];
            path.remove(path.size() - 1);
            used[i] = false;
        }
    }

    // 17 电话号码字母组合
    List<String> strRes = new ArrayList<>();

    public List<String> letterCombinations(String digits) {
        if (digits == null || digits.length() == 0) {
            return strRes;
        }
        String[] strs = new String[] {
                "",
                "",
                "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"
        };
        dfs5(strs, digits, 0);
        return strRes;
    }

    StringBuilder tmp = new StringBuilder();

    public void dfs5(String[] strs, String digits, int index) {
        if (index == digits.length()) {
            strRes.add(tmp.toString());
            return;
        }
        // 当前层对应数字
        int digit = digits.charAt(index) - '0';
        // 当前层对应字符集
        String str = strs[digit];
        for (int i = 0; i < str.length(); i++) {
            tmp.append(str.charAt(i));
            dfs5(strs, digits, index + 1);
            tmp.deleteCharAt(tmp.length() - 1);
        }
    }

    // 4.节点与祖先之间最大差值
    int maxDiff = 0;

    public int maxAncestorDiff(TreeNode root) {
        return dfs6(root, root.val, root.val);
    }

    public int dfs6(TreeNode root, int mi, int ma) {
        if (root == null) {
            return 0;
        }
        // 用上一层计算
        maxDiff = Math.max(Math.abs(root.val - mi), Math.abs(ma - root.val));
        // 更新最大和最小的祖先
        int min = Math.min(mi, root.val);
        int max = Math.max(ma, root.val);
        // 递归计算左右子树的最大差值
        maxDiff = Math.max(maxDiff, dfs6(root.left, min, max));
        maxDiff = Math.max(maxDiff, dfs6(root.right, min, max));
        return maxDiff;
    }

    /**
     * 四、排序
     */
    // 1.堆排序
    // 堆排序 O(nlogn)
    public void heapSort(int[] arr) {
        // 建堆 O(n) 自底向上
        for (int i = arr.length - 1; i >= 0; i--) {
            heapAdjust(arr, i, arr.length);
        }
        // 自底向上交换根节点和最后一个节点并调整
        for (int i = arr.length - 1; i >= 0; i--) {
            int tmp = arr[0];
            arr[0] = arr[i];
            arr[i] = tmp;
            // 每弹出一个堆顶len--
            heapAdjust(arr, 0, i);
        }
    }

    // 堆调整 O(nlogn)
    // 调整以parent为根节点的子树
    public void heapAdjust(int[] arr, int parent, int len) {
        // 记录调整的根节点值
        int tmp = arr[parent];
        // 初始根节点对应左孩子节点的下标
        int child = parent * 2 + 1;
        while (child < len) {
            // 判断右孩子是不是更大
            if (child + 1 < len && arr[child + 1] > arr[child]) {
                child++;
            }
            // 如果父节点比孩子都大，则直接退出（自底向上）
            if (tmp >= arr[child]) {
                break;
            }
            // 父节点下沉、更新孩子节点
            arr[parent] = arr[child];
            parent = child;
            child = parent * 2 + 1;
        }
        // 最初的父节点值放在下沉的最后的位置
        arr[parent] = tmp;
    }

    // 2.快排
    Random random = new Random();

    public void quickSort(int[] nums, int low, int high) {
        if (low >= high) {
            return;
        }
        int i = low, j = high;
        // 随机选取分割点
        int rand = random.nextInt(high - low + 1) + low;
        int p = nums[rand];
        // 将分割点放在第一个
        int tmp = nums[rand];
        nums[rand] = nums[low];
        nums[low] = tmp;

        while (i < j) {
            while (nums[j] >= p && i < j) {
                j--;
            }
            while (nums[i] <= p && i < j) {
                i++;
            }
            if (i < j) {
                int temp = nums[i];
                nums[i] = nums[j];
                nums[i] = temp;
            }
        }

        nums[low] = nums[i];
        // 将分割点放在正确位置
        nums[i] = p;
        // 递归处理左右子问题
        quickSort(nums, low, i - 1);
        quickSort(nums, i + 1, high);
    }

    // 3.归并排序
    public static void merge(int[] arr, int l, int mid, int r) {
        int[] tmp = Arrays.copyOfRange(arr, l, r + 1);
        int i = l, j = mid + 1;
        for (int k = l; k <= r; k++) {
            if (i > mid) {
                // 左边遍历完了
                arr[k] = tmp[j - l];
                j++;
            } else if (j > r) {
                // 右边遍历完了
                arr[k] = tmp[i - l];
                i++;
            } else if (tmp[i - l] < tmp[j - l]) {
                arr[k] = tmp[i - l];
                i++;
            } else {
                arr[k] = tmp[j - l];
                j++;
            }
        }
    }

    public static void mergeSort(int[] arr, int l, int r) {
        if (l >= r) {
            return;
        }
        int mid = l + (r - l) / 2;
        mergeSort(arr, l, mid);
        mergeSort(arr, mid + 1, r);
        if (arr[mid] > arr[mid + 1]) {
            merge(arr, l, mid, r);
        }
    }

    /**
     * 五、位图
     */
    // 1.在[1 .. 100]范围内有1亿个数，找出重复次数最多的那个数字，假设重复次数互不相同
    public static void repeatTopN(List<Integer> data, int topN) {
        // 存放key:数字 value:出现次数
        Map<Integer, Integer> map1 = new HashMap<Integer, Integer>();
        // 存放 key:出现次数 value:数字
        Map<Integer, List<Integer>> map2 = new HashMap<Integer, List<Integer>>();

        // 填充两个map
        for (int number : data) {
            int count = 1;

            if (map1.containsKey(number)) {
                count = map1.get(number) + 1;
                map1.put(number, count);
            } else {
                map1.put(number, 1);
            }

            if (map2.containsKey(count)) {
                map2.get(count).add(number);
            } else {
                List list = new ArrayList<Integer>();
                list.add(number);
                map2.put(count, list);
            }
        }
        // 使用位图
        BitSet bs = new BitSet();
        // 填充位图
        // 设置出现次数为bs对应下标
        for (Map.Entry<Integer, Integer> entry : map1.entrySet()) {
            bs.set(entry.getValue());
        }
        // 注意：这里使用TreeSet 是为了解决上面最多出现次数的数字会在之前的list都出现一遍，为了去重和按照顺序插入到result中
        Set<Integer> result = new TreeSet<Integer>();
        // 倒序循环bs。只要值为true的。说明是出现过的最大次数。
        // 然后倒序输出，使用出现次数从map2中查找对应的数字列表, 查找topN个数字

        a: for (int j = bs.size(); j >= 0; j--) { // 这里bs的size是2的n次幂。
            if (bs.get(j)) {
                b: for (Integer num : map2.get(j)) {
                    if (result.size() > topN) { // result的大小 大于用户需要取出的topN。退出循环
                        break a;
                    }
                    // 取出重复次数前topN的数字
                    boolean flag = result.add(num);
                    if (flag) {// 如果是不重复并且增加进去的则输出数字出现的次数。
                        System.out.println("[" + num + "]出现过:" + j + "次");
                    }
                }
            }
        }
        System.out.println(result);
    }

    // public static void main(String[] args) {
    // Random r = new Random();
    // List<Integer> data = new ArrayList<>();
    // for (int i = 0; i < 10000; i++){
    // data.add(r.nextInt(1000));
    // }
    // repeatTopN(data,5);
    // }

    /**
     * 八、字符串
     */
    // 1.将句子排序
    public String sortSentence(String s) {
        String[] strings = s.split(" ");
        if (strings.length < 2) {
            return s.substring(0, s.length() - 1);
        }
        StringBuilder stringBuilder = new StringBuilder("");
        Arrays.sort(strings, (s1, s2) -> (s1.charAt(s1.length() - 1) - 'a') - (s2.charAt(s2.length() - 1) - 'a'));
        for (String string : strings) {
            stringBuilder.append(string.substring(0, string.length() - 1) + " ");
        }
        stringBuilder.deleteCharAt(stringBuilder.length());
        return stringBuilder.toString();
    }

    // 2.最长公共前缀
    // a.普通扫描
    public String longestCommonPrefix1(String[] strs) {
        if (strs.length < 2) {
            return strs[0];
        }
        String res = strs[0];
        int len = strs.length;
        for (int i = 1; i < strs.length; i++) {
            res = helper1(res, strs[i]);
            if (res == "") {
                return res;
            }
        }
        return res;
    }

    public String helper1(String s1, String s2) {
        String res = "";
        int i = 0, j = 0;
        while (i < s1.length() && j < s2.length()) {
            if (s1.charAt(i) != s2.charAt(j)) {
                break;
            }
            i++;
            j++;
        }
        res = s1.substring(0, i);
        return res;
    }

    // b.分治法
    public String longestCommonPrefix2(String[] strs) {
        if (strs.length < 2) {
            return strs[0];
        } else {
            return helper2(strs, 0, strs.length - 1);
        }
    }

    public String helper2(String[] strs, int start, int end) {
        if (start == end) {
            return strs[start];
        } else {
            int mid = start + (end - start) / 2;
            String leftCommon = helper2(strs, start, mid);
            String rightCommon = helper2(strs, mid + 1, end);
            int i = 0, j = 0;
            while (i < leftCommon.length() && j < rightCommon.length()) {
                if (leftCommon.charAt(i) != rightCommon.charAt(j)) {
                    return leftCommon.substring(0, i);
                }
                i++;
                j++;
            }
            return leftCommon.substring(0, i);
        }
    }

    // c.二分
    public String longestCommonPrefix(String[] strs) {
        if (strs.length < 2) {
            return strs[0];
        }
        int minLen = Integer.MAX_VALUE;
        for (String s : strs) {
            minLen = Math.min(minLen, s.length());
        }
        int low = 0, high = minLen;
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (helper3(strs, mid)) {
                // mid长度是公共前缀，那么最终结果长度一定大于或等于mid
                low = mid;
            } else {
                high = mid - 1;
            }
        }
        return strs[0].substring(0, low);
    }

    public boolean helper3(String[] strs, int len) {
        String str = strs[0].substring(0, len);
        for (int i = 0; i < strs.length; i++) {
            for (int j = 0; j < len; j++) {
                if (strs[i].charAt(j) != str.charAt(j)) {
                    return false;
                }
            }
        }
        return true;
    }

    // 3.计算器
    public int calculate(String s) {
        char[] chars = s.toCharArray();
        // 栈内存储的是最后需要相加的
        Stack<Integer> stack = new Stack<>();
        int curNum = 0;
        int preOp = '+';
        for (int i = 0; i < s.length(); i++) {
            if (Character.isDigit(s.charAt(i))) {
                curNum = curNum * 10 + s.charAt(i) - '0';
            }
            if (!Character.isDigit(s.charAt(i)) && s.charAt(i) != ' '
                    || i == s.length() - 1) {
                switch (preOp) {
                    case '+':
                        stack.push(curNum);
                        break;
                    case '-':
                        stack.push(-curNum);
                        break;
                    case '*':
                        stack.push(stack.pop() * curNum);
                        break;
                    default:
                        stack.push(stack.pop() / curNum);
                }
                preOp = s.charAt(i);
                curNum = 0;
            }
        }
        int res = 0;
        while (!stack.isEmpty()) {
            res += stack.pop();
        }
        return res;
    }

    // 4.最长不含重复字符的子字符串
    // 滑动窗口
    public static int lengthOfLongestSubstring(String s) {
        if (s == null || s.length() == 0) {
            return 0;
        }
        // 用来判断当前窗口内是否有重复字符
        HashSet<Character> set = new HashSet<>();
        set.add(s.charAt(0));
        int l = 0, r = 0;
        int maxLen = 1;
        while (l < s.length()) {
            if (l > 0) {
                set.remove(s.charAt(l - 1));
            }
            while (r + 1 < s.length() && !set.contains(s.charAt(r + 1))) {
                set.add(s.charAt(r + 1));
                r++;
            }
            maxLen = Math.max(maxLen, r - l + 1);
            l++;
        }
        return maxLen;
    }

    // public static void main(String[] args) {
    // String s = "pwwkew";
    // System.out.println(lengthOfLongestSubstring(s));
    // }
    // 5.单词子集
    // 记录words2字符串中所有字符的最大出现次数
    // 遍历words1字符串，若能够包含记录中的所有字符，则加入结果集
    public List<String> wordSubsets(String[] words1, String[] words2) {
        int n = words1.length;
        int m = words2.length;

        int[] count = new int[26];
        for (int i = 0; i < m; i++) {
            int[] tmp = new int[26];
            String s = words2[i];
            for (int j = 0; j < s.length(); j++) {
                tmp[s.charAt(j) - 'a']++;
            }
            for (int j = 0; j < 26; j++) {
                count[j] = Math.max(count[j], tmp[j]);
            }
        }
        List<String> list = new ArrayList<>();
        a: for (int i = 0; i < n; i++) {
            int[] tmp = new int[26];
            String s = words1[i];
            for (int j = 0; j < s.length(); j++) {
                tmp[s.charAt(j) - 'a']++;
            }
            for (int j = 0; j < 26; j++) {
                if (count[j] > tmp[j]) {
                    continue a;
                }
            }
            list.add(s);
        }
        return list;
    }

    // 6.最长回文子串
    // dp
    public String longestPalindrome1(String s) {
        int len = s.length();
        if (len < 2) {
            return s;
        }
        int start = 0;
        int maxLen = 0;
        // s[i...j]是不是回文串
        boolean[][] dp = new boolean[len][len];
        for (int i = 0; i < len; i++) {
            dp[i][i] = true;
        }
        char[] chars = s.toCharArray();
        // 枚举子串长度
        for (int l = 2; l <= len; l++) {
            for (int i = 0; i < len; i++) {
                int j = i + l - 1;

                if (j >= len) {
                    break;
                }
                if (chars[i] != chars[j]) {
                    dp[i][j] = false;
                } else {
                    if (j - i < 3) {
                        dp[i][j] = true;
                    } else {
                        dp[i][j] = dp[i + 1][j - 1];
                    }
                }
                // 更新start和max
                if (dp[i][j] && j - i + 1 > maxLen) {
                    maxLen = j - i + 1;
                    start = i;
                }
            }
        }
        return s.substring(start, start + maxLen);
    }

    // 中心扩展
    public String longestPalindrome(String s) {
        if (s == null || s.length() == 0) {
            return "";
        }
        int start = 0, end = 0;
        for (int i = 0; i < s.length(); i++) {

        }
        return "";
    }
}
