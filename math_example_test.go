package toolbox

import "fmt"

func ExampleRange() {
	result1 := Range[int64](3)
	result2 := Range[int64](-3)
	fmt.Println(result1)
	fmt.Println(result2)
	// Output:
	// [0 1 2]
	// [0 -1 -2]
}

func ExampleRangeFrom() {
	result1 := RangeFrom[int64](8, 3)
	result2 := RangeFrom[int64](8, -3)
	fmt.Println(result1)
	fmt.Println(result2)
	// Output:
	// [8 9 10]
	// [8 7 6]
}

func ExampleRangeWithStep() {
	result1 := RangeWithStep[int64](8, 2, 3)
	result2 := RangeWithStep[int64](8, -2, 3)
	fmt.Println(result1)
	fmt.Println(result2)
	// Output:
	// [8 10 12]
	// [8 6 4]
}

func ExampleMax() {
	result := Max(1, 2, 3, 6, 5)
	fmt.Println(result)
	// Output:
	// 6
}

func ExampleMin() {
	result := Min(7, -1, 3, 6, 7)
	fmt.Println(result)
	// Output:
	// -1
}

func ExampleSum() {
	result := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(result)
	// Output:
	// 15
}

func ExampleSumBy() {
	result := SumBy([]int{1, 2, 3, 4, 5}, func(t, index int) int {
		return t + 1
	})
	fmt.Println(result)
	// Output:
	// 20
}
