package interview;
import java.time.Duration;
import java.util.*;
/**
 * 分析下面的题目，应该采用什么方法解决问题？
小歪在操场做志愿者时，观察篮球收纳盒中篮球编号变化的过程，并试图找出可以形成的不同数字序列的数量。具体规则如下：

1. **操作定义**：
   - `push x`：表示有人将编号为x的篮球放入盒内。
   - `pop`：表示有人从盒中取走了最后放入的篮球。如果盒子为空，则忽略这一步操作。

2. **数字序列形成**：
   - 将盒内有的篮球编号按放入顺序依次连接起来，可以得到一个数字序列。

3. **目标**：
   - 计算在整个过程中，可以形成多少个不同的数字序列。

4. **输入描述**：
   - 第一行输入一个整数n（1≤n≤2×10^5），代表小歪的记录次数。
   - 随后n行，每行首先输入一个字符串s（s∈{push, pop}），代表这一条记录的类别。
   - 如果s=push，在同一行上输入一个整数x（1≤x≤2×10^5），代表新放入一个编号为x的篮球，编号可能会有相同。
   - 如果s=pop，代表最后放入的篮球被取出。

5. **输出描述**：
   - 输出在整个过程中可以形成的不同数字序列的数量。

题目提供了两个示例来说明数字序列的变化过程和计算方法：

**示例1**：
- 输入：
  ```
  push 11
  push 4
  push 51
  pop
  pop
  push 45
  ```
- 输出：
  ```
  4
  ```
- 说明：数字序列的变化过程如下：
  - 放入一个篮球，{11}；
  - 放入一个篮球，{11,4}；
  - 放入一个篮球，{11,4,51}；
  - 取走最后放入的篮球，{11,4}；
  - 取走最后放入的篮球，{11}；
  - 放入一个篮球，{11,45}。
  - 一共可以得到4个不同的数字序列：{11}，{11,4}，{11,4,51}，{11,45}。

**示例2**：
- 输入：
  ```
  push 11
  push 4
  push 51
  pop
  pop
  push 45
  ```
- 输出：
  ```
  4
  ```
- 说明：与示例1相同，得到4个不同的数字序列。
 */
public class UniqueSequences {

    // 定义字典树的节点
    static class TrieNode {
        Map<Integer, TrieNode> children = new HashMap<>();
        boolean end = false;
    }

    static class Trie {
        TrieNode root;
        int uniqueSequences;

        public Trie() {
            root = new TrieNode();
            uniqueSequences = 0;
        }

        // 向字典树中插入一个序列，返回是否插入了新序列
        public boolean insert(List<Integer> sequence) {
            TrieNode current = root;
            boolean isNewSequence = false;

            for (int num : sequence) {
                if (!current.children.containsKey(num)) {
                    current.children.put(num, new TrieNode());
                }
                current = current.children.get(num);
            }

            if (!current.end) {
                current.end = true;
                isNewSequence = true;
            }

            return isNewSequence;
        }
    }

    public static int countUniqueSequences(List<String[]> operations) {
        Trie trie = new Trie();
        Stack<Integer> stack = new Stack<>();

        for (String[] op : operations) {
            if (op[0].equals("push")) {
                stack.push(Integer.parseInt(op[1]));
            } else if (op[0].equals("pop")) {
                if (!stack.isEmpty()) {
                    stack.pop();
                }
            }

            List<Integer> currentSequence = new ArrayList<>(stack);
            if (trie.insert(currentSequence)) {
                trie.uniqueSequences++;
            }
        }

        return trie.uniqueSequences;
    }

    public static void main(String[] args) {
        long t1 = System.currentTimeMillis();
        Scanner scanner = new Scanner(System.in);
        
        // 读取记录次数
        int n = Integer.parseInt(scanner.nextLine());
        List<String[]> operations = new ArrayList<>();

        // 读取每一条操作
        for (int i = 0; i < n; i++) {
            String line = scanner.nextLine();
            String[] parts = line.split(" ");
            operations.add(parts);
        }

        // 计算不同的数字序列数量并输出
        System.out.println(countUniqueSequences(operations));
        System.out.println("time used:" + Duration.ofMillis(System.currentTimeMillis()-t1).toString());
    }
}
