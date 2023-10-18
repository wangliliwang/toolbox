package toolbox

import (
	"fmt"
	"strconv"
)

func ExampleFilter() {
	list := []int64{1, 2, 3, 4, 5}
	result := Filter(list, func(item int64, index int) bool {
		return item%2 == 0
	})
	fmt.Printf("%+v", result)
	// Output: [2 4]
}

func ExampleMap() {
	list := []int64{1, 2, 3}
	result := Map(list, func(item int64, index int) string {
		return strconv.FormatInt(item, 10)
	})
	fmt.Printf("%+v", result)
	// Output: [1 2 3]
}

func ExampleFilterMap() {
	list := []int64{1, 2, 3, 4, 5}
	result := FilterMap(list, func(item int64, index int) (string, bool) {
		return strconv.FormatInt(item, 10), item%2 == 0
	})
	fmt.Printf("%+v", result)
	// Output: [2 4]
}

func ExampleFlatMap() {
	list := []int64{1, 23}
	result := FlatMap(list, func(item int64, index int) []string {
		return []string{
			strconv.FormatInt(item, 10),
			strconv.FormatInt(item, 26),
		}
	})
	fmt.Printf("%+v", result)
	// Output: [1 1 23 n]
}

func ExampleReduce() {
	list := Range[int64](101)
	result := Reduce[int64, int64](list, func(agg int64, item int64, index int) int64 {
		return agg + item
	}, 0)
	fmt.Printf("%+v", result)
	// Output: 5050
}
