n,m = map(int, input().split())

sc = [0]*m
tmp = [[] for _ in range(m)]
for _ in range(n):
    s = input()
    for i, c in enumerate(s):
        score = ord(c) if 'a' <= c <= 'z' else ord(c)-ord('A')+ord('a')
        sc[i] += score-ord('a')
        tmp[i].append(score-ord('a'))

for i in range(m):
    tmp[i].sort()
    sc[i] -= tmp[i][0]+tmp[i][-1]
    sc[i] = chr((sc[i]-1)//(n-2)+1+ord('a'))

res = list(zip(sc, range(1,m+1)))
res.sort()
print(res)

print(' '.join([str(res[i][1]) for i in range(m)]))
print(' '.join([res[i][0] for i in range(m)]))
