from collections import Counter


def main():
    """
3
abc
abcbc
abc
acdbc
xyz
xzyxz
    """
    t = int(input())
    while t > 0:
        t = t - 1
        s = input()
        target = input()
        ans = solve1(s, target)
        print(ans)


def solve1(s: str, target: str):
    cnt1 = Counter(s)
    cnt2 = Counter(target)

    for ch1 in cnt2.keys():
        if ch1 not in cnt1.keys():
            return -1

    ans = 0

    j = 0
    while j < len(target):

        for i, ch in enumerate(s):
            if ch == target[j]:
                j = j + 1
        ans = ans + 1
    return ans


if __name__ == '__main__':
    main()
