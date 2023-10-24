package toolbox

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"sort"
	"strings"
)

var exampleMap = map[string]int{
	"apple":  10,
	"orange": 12,
	"banana": 15,
}

type entries[K constraints.Ordered, V any] []Entry[K, V]

func (e entries[K, V]) Len() int {
	return len(e)
}

func (e entries[K, V]) Less(i, j int) bool {
	return e[i].Key < e[j].Key
}

func (e entries[K, V]) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func sprintSortedMapByKey[K constraints.Ordered, V any](in map[K]V) string {
	// transform to entry
	es := ToEntries(in)

	// sort
	sort.Sort(entries[K, V](es))

	// format
	result := "map["
	itemStrs := make([]string, 0, len(es))
	for _, item := range es {
		itemStrs = append(itemStrs, fmt.Sprintf("%+v:%+v", item.Key, item.Value))
	}
	result += strings.Join(itemStrs, " ") + "]"
	return result
}

func ExampleOmitBy() {
	result := OmitBy(exampleMap, func(key string, value int) bool {
		return strings.HasPrefix(key, "or")
	})
	fmt.Printf("%+v ", sprintSortedMapByKey(result))
	// Output:
	// map[apple:10 banana:15]
}

func ExampleOmitByKeys() {
	result := OmitByKeys(exampleMap, []string{"apple"})
	fmt.Printf("%+v", sprintSortedMapByKey(result))
	// Output:
	// map[banana:15 orange:12]
}

func ExampleOmitByValues() {
	result := OmitByValues(exampleMap, []int{10})
	fmt.Printf("%+v", sprintSortedMapByKey(result))
	// Output:
	// map[banana:15 orange:12]
}

func ExamplePickBy() {
	result := PickBy(exampleMap, func(key string, value int) bool {
		return strings.HasPrefix(key, "or")
	})
	fmt.Printf("%+v", sprintSortedMapByKey(result))
	// Output:
	// map[orange:12]
}

func ExamplePickByKeys() {
	result := PickByKeys(exampleMap, []string{"apple"})
	fmt.Printf("%+v", sprintSortedMapByKey(result))
	// Output:
	// map[apple:10]
}

func ExamplePickByValues() {
	result := PickByValues(exampleMap, []int{10})
	fmt.Printf("%+v", sprintSortedMapByKey(result))
	// Output:
	// map[apple:10]
}

func ExampleToEntries() {
	result := ToEntries(exampleMap)
	sort.Sort(entries[string, int](result))
	fmt.Printf("%+v", result)
	// Output:
	// [{Key:apple Value:10} {Key:banana Value:15} {Key:orange Value:12}]
}

func ExampleFromEntries() {
	es := []Entry[string, int]{
		{"wang", 1},
		{"zhao", 3},
	}
	result := FromEntries(es)
	fmt.Printf("%+v", sprintSortedMapByKey(result))
	// Output:
	// map[wang:1 zhao:3]
}

func ExampleInvert() {
	result := Invert(exampleMap)
	fmt.Printf("%+v", sprintSortedMapByKey(result))
	// Output:
	// map[10:apple 12:orange 15:banana]
}

func ExampleMapOverMap() {
	MapOverMap(exampleMap, func(k string, v int) {
		fmt.Println(k, v)
	})
}

func ExampleMapKeys() {
	result := MapKeys(exampleMap, func(k string, v int) string {
		return "common_key"
	})
	fmt.Printf("%+v", Keys(result))
	// Output:
	// [common_key]
}

func ExampleMapValues() {
	result := MapValues(exampleMap, func(k string, v int) int {
		return 0
	})
	fmt.Printf("%+v", Values(result))
	// Output:
	// [0 0 0]
}

func ExampleMapEntries() {
	result := MapEntries(exampleMap, func(k string, v int) (string, int) {
		return "common_key", 0
	})
	fmt.Printf("%+v", result)
	// Output:
	// map[common_key:0]
}

func ExampleMerge() {
	result := Merge(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"c": 3},
	)
	fmt.Printf("%+v", sprintSortedMapByKey(result))
	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleKeys() {
	result := Keys(exampleMap)
	fmt.Printf("%+v", len(result))
	// Output:
	// 3
}

func ExampleValues() {
	result := Values(exampleMap)
	fmt.Printf("%+v", len(result))
	// Output:
	// 3
}

func ExampleValueOr() {
	result1 := ValueOr(exampleMap, "apple", 11)
	result2 := ValueOr(exampleMap, "bear", 11)
	fmt.Printf("%+v %+v", result1, result2)
	// Output:
	// 10 11
}

func ExampleMapToSlice() {
	result := MapToSlice(exampleMap, func(k string, v int) int64 {
		return 1
	})
	fmt.Printf("%+v", result)
	// Output:
	// [1 1 1]
}
