package main

import "fmt"

// 在这个示例中，我们定义了一个 HashSet 类型，它实际上是一个 map[string]struct{} 类型的别名。
// 我们定义了 Add、Remove 和 Contains 方法来操作哈希集合。添加元素时，我们将元素作为 map 的键，对应的值设为一个空的 struct{} 结构体。
// 这样就可以模拟一个只包含键而不包含值的集合。
// 在 Go 中，将 value 定义为 struct{} 和 interface{} 在语义上有一些区别。

// struct{}： 表示一个空的结构体，不包含任何字段。因为 struct{} 不包含任何数据，所以将其用作哈希集合的 value 类型时，只关注键而不关注值的情况。
// 这种方式在内存使用上更加高效，因为不会占用额外的内存空间。

// interface{}： 表示一个空接口，可以存储任意类型的值。将 value 定义为 interface{} 时，可以存储任何类型的数据，
// 但这也意味着会占用更多的内存空间，并且可能会带来类型转换的开销。

// 因此，如果在哈希集合中只需要关注键而不需要关注值的情况下，通常会选择使用 struct{} 来定义 value 类型，这样可以节省内存并提高性能。
// 而如果需要存储任意类型的值，则可以使用 interface{}。
type HashSet map[string]struct{}

func (set HashSet) Add(value string) {
	set[value] = struct{}{}
}

func (set HashSet) Remove(value string) {
	delete(set, value)
}

func (set HashSet) Contains(value string) bool {
	_, exists := set[value]
	return exists
}

func main() {
	mySet := make(HashSet)

	mySet.Add("apple")
	mySet.Add("banana")
	mySet.Add("orange")

	fmt.Println(mySet.Contains("apple")) // true
	fmt.Println(mySet.Contains("grape")) // false

	mySet.Remove("banana")
	fmt.Println(mySet.Contains("banana")) // false
}
