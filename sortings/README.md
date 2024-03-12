# Sortings

## 排序算法

- 插入排序：insert sort
- 归并排序：merge sort
- 堆排序：heap sort
- 快速排序：quick sort

## 测试脚本

The test script uses python built-in library *subprocess*.

- Create a Popen instance: ```Popen(<cammand line as list of string>, stdin=PIPE, stdout=PIPE, stderr=PIPE)```
- Interact with process (Run process with optional data and get output): ```Popen.communicate(input=input_str.encode(), timeout=timeout)```
- Terminate process if timeout expires:

```python
try:
  stdout, stderr = p.communicate(data, timeout)
except subprocess.TIMEOUTEXPIRES:
  p.kill()
  ...
```

- Get return value of process: ```p.returncode```
