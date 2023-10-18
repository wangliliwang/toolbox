package toolbox

import "golang.org/x/exp/constraints"

func Range[T constraints.Integer](elementNum int) []T {
	result := make([]T, 0, elementNum)
	for i := 0; i < elementNum; i++ {
		result = append(result, T(i))
	}
	return result
}
