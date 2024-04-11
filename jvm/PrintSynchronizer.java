import java.util.concurrent.Semaphore;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

/**
 * 实现一个交替输出的线程同步
 */
public class PrintSynchronizer {
    static Lock lk = new ReentrantLock(true);
    public static volatile int num = 0;
    public static int limit = 100;

    public static class Task implements Runnable {
        int n;
        int i;
        public Task(int cnt, int idx) {
            n = cnt;
            i = idx;
        }

        @Override
        public void run() {
            while (num < limit) {
                try {
                    lk.lock();
                    if (num < limit && num % n == i) {
                        System.out.println("Thread-" + Thread.currentThread().getName() + String.format(":%d", num ++));
                    }
                } catch (Exception e) {
                    
                } finally{
                    lk.unlock();
                }
            }
        }
        
    }

    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(new Task(3, 0), "1");
        Thread t2 = new Thread(new Task(3, 1), "2");
        Thread t3 = new Thread(new Task(3, 2), "3");

        t1.start();
        t2.start();
        t3.start();
        t1.join();
        t2.join();
        t3.join();
    }

    static void print1() {
        Semaphore semaphoreOdd = new Semaphore(1);
        Semaphore semaphoreEven = new Semaphore(0);
        Thread oddThread = new Thread(() -> printNumbers(semaphoreOdd, semaphoreEven, true));
        Thread evenThread = new Thread(() -> printNumbers(semaphoreEven, semaphoreOdd, false));

        oddThread.start();
        evenThread.start();
    }


    static void printNumbers(Semaphore currentSemaphore, Semaphore nextSemaphore, boolean isOdd) {
        int startNumber = isOdd ? 1 : 2;
        while (startNumber <= 10) {
            try {
                currentSemaphore.acquire();
                System.out.println(Thread.currentThread().getName() + ": " + startNumber);
                startNumber += 2; // Increment by 2 for odd or even numbers
            } catch (InterruptedException e) {
                e.printStackTrace();
            } finally {
                nextSemaphore.release();
            }
        }
    }


    public static void waitNotify() throws InterruptedException {
        Object l1 = new Object();
        Object l2 = new Object();
        Thread t1 = new Thread(() -> printNumbers(l1, l2, true));
        Thread t2 = new Thread(() -> printNumbers(l1, l1, false));
        t1.start();
        t2.start();
    }

    static void printNumbers(Object lock, Object next, boolean isOdd) {
        int startNumber = isOdd ? 1 : 2;
        while (startNumber <= 10) {
            synchronized (lock) {
                try {

                    System.out.println(Thread.currentThread().getName() + ": " + startNumber);
                    startNumber += 2; // Increment by 2 for odd or even numbers
                    lock.notifyAll();
                    lock.wait();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                } finally {
                }
            }
        }
    }
}
