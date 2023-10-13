package toolbox

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

// Intersect returns intersection of list1 and list2
func Intersect[T comparable](list1 []T, list2 []T) []T {
	result := make([]T, 0)
	seen := CollectionToSet(list1)
	for _, item := range list2 {
		if _, ok := seen[item]; ok {
			result = append(result, item)
		}
	}
	return result
}

// Subtract returns subtraction of list1 and list2, that is result of list1 - list2.
func Subtract[T comparable](list1 []T, list2 []T) []T {
	result := make([]T, 0)
	seen2 := CollectionToSet(list2)
	for _, item := range list1 {
		if _, ok := seen2[item]; !ok {
			result = append(result, item)
		}
	}
	return result
}

// Union returns union of list.
func Union[T comparable](lists ...[]T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)
	for _, list := range lists {
		for _, item := range list {
			if _, ok := seen[item]; !ok {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

func Without[T comparable](collection []T, exclude ...T) []T {
	result := make([]T, 0, len(collection))
	excludeMp := CollectionToSet(exclude)
	for _, item := range collection {
		if !containsInSet(excludeMp, item) {
			result = append(result, item)
		}
	}
	return result
}

func WithoutEmpty[T comparable](collection []T) []T {
	result := make([]T, 0, len(collection))
	var empty T
	for _, item := range collection {
		if empty != item {
			result = append(result, item)
		}
	}
	return result
}
