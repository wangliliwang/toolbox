package toolbox

import (
	"fmt"
	"github.com/spf13/cast"
	"math"
	"math/rand"
	"strings"
)

// Filter iterates over elements of collection, returns an array of all elements that predicate function returns truthy for.
func Filter[T any](collection []T, predicate func(item T, index int) bool) []T {
	result := make([]T, 0)
	for index, item := range collection {
		if predicate(item, index) {
			result = append(result, item)
		}
	}
	return result
}

// Reject iterates over elements of collection, returns an array of all elements that predicate function returns falsey for.
func Reject[T any](collection []T, predicate func(item T, index int) bool) []T {
	result := make([]T, 0)
	for index, item := range collection {
		if !predicate(item, index) {
			result = append(result, item)
		}
	}
	return result
}

// Map manipulates a slice and transforms it into a slice with another type.
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, 0, len(collection))
	for index, item := range collection {
		result = append(result, iteratee(item, index))
	}
	return result
}

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMap[T any, R any](collection []T, callback func(item T, index int) (R, bool)) []R {
	result := make([]R, 0, len(collection))
	for index, item := range collection {
		if value, ok := callback(item, index); ok {
			result = append(result, value)
		}
	}
	return result
}

// FlatMap manipulates a slice and transforms and flattens it to a slice of another type.
// The transform function either return a slice or a `nil`.
func FlatMap[T any, R any](collection []T, iteratee func(item T, index int) []R) []R {
	result := make([]R, 0)
	for index, item := range collection {
		result = append(result, iteratee(item, index)...)
	}
	return result
}

// Reduce reduces a collection into a value which is the accumulate result of running each element in collection through accumulator,
// where each successive invocation is supplied the return value of the previous.
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for index, item := range collection {
		initial = accumulator(initial, item, index)
	}
	return initial
}

// ReduceRight helper is like Reduce except that it iterates over elements from right to left.
func ReduceRight[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for index := len(collection) - 1; index >= 0; index-- {
		initial = accumulator(initial, collection[index], index)
	}
	return initial
}

// ForEach iterates over collection and invokes iteratee function for each element.
func ForEach[T any](collection []T, iteratee func(item T, index int)) {
	for index, item := range collection {
		iteratee(item, index)
	}
}

// ForEachRight iterates over collection from the end and invokes iteratee function for each element.
func ForEachRight[T any](collection []T, iteratee func(item T, index int)) {
	for index := len(collection) - 1; index >= 0; index-- {
		iteratee(collection[index], index)
	}
}

// Times invokes the iteratee function n times, returning an array of results of each invocation.
func Times[T any](count int, iteratee func(index int) T) []T {
	result := make([]T, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, iteratee(i))
	}
	return result
}

// Uniq returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
func Uniq[T comparable](collection []T) []T {
	result := make([]T, 0)
	seen := make(map[T]struct{})
	for _, item := range collection {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// UniqBy returns a duplicate-free version of an array, in which only the first occurrence of the element is kept.
// It accepts `iteratee` which is invoked for each element in array to generate the criterion by which uniqueness is computed.
// The order of result values is determined by the order they occur in the array.
func UniqBy[T any, U comparable](collection []T, iteratee func(item T, index int) U) []T {
	result := make([]T, 0)
	seen := make(map[U]struct{})
	for index, item := range collection {
		u := iteratee(item, index)
		if _, ok := seen[u]; !ok {
			seen[u] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
func GroupBy[T any, U comparable](collection []T, iteratee func(item T, index int) U) map[U][]T {
	result := make(map[U][]T)
	for index, item := range collection {
		u := iteratee(item, index)
		result[u] = append(result[u], item)
	}
	return result
}

// Chunk returns an array of elements split into groups the length of size.
// If collection can't be split evenly, the final chunk will be the remaining elements.
func Chunk[T any](collection []T, size int) [][]T {
	if size <= 0 {
		panic("size should be greater than 0")
	}
	chunkNum := int(math.Ceil(float64(len(collection)) / float64(size)))
	result := make([][]T, chunkNum)
	for chunkIndex := 0; chunkIndex < chunkNum; chunkIndex++ {
		endIndex := (chunkIndex + 1) * size
		if endIndex > len(collection) {
			endIndex = len(collection)
		}
		result[chunkIndex] = collection[chunkIndex*size : endIndex]
	}
	return result
}

// PartitionBy returns an array of elements split into groups.
// The order of grouped values is determined by the order they occur in collection.
// The grouping is generated from the results of running each element of collection through iteratee function.
func PartitionBy[T any, P comparable](collection []T, iteratee func(item T, index int) P) [][]T {
	result := make([][]T, 0)
	seen := make(map[P]int)
	for index, item := range collection {
		key := iteratee(item, index)
		if index, ok := seen[key]; ok {
			result[index] = append(result[index], item)
		} else {
			seen[key] = len(result)
			result = append(result, []T{item})
		}
	}
	return result
}

// Flatten returns an array of single level deep.
func Flatten[T any](collection [][]T) []T {
	result := make([]T, 0)
	for _, item := range collection {
		result = append(result, item...)
	}
	return result
}

// Interleave round-robin alternating input slices and sequentially appending value at index into result.
func Interleave[T any](collections ...[]T) []T {
	result := make([]T, 0)
	maxLengthCollection := FindBy[[]T](collections, func(a, b []T) bool {
		return len(a) > len(b)
	})
	maxLength := len(maxLengthCollection)
	for lengthIndex := 0; lengthIndex < maxLength; lengthIndex++ {
		for _, collection := range collections {
			if len(collection)-1 >= lengthIndex {
				result = append(result, collection[lengthIndex])
			}
		}
	}
	return result
}

// Shuffle shuffles the values in the collection in-place.
// Using the Fisher-Yates shuffle algorithm.
func Shuffle[T any](collection []T) []T {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})
	return collection
}

// Reverse reverses array in-place so that the first element become the last, the second element becomes the second to the last, and so on.
func Reverse[T any](collection []T) []T {
	length := len(collection)
	half := length / 2
	for i := 0; i < half; i++ {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}
	return collection
}

// Fill fills the collection with initial value.
func Fill[T any](collection []T, initial T) []T {
	for index := range collection {
		collection[index] = initial
	}
	return collection
}

// FillWithClone fills the collection with cloned initial value.
func FillWithClone[T Cloneable[T]](collection []T, initial T) []T {
	for index := range collection {
		collection[index] = initial
	}
	return collection
}

// Repeat returns a array of count initial values.
func Repeat[T any](count int, initial T) []T {
	result := make([]T, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, initial)
	}
	return result
}

// RepeatWithClone returns an array of count initial cloned values.
func RepeatWithClone[T Cloneable[T]](count int, initial T) []T {
	result := make([]T, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, initial.Clone())
	}
	return result
}

// RepeatBy returns an array of count values that generated by predicate function.
func RepeatBy[T any](count int, predicate func(index int) T) []T {
	result := make([]T, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, predicate(i))
	}
	return result
}

// KeyBy transforms a slice to a map based on a pivot callback.
func KeyBy[K comparable, V any](collection []V, iteratee func(v V, index int) K) map[K]V {
	result := make(map[K]V)
	for index, item := range collection {
		result[iteratee(item, index)] = item
	}
	return result
}

// Associate transforms a slice to a map whose key-value pairs are generated by transform function.
func Associate[T any, K comparable, V any](collection []T, transform func(t T, index int) (K, V)) map[K]V {
	result := make(map[K]V)
	for index, item := range collection {
		k, v := transform(item, index)
		result[k] = v
	}
	return result
}

// Drop returns a slice with n elements dropped from the beginning of the collection.
func Drop[T any](collection []T, n int) []T {
	if n >= len(collection) {
		return make([]T, 0)
	}
	result := make([]T, 0, len(collection)-n)
	result = append(result, collection[n:]...)
	return result
}

// DropRight returns a slice with n elements dropped from the end of the collection.
func DropRight[T any](collection []T, n int) []T {
	if n >= len(collection) {
		return make([]T, 0)
	}
	result := make([]T, 0, len(collection)-n)
	result = append(result, collection[:len(collection)-n]...)
	return result
}

// DropWhile returns a slice excluding elements dropped from the beginning.
// Elements are dropped until the predicate function returns falsey.
func DropWhile[T any](collection []T, predicate func(item T, index int) bool) []T {
	startIndex := 0
	for ; startIndex < len(collection); startIndex++ {
		if !predicate(collection[startIndex], startIndex) {
			break
		}
	}
	return Drop(collection, startIndex)
}

// DropRightWhile returns a slice excluding elements dropped from the end.
// Elements are dropped until the predicate function returns falsey.
func DropRightWhile[T any](collection []T, predicate func(item T, index int) bool) []T {
	endIndex := len(collection) - 1
	for ; endIndex >= 0; endIndex-- {
		if !predicate(collection[endIndex], endIndex) {
			break
		}
	}
	return DropRight(collection, len(collection)-1-endIndex)
}

// Take creates a slice of n elements taken from the beginning.
func Take[T any](collection []T, n int) []T {
	return Slice(collection, 0, n)
}

// TakeRight creates a slice of n elements taken from the end.
func TakeRight[T any](collection []T, n int) []T {
	return Slice(collection, len(collection)-n, len(collection))
}

// TakeWhile creates a slice of n elements taken from the beginning.
// Elements are taken until the predicate function returns falsey.
func TakeWhile[T any](collection []T, predicate func(item T, index int) bool) []T {
	n := 0
	for index, item := range collection {
		if !predicate(item, index) {
			n = index
			break
		}
	}
	return Take(collection, n)
}

// TakeRightWhile creates a slice of n elements taken from the end.
// Elements are taken until the predicate function returns falsey.
func TakeRightWhile[T any](collection []T, predicate func(item T, index int) bool) []T {
	endIndex := len(collection) - 1
	for ; endIndex >= 0; endIndex-- {
		if !predicate(collection[endIndex], endIndex) {
			break
		}
	}
	return TakeRight(collection, len(collection)-1-endIndex)
}

// Head returns first element in the collection.
// If the collection is empty, return zero value of T and false.
func Head[T any](collection []T) (T, bool) {
	if len(collection) > 0 {
		return collection[0], true
	}
	var zero T
	return zero, false
}

// Last returns last element in the collection.
// If the collection is empty, return zero value of T and false.
func Last[T any](collection []T) (T, bool) {
	if len(collection) > 0 {
		return collection[len(collection)-1], true
	}
	var zero T
	return zero, false
}

// Initial returns all but the last element of the collection.
func Initial[T any](collection []T) []T {
	return DropRight(collection, 1)
}

// Tail returns all but the first element of the collection.
func Tail[T any](collection []T) []T {
	return Drop(collection, 1)
}

// Join converts all elements in array into a string separated by separator
// TODO(@wangli) thinking about how to check if T is string
func Join[T any](collection []T, separator string) string {
	return strings.Join(Map[T, string](collection, func(item T, index int) string {
		// efficient any to string
		result, err := cast.ToStringE(item)
		if err != nil {
			return fmt.Sprintf("%v", item)
		}
		return result
	}), separator)
}

// Remove excludes all excludedValues in collection.
func Remove[T comparable](collection []T, excludedValues ...T) []T {
	excludeMap := CollectionToSet(excludedValues)
	return Filter[T](collection, func(item T, index int) bool {
		_, ok := excludeMap[item]
		return !ok
	})
}

// RemoveBy removes all elements that predicate function returns true in collection.
func RemoveBy[T any](collection []T, predicate func(item T, index int) bool) []T {
	return Filter[T](collection, func(item T, index int) bool {
		return !predicate(item, index)
	})
}

// Slice returns a slice of collection from start up to, but not including, end.
// It's like collection[start:end], but will not panic on overflow.
func Slice[T any](collection []T, start, end int) []T {
	length := len(collection)
	if start >= end || length == 0 {
		return make([]T, 0)
	}
	if start < 0 {
		start = 0
	}
	if start > length-1 {
		start = length - 1
	}
	if end < 0 {
		end = 0
	}
	if end > length {
		end = length
	}
	return collection[start:end]
}

// SliceWithCopy returns a copy of slice in collection from start up to, but not including, end.
func SliceWithCopy[T any](collection []T, start, end int) []T {
	src := Slice(collection, start, end)
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

// Count returns the number of times of the value.
func Count[T comparable](collection []T, value T) int {
	count := 0
	for _, item := range collection {
		if item == value {
			count++
		}
	}
	return count
}

// CountBy returns a map composed of keys generated from iteratee.
// The corresponding value of each key is the number of times the key was returned by iteratee.
func CountBy[T any, K comparable](collection []T, iteratee func(item T, index int) K) map[K]int {
	result := make(map[K]int)
	for index, item := range collection {
		k := iteratee(item, index)
		result[k]++
	}
	return result
}

// Every returns true if all elements of a subset are contained into a collection or if the subset is empty.
func Every[T comparable](collection []T, subset []T) bool {
	if len(subset) <= 1 {
		for _, item := range subset {
			if !Contains(collection, item) {
				return false
			}
		}
		return true
	} else {
		mp := CollectionToSet(collection)
		for _, item := range subset {
			if !containsInSet(mp, item) {
				return false
			}
		}
		return true
	}
}

// EveryBy returns true if predicate function returns true for all elements in the collection or if the collection is empty.
func EveryBy[T comparable](collection []T, predicate func(item T) bool) bool {
	for _, item := range collection {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Contains returns true if an element present in a collection.
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}

func containsInSet[T comparable](set map[T]struct{}, element T) bool {
	_, ok := set[element]
	return ok
}

// ContainsBy returns true if predicate function return true.
func ContainsBy[T comparable](collection []T, predicate func(item T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Some returns true if at lease one element of a subset is contained into a collection.
// If the subset is empty Some returns false.
func Some[T comparable](collection []T, subset []T) bool {
	if len(subset) <= 1 {
		for _, item := range subset {
			if Contains(collection, item) {
				return true
			}
		}
		return false
	} else {
		mp := CollectionToSet(collection)
		for _, item := range subset {
			if containsInSet(mp, item) {
				return true
			}
		}
		return false
	}
}

// SomeBy returns true if predicate function returns true for at least one element in the collection.
// If the subset is empty SomeBy returns false.
func SomeBy[T comparable](collection []T, predicate func(item T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
}

// None returns true if all elements of a subset is NOT contained into a collection or if the subset is empty.
func None[T comparable](collection []T, subset []T) bool {
	if len(subset) <= 1 {
		for _, item := range subset {
			if Contains(collection, item) {
				return false
			}
		}
		return true
	} else {
		mp := CollectionToSet(collection)
		for _, item := range subset {
			if containsInSet(mp, item) {
				return false
			}
		}
		return true
	}
}

// NoneBy returns true if predicate function returns false for all elements in the collection or if the subset is empty.
func NoneBy[T comparable](collection []T, predicate func(item T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return false
		}
	}
	return true
}

// Sample returns a random element in collection.
// If collection is empty, return zero value and false.
func Sample[T any](collection []T) (T, bool) {
	if len(collection) == 0 {
		var zero T
		return zero, false
	}
	// perm
	return collection[rand.Intn(len(collection))], true
}

// SampleSize returns n random elements with diffent indexes in collection.
// TODO(@wangli) make it more efficient.
func SampleSize[T any](collection []T, n int) []T {
	min := Min(len(collection), n)
	indexes := rand.Perm(len(collection))
	return RepeatBy(min, func(index int) T {
		return collection[indexes[index]]
	})
}
