import java.util.Random;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;
import java.util.function.Function;

/**
 * 生产者消费者模型
 */
public class ProducerConsumer {
    static volatile int[] store = new int[5];
    static volatile int start = 0;
    static volatile int end = 0;
    static Lock lk1 = new ReentrantLock();
    static Lock lk2 = new ReentrantLock();

    public static class Task implements Runnable {
        Lock lk;
        Function<Thread, Void> act;

        public Task(Lock lock, Function<Thread, Void> action) {
            lk = lock;
            act = action;
        }

        @Override
        public void run() {
            while (true) {
                try {
                    lk.lock();
                    act.apply(Thread.currentThread());
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    // TODO Auto-generated catch block
                    e.printStackTrace();
                } finally {
                    lk.unlock();
                }
            }
        }

    }

    public static Void put(Thread t) {
        if (end - start + 1 < store.length) {
            int val = new Random().nextInt(100);
            System.out.println("Thread-" + t.getName() + " produces " + val);
            store[end++ % store.length] = val;
        }
        return null;
    }

    public static Void take(Thread t) {
        if (start < end) {
            int v = store[(start++) % store.length];
            System.out.println("Thread-" + t.getName() + " consumes " + v);
        }
        return null;
    }

    public static void main(String[] args) throws InterruptedException {
        Thread p1 = new Thread(new Task(lk1, ProducerConsumer::put), "p1");
        Thread p2 = new Thread(new Task(lk1, ProducerConsumer::put), "p2");
        Thread c1 = new Thread(new Task(lk2, ProducerConsumer::take), "c1");
        Thread c2 = new Thread(new Task(lk2, ProducerConsumer::take), "c2");

        p1.start();
        p2.start();
        c1.start();
        c2.start();
        p1.join();
        p2.join();
        c1.join();
        c2.join();
    }
}
