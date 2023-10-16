package toolbox

// FindBy returns minimum element in collection by comparison function.
// If there are multi minimum values, MinBy returns the first.
// If collection is empty, MinBy returns zero value of T.
// todo(wangli) Change name and description.
func FindBy[T any](collection []T, comparison func(a, b T) bool) T {
	var min T
	if len(collection) == 0 {
		return min
	}
	min = collection[0]
	for i := 1; i < len(collection); i++ {
		if comparison(collection[i], min) {
			min = collection[i]
		}
	}
	return min
}
