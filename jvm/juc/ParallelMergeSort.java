import java.io.FileWriter;
import java.io.IOException;
import java.io.RandomAccessFile;
import java.util.Arrays;
import java.util.concurrent.*;

public class ParallelMergeSort {
    public static void main(String[] args) throws InterruptedException, ExecutionException, IOException {
        ForkJoinPool pool = new ForkJoinPool();
        CompletionService<TaskResult> completionService = new ExecutorCompletionService<>(pool);

        int processors = Runtime.getRuntime().availableProcessors();
        int tasks = 20;
        int dataSize = 2000000; // 每个任务处理的数据量
        int totalSize = dataSize * tasks; // 所有任务处理的数据量

        // 初始化原始数组，这里用随机数代替
        int[] array = generateRandomArray(totalSize);

        for (int i = 0; i < tasks; i++) {
            int[] subArray = Arrays.copyOfRange(array, i * dataSize, (i + 1) * dataSize);
            String fileName = "task_" + i + ".txt";
            completionService.submit(new MergeSortTask(subArray, i * dataSize, (i + 1) * dataSize, fileName));
        }

        TaskResult[] results = new TaskResult[tasks];
        int[] starts = new int[tasks];
        int chunksize =  dataSize/(processors);
        for (int i = 0; i < results.length; i++) {
            results[i] = completionService.take().get();
        }
        
        

        pool.shutdown();
    }

    private static int[] generateRandomArray(int size) {
        int[] array = new int[size];
        for (int i = 0; i < size; i++) {
            array[i] = (int) (Math.random() * 1000000); // 生成1-100w之间的随机数
        }
        return array;
    }



    private static void writeToFile(int[] array, String fileName) throws IOException {
        FileWriter writer = new FileWriter(fileName);
        for (int num : array) {
            writer.write(num + "\n");
        }
        writer.close();
    }

    private static int[] readFromFile(String fileName) throws IOException {
        RandomAccessFile file = new RandomAccessFile(fileName, "r");
        int[] array = new int[(int) file.length()];
        int index = 0;
        String line;
        while ((line = file.readLine()) != null) {
            array[index++] = Integer.parseInt(line);
        }
        file.close();
        return array;
    }
}

class MergeSortTask implements Callable<TaskResult> {
    private int[] array;
    private int start;
    private int end;
    private String fileName;

    public MergeSortTask(int[] array, int start, int end, String fileName) {
        this.array = array;
        this.start = start;
        this.end = end;
        this.fileName = fileName;
    }

    @Override
    public TaskResult call() throws Exception {
        // 对子数组进行归并排序
        mergeSort(array);

        // 将排好序的子数组写入文件
        writeToFile(array, fileName);
        return new TaskResult(start, end, fileName);
    }

    private void mergeSort(int[] array) {
        // 归并排序算法，这里省略
        Arrays.sort(array);
    }

    private void writeToFile(int[] array, String fileName) throws IOException {
        FileWriter writer = new FileWriter(fileName);
        for (int num : array) {
            writer.write(num + "\n");
        }
        writer.close();
    }
}

class TaskResult {
    private int start;
    private int end;
    private String fileName;

    public TaskResult(int start, int end, String fileName) {
        this.start = start;
        this.end = end;
        this.fileName = fileName;
    }

    public int getStart() {
        return start;
    }

    public int getEnd() {
        return end;
    }

    public String getFileName() {
        return fileName;
    }
}
