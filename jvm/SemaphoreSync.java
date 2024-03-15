import java.util.concurrent.Semaphore;

public class SemaphoreSync {

    public static void main(String[] args) throws InterruptedException {
        waitNotify();
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
