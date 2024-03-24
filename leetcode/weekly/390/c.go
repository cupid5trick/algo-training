package main

// 100258. 最高频率的 ID - 力扣（LeetCode）: https://leetcode.cn/problems/most-frequent-ids/description/
import "container/heap"

type pair struct {
	c int64
	x int
}
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].c > h[j].c } // 最大堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() interface{} {
	old := *h
	*h = old[:len(old)-1]
	return old[0]
}

func mostFrequentIDs(nums []int, freq []int) []int64 {
	n := len(nums)
	cnt := map[int]int64{}

	h := hp{}
	heap.Init(&h)

	ans := make([]int64, n)
	for i, x := range freq {
		cnt[nums[i]] += int64(x)
		heap.Push(&h, pair{cnt[nums[i]], nums[i]})
		for len(h) > 0 && h[0].c != cnt[h[0].x] {
			heap.Pop(&h)
		}

		ans[i] = int64(h[0].c)
	}

	return ans
}
