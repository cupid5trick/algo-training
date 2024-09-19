package main

/// 1. 求出硬币游戏赢家
func losingPlayer(x int, y int) string {
    turn := 0
    for x >=1 && y >=4 {
        x --
        y -= 4
        turn ++
    }
    if turn % 2 == 0 {
        return "Bob"
    } else {
        return "Alice"
    }
}

/// 2. 操作后字符串最短长度
/// 没理解对样例，很简单
func minimumLength(s string) int {
    stat := [26][]int{}
    for i, c := range []byte(s) {
        stat[c-'a'] = append(stat[c-'a'], i)
    }
    
    cnt := 0
    for _, list := range stat {
        sz := len(list)
        if len(list) <= 2 {
            cnt += sz
        } else {
            for sz > 2 {
                sz -= 2
            }
            cnt +=sz
        }
    }
    return cnt
}

/// 3. 使差值相等的最少数组改动次数
