package toolbox

import (
	"fmt"
	"strconv"
)

// See https://go.dev/blog/examples

func ExampleFilter() {
	list := []int64{1, 2, 3, 4, 5}
	result := Filter(list, func(item int64, index int) bool {
		return item%2 == 0
	})
	fmt.Printf("%+v", result)
	// Output: [2 4]
}

func ExampleReject() {
	list := []int64{1, 2, 3, 4, 5}
	result := Reject(list, func(item int64, index int) bool {
		return item%2 == 0
	})
	fmt.Printf("%+v", result)
	// Output: [1 3 5]
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

func ExampleReduceRight() {
	// reverse list
	list := []int64{1, 2, 3, 4, 5}
	result := ReduceRight[int64, []int64](list, func(agg []int64, item int64, index int) []int64 {
		agg = append(agg, item)
		return agg
	}, make([]int64, 0))
	fmt.Printf("%+v", result)
	// Output: [5 4 3 2 1]
}

func ExampleForEach() {
	list := []int64{1, 2, 3, 4, 5}
	ForEach(list, func(item int64, index int) {
		fmt.Printf("%d:%d,", index, item)
	})
	// Output: 0:1,1:2,2:3,3:4,4:5,
}

func ExampleForEachRight() {
	list := []int64{1, 2, 3, 4, 5}
	ForEachRight(list, func(item int64, index int) {
		fmt.Printf("%d:%d,", index, item)
	})
	// Output: 4:5,3:4,2:3,1:2,0:1,
}

func ExampleTimes() {
	result := Times(5, func(index int) int64 {
		return int64(index)
	})
	fmt.Printf("%+v", result)
	// Output: [0 1 2 3 4]
}

func ExampleUniq() {
	list := []int64{1, 1, 2, 2, 3, 3}
	result := Uniq(list)
	fmt.Printf("%+v", result)
	// Output: [1 2 3]
}

func ExampleUniqBy() {
	list := []int64{1, 2, 3, 4, 5}
	result := UniqBy(list, func(item int64, index int) int64 {
		return item % 2
	})
	fmt.Printf("%+v", result)
	// Output: [1 2]
}

func ExampleGroupBy() {
	list := []int64{1, 2, 3, 4, 5}
	result := GroupBy(list, func(item int64, index int) int64 {
		return item % 2
	})
	fmt.Printf("%+v", result)
	// Output: map[0:[2 4] 1:[1 3 5]]
}

func ExampleChunk() {
	list := []int64{1, 2, 3, 4, 5}
	result := Chunk(list, 2)
	fmt.Printf("%+v", result)
	// Output: [[1 2] [3 4] [5]]
}

func ExamplePartitionBy() {
	list := []int64{1, 2, 3, 4, 5}
	result := PartitionBy(list, func(item int64, index int) bool {
		return item%2 == 0
	})
	fmt.Printf("%+v", result)
	// Output: [[1 3 5] [2 4]]
}

func ExampleFlatten() {
	list := [][]int64{{1, 2}, {3, 4}, {5}}
	result := Flatten(list)
	fmt.Printf("%+v", result)
	// Output: [1 2 3 4 5]
}

func ExampleInterleave() {
	list := [][]int64{{1, 2}, {3, 4}, {5}, {6, 7, 8}}
	result := Interleave(list...)
	fmt.Printf("%+v", result)
	// Output: [1 3 5 6 2 4 7 8]
}

func ExampleShuffle() {
	list := []int64{1, 2, 3, 4, 5}
	result := Shuffle(list)
	fmt.Printf("%+v", len(result))
	// Output: 5
}

func ExampleReverse() {
	list := []int64{1, 2, 3, 4, 5}
	result := Reverse(list)
	fmt.Printf("%+v", result)
	// Output: [5 4 3 2 1]
}

type cloneableInt64 int64

func (c cloneableInt64) Clone() cloneableInt64 {
	return cloneableInt64(c)
}

func ExampleFill() {
	list := []int64{1, 2, 3, 4, 5}
	result := Fill(list, 0)
	fmt.Printf("%+v", result)
	// Output: [0 0 0 0 0]
}

func ExampleFillWithClone() {
	list := []cloneableInt64{1, 2, 3, 4, 5}
	result := FillWithClone[cloneableInt64](list, cloneableInt64(0))
	fmt.Printf("%+v", result)
	// Output: [0 0 0 0 0]
}

func ExampleRepeat() {
	result := Repeat(5, 5)
	fmt.Printf("%+v", result)
	// Output: [5 5 5 5 5]
}

func ExampleRepeatWithClone() {
	result := RepeatWithClone(5, cloneableInt64(5))
	fmt.Printf("%+v", result)
	// Output: [5 5 5 5 5]
}

func ExampleRepeatBy() {
	result := RepeatBy(5, func(index int) int64 {
		return int64(index)
	})
	fmt.Printf("%+v", result)
	// Output: [0 1 2 3 4]
}

func ExampleKeyBy() {
	list := []int64{1, 2, 3, 4, 5}
	result := KeyBy(list, func(item int64, index int) string {
		return fmt.Sprintf("key-%d", item)
	})
	fmt.Printf("%+v", result)
	// Output: map[key-1:1 key-2:2 key-3:3 key-4:4 key-5:5]
}

func ExampleAssociate() {
	list := []int64{1, 2, 3, 4, 5}
	result := Associate(list, func(item int64, index int) (string, string) {
		return fmt.Sprintf("key-%d", item), fmt.Sprintf("value-%d", index)
	})
	fmt.Printf("%+v", result)
	// Output: map[key-1:value-0 key-2:value-1 key-3:value-2 key-4:value-3 key-5:value-4]
}

func ExampleDrop() {
	list := []int64{1, 2, 3, 4, 5}
	result := Drop(list, 2)
	fmt.Printf("%+v", result)
	// Output: [3 4 5]
}

func ExampleDropRight() {
	list := []int64{1, 2, 3, 4, 5}
	result := DropRight(list, 2)
	fmt.Printf("%+v", result)
	// Output: [1 2 3]
}

func ExampleDropWhile() {
	list := []int64{1, 2, 3, 4, 5}
	result := DropWhile(list, func(item int64, index int) bool {
		return item < 3
	})
	fmt.Printf("%+v", result)
	// Output: [3 4 5]
}

func ExampleDropRightWhile() {
	list := []int64{1, 2, 3, 4, 5}
	result := DropRightWhile(list, func(item int64, index int) bool {
		return item > 3
	})
	fmt.Printf("%+v", result)
	// Output: [1 2 3]
}

func ExampleTake() {
	list := []int64{1, 2, 3, 4, 5}
	result := Take(list, 2)
	fmt.Printf("%+v", result)
	// Output: [1 2]
}

func ExampleTakeRight() {
	list := []int64{1, 2, 3, 4, 5}
	result := TakeRight(list, 2)
	fmt.Printf("%+v", result)
	// Output: [4 5]
}

func ExampleTakeWhile() {
	list := []int64{1, 2, 3, 4, 5}
	result := TakeWhile(list, func(item int64, index int) bool {
		return item < 3
	})
	fmt.Printf("%+v", result)
	// Output: [1 2]
}

func ExampleTakeRightWhile() {
	list := []int64{1, 2, 3, 4, 5}
	result := TakeRightWhile(list, func(item int64, index int) bool {
		return item > 3
	})
	fmt.Printf("%+v", result)
	// Output: [4 5]
}

func ExampleHead() {
	list := []int64{1, 2, 3, 4, 5}
	result, ok := Head(list)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", ok)
	// Output:
	// 1
	// true
}

func ExampleLast() {
	list := []int64{1, 2, 3, 4, 5}
	result, ok := Last(list)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", ok)
	// Output:
	// 5
	// true
}

func ExampleInitial() {
	list := []int64{1, 2, 3, 4, 5}
	result := Initial(list)
	fmt.Printf("%+v", result)
	// Output: [1 2 3 4]
}

func ExampleTail() {
	list := []int64{1, 2, 3, 4, 5}
	result := Tail(list)
	fmt.Printf("%+v", result)
	// Output: [2 3 4 5]
}

func ExampleJoin() {
	list := []int64{1, 2, 3, 4, 5}
	result := Join(list, ",")
	fmt.Printf("%+v", result)
	// Output: 1,2,3,4,5
}

func ExampleRemove() {
	list := []int64{1, 2, 3, 4, 5}
	result := Remove(list, 1, 3)
	fmt.Printf("%+v", result)
	// Output: [2 4 5]
}

func ExampleRemoveBy() {
	list := []int64{1, 2, 3, 4, 5}
	result := RemoveBy(list, func(item int64, index int) bool {
		return item%2 == 0
	})
	fmt.Printf("%+v", result)
	// Output: [1 3 5]
}

func ExampleSlice() {
	list := []int64{1, 2, 3, 4, 5}
	result := Slice(list, 1, 3)
	fmt.Printf("%+v", result)
	// Output: [2 3]
}

func ExampleSliceWithCopy() {
	list := []int64{1, 2, 3, 4, 5}
	result := SliceWithCopy(list, 1, 3)
	fmt.Printf("%+v", result)
	// Output: [2 3]
}

func ExampleSampleSize() {
	list := []int64{1, 2, 3, 4, 5, 6, 7}
	result := SampleSize(list, 2)
	fmt.Printf("%+v", len(result))
	// Output: 2
}

func ExampleNth() {
	collection := []int{1, 2, 3, 4, 5}
	result1, err1 := Nth(collection, 1)
	result2, err2 := Nth(collection, -1)
	fmt.Println(result1, err1, result2, err2)
	// Output:
	// 2 <nil> 5 <nil>
}
