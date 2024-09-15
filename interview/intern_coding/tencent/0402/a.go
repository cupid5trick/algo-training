package main

import (
	"container/heap"
	. "fmt"
)

// 实现LRU缓存
// 运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
// 获取数据 get(key)：
//     如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1
// 写入数据 put(key, value)：
//     如果关键字已经存在，则变更其数据值
//     如果关键字不存在，则插入该组「关键字/值」
//     当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间
// 示例：

// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1); // 返回 1
// cache.put(3, 3); // 该操作会使得关键字 2 作废
// cache.get(2); // 返回 -1 (未找到)
// cache.put(4, 4); // 该操作会使得关键字 1 作废
// cache.get(1); // 返回 -1 (未找到)
// cache.get(3); // 返回 3
// cache.get(4); // 返回 4

type item struct {
	t int
	k int
}

type myheap struct {
	arr []item
}

func (hp myheap) Len() int           { return len(hp.arr) }
func (hp myheap) Swap(a, b int)      { hp.arr[a], hp.arr[b] = hp.arr[b], hp.arr[a] }
func (hp myheap) Less(a, b int) bool { return hp.arr[a].t < hp.arr[b].t }
func (hp *myheap) Push(x any) {
	hp.arr = append(hp.arr, x.(item))
}

func (hp *myheap) Pop() (x any) {
	old := hp.arr
	n := len(hp.arr)
	x = old[n-1]
	hp.arr = old[:n-1]
	return
}

type Cache struct {
	t  int
	hp *myheap
	mp map[int]int
	size int
}

func NewCache(size int) *Cache {
	hp := &myheap{[]item{}}
	heap.Init(hp)
	cache := &Cache{0, hp, map[int]int{}, size}
	return cache
}

func (ch *Cache) put(k, v int) {
	// _ := ch.get(k)
	if ch.hp.Len() >= ch.size {
		olditem := heap.Pop(ch.hp).(item)
		delete(ch.mp, olditem.k)
		Println("remove: ", olditem.k)
	}
	ch.mp[k] = v
	heap.Push(ch.hp, item{ch.t, k})
	Println("lru size", ch.hp.Len())
	ch.t ++
}

func (ch *Cache) get(k int) (v int) {
	value, ok := ch.mp[k]
	if ok {
		v = value
	} else {
		v = -1
	}
	return
}

// func main() {
	func main_a() {
		cache := NewCache(2)
	cache.put(1, 1)
	cache.put(2, 2)
	Println(cache.get(1))    // 返回 1
	cache.put(3, 3) // 该操作会使得关键字 2 作废
	Println(cache.get(2))    // 返回 -1 (未找到)
	cache.put(4, 4) // 该操作会使得关键字 1 作废
	Println(cache.get(1))    // 返回 -1 (未找到)
	Println(cache.get(3))    // 返回 3
	Println(cache.get(4))    // 返回 4
}
