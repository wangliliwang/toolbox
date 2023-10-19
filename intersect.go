package toolbox

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
