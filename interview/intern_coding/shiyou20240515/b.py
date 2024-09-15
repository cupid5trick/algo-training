def main():
    """
4
bge)))))))))
((IIII))))))
()()()()(uuu
))))UUUU((()
    """
    t = int(input())
    while t  > 0:
        t = t- 1
        s = input()
        st = []
        mismatches = []
        for i, ch in enumerate(list(s)):
            if ch == '(':
                st.append((ch, i))
            elif ch == ')':
                if len(st)>0 and st[-1][0] == '(':
                    st = st[:-1]
                    print('ppop')
                else:
                    mismatches.append((2,i))

        for v in st:
            mismatches.append((1, v[1]))
            
        # print(mismatches)

        ans = [' '] * len(s)
        for v in mismatches:
            typ, idx = v
            if typ == 1:
                ans[idx] = 'x'
            elif typ == 2:
                ans[idx] = '?'

        # print(ans)
        print(s)
        print(''.join(ans))

if __name__ == '__main__':
    main()
    