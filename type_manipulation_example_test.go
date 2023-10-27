package toolbox

import "fmt"

func ExampleToPtr() {
	x := 1
	result := ToPtr(x)
	fmt.Println(*result)
	// Output:
	// 1
}

func ExampleFromPtr() {
	x := 1
	y := &x
	result := FromPtr(y)
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleFromPtrWithFallback() {
	var x *int = nil
	result := FromPtrWithFallback(x, 1)
	fmt.Println(result)
	// Output:
	// 1
}

func ExampleToSlicePtr() {
	collection := []int{1, 2, 3, 4, 5}
	ToSlicePtr(collection)
}

func ExampleFromSlicePtr() {
	collection := []*int{ToPtr(1), ToPtr(2), ToPtr(3), ToPtr(4), ToPtr(5)}
	result := FromSlicePtr(collection)
	fmt.Println(result)
	// Output:
	// [1 2 3 4 5]
}

func ExampleFromSlicePtrWithFallback() {
	var nilValue *int = nil
	collection := []*int{ToPtr(1), ToPtr(2), nilValue, ToPtr(4), ToPtr(5)}
	result := FromSlicePtrWithFallback(collection, 7)
	fmt.Println(result)
	// Output:
	// [1 2 7 4 5]
}

func ExampleToAnySlice() {
	collection := []int{1, 2, 3, 4, 5}
	result := ToAnySlice(collection)
	fmt.Println(result)
	// Output:
	// [1 2 3 4 5]
}

func ExampleFromAnySlice() {
	collection := []any{1, 2, "3", 4, 5}
	result, ok := FromAnySlice[int](collection)
	fmt.Println(result, ok)
	// Output:
	// [1 2 0 4 5] false
}

func ExampleZero() {
	result := Zero[int]()
	fmt.Println(result)
	// Output:
	// 0
}

func ExampleIsZero() {
	result := IsZero[int](1)
	fmt.Println(result)
	// Output:
	// false
}

func ExampleIsNotZero() {
	result := IsNotZero[int](1)
	fmt.Println(result)
	// Output:
	// true
}
