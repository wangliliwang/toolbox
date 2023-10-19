package toolbox

import "golang.org/x/exp/constraints"

func Range[T constraints.Integer](elementNum int) []T {
	result := make([]T, 0, elementNum)
	for i := 0; i < elementNum; i++ {
		result = append(result, T(i))
	}
	return result
}

func Max[T constraints.Ordered](elements ...T) T {
	if len(elements) == 0 {
		panic("need at least one elements")
	}
	return Reduce[T, T](elements[1:], func(agg T, item T, index int) T {
		if agg < item {
			agg = item
		}
		return agg
	}, elements[0])
}

func Min[T constraints.Ordered](elements ...T) T {
	if len(elements) == 0 {
		panic("need at least one elements")
	}
	return Reduce[T, T](elements[1:], func(agg T, item T, index int) T {
		if agg > item {
			agg = item
		}
		return agg
	}, elements[0])
}
