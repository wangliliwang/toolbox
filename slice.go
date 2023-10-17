package toolbox

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
