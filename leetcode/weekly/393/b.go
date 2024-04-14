// 100265. 素数的最大距离 - 力扣（LeetCode）: https://leetcode.cn/problems/maximum-prime-difference/description/
// 很简单。预处理找出所有质数，然后遍历数组求素数下标之间的最大距离。
// 可以考虑双指针优化，不过复杂度还是 O(n)

func isPrime(x int) (ans bool) {
    ans = true
    for i:=2; i *i <= x; i ++{
        if x % i == 0 {
            ans = false
        }
    }
    return
}
func maximumPrimeDifference(nums []int) int {
    primes := make([]bool, 110)
    for x:=2; x <= 100; x ++ {
        if isPrime(x) {
            primes[x] = true
        }
    }
    
    st := []int{}
    for i, x := range nums {
        if primes[x] {
            st = append(st, i)
        }
    }
    m := len(st)
    ans := st[m-1] - st[0] 
    return ans
}