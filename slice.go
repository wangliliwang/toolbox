package toolbox

import (
	"math"
	"math/rand"
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

// Times invokes the iteratee function n times, returning an array of results of each invocation.
func Times[T any](count int, iteratee func(index int) T) []T {
	result := make([]T, count)
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
func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	result := make([]T, 0)
	seen := make(map[U]struct{})
	for _, item := range collection {
		u := iteratee(item)
		if _, ok := seen[u]; !ok {
			seen[u] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	result := make(map[U][]T)
	for _, item := range collection {
		u := iteratee(item)
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
func PartitionBy[T any, P comparable](collection []T, iteratee func(item T) P) [][]T {
	result := make([][]T, 0)
	seen := make(map[P]int)
	for _, item := range collection {
		key := iteratee(item)
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
			if len(collections)-1 >= lengthIndex {
				result = append(result, collection[lengthIndex])
			}
		}
	}
	return result
}

// Shuffle shuffle the values in the collection in-place.
// Using the Fisher-Yates shuffle algorithm.
// todo(wangli) learn this algo
func Shuffle[T any](collection []T) {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})
}

// Reverse reverses array in-place so that the first element become the last, the second element becomes the second to the last, and so on.
func Reverse[T any](collection []T) {
	length := len(collection)
	half := length / 2
	for i := 0; i < half; i++ {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}
}

// Fill fills the collection with initial value.
func Fill[T Clonable[T]](collection []T, initial T) {
	for index := range collection {
		collection[index] = initial.Clone()
	}
}

// Repeat returns a array of count initial values.
func Repeat[T Clonable[T]](count int, initial T) []T {
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
func KeyBy[K comparable, V any](collection []V, iteratee func(v V) K) map[K]V {
	result := make(map[K]V)
	for _, item := range collection {
		result[iteratee(item)] = item
	}
	return result
}

// Associate transforms a slice to a map whose key-value pairs are generated by transform function.
func Associate[T any, K comparable, V any](collection []T, transform func(t T) (K, V)) map[K]V {
	result := make(map[K]V)
	for _, item := range collection {
		k, v := transform(item)
		result[k] = v
	}
	return result
}
