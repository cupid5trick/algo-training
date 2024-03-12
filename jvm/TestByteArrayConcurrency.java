import java.util.Arrays;
import java.util.concurrent.atomic.AtomicBoolean;

public class TestByteArrayConcurrency {
    private static final int ARRAY_SIZE = 16;
    private static final byte[] byteArray = new byte[ARRAY_SIZE];
    private static final AtomicBoolean foundMismatch = new AtomicBoolean(false);

    public static void main(String[] args) {
        // 初始化byte数组
        for (int i = 0; i < ARRAY_SIZE; i++) {
            byteArray[i] = (byte) i;
        }

        // Create and start background thread to check for mismatches
        Thread checkerThread = new Thread(() -> {
            while (!foundMismatch.get()) {
                for (int i = 0; i < ARRAY_SIZE; i++) {
                    if (byteArray[i] != (byte) i) {
                        foundMismatch.set(true);
                        System.out.println("Mismatch detected at index " + i);
                        break;
                    }
                }
            }
        });
        checkerThread.setDaemon(true);
        checkerThread.start();

        // Create and start worker threads
        Thread[] threads = new Thread[ARRAY_SIZE];
        int thread_cnt = 16;

        // 创建并启动16个线程并发修改byte数组
        for (int i = 0; i < thread_cnt; i++) {
            final int index = i % ARRAY_SIZE;
            threads[index] = new Thread(() -> {
                while (!foundMismatch.get()) {
                    // 并发修改byte数组
                    byteArray[index] = (byte) index;

                    // 检查是否有元素不等于其下标，如果有则设置foundMismatch为true
                    for (int j = 0; j < ARRAY_SIZE; j++) {
                        if (byteArray[j] != j) {
                            foundMismatch.set(true);
                            System.out.printf("Found mismatch at index: %d, %s" + j, Arrays.toString(byteArray));
                            break;
                        }
                    }
                }
            });
            threads[index].start();
        }

        // Wait for all worker threads to finish
        for (Thread thread : threads) {
            try {
                thread.join();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

    }
}
