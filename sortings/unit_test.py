import os
import subprocess
import errno
from ast import literal_eval

programes = [
    'insert_sort/insert_sort.exe',
    'merge_sort/merge_sort.exe',
    'heap_sort/heap_sort.exe',
    'quick_sort/quick_sort.exe',
]

def random_array(length=20):
    import random
    return ' '.join([str(random.randint(-0x80000000, 0x7fffffff)) for _ in range(length)])

def test(prog, times=1000):
    p = subprocess.Popen([prog],
                         stdin=subprocess.PIPE,
                         stdout=subprocess.PIPE,
                         stderr=subprocess.PIPE,
                         )

    arrays = [random_array() for _ in range(times)]
    input_str = '%d\n%s' % (len(arrays), '\n'.join(['%d %s' % (len(a.split(' ')), a) for a in arrays]))
    # print(input_str)
    try:
        stdout, stderr = p.communicate(
            input_str.encode('utf-8'),
            5,
        )
    except subprocess.TimeoutExpired:
        p.kill()
        return

    lines = stdout.decode('utf-8').strip(' \b\r\n').split('\r\n')
    correct = True
    if p.returncode != 0:

        p.kill()
        print(p.returncode, errno.errorcode[p.returncode], stdout.decode('utf-8'))
        return

    for i, l in enumerate(lines):
        if len(l) != 0:
            # print('[line %d]: ' % (i+1), l, sep='\n')
            l = '[%s]' % l.strip(' \b').replace(' ', ',')
            a = '[%s]' % arrays[i].strip(' \b').replace(' ', ',')
            l = literal_eval(l)
            a = literal_eval(a)
            if l != sorted(a):
                correct = False
                print('[Test Error for input %s]: ' % i, l)
                print('Sorted array should be: ', sorted(l))

    if correct:
        print('Test ok: %s inputs correct' % len(lines))



if __name__ == '__main__':
    for prog in programes:
        test(prog)