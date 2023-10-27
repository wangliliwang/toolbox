package toolbox

// ToPtr returns a pointer copy of value.
func ToPtr[T any](x T) *T {
	return &x
}

// FromPtr returns the pointer value, or zero if pointer is nil.
func FromPtr[T any](x *T) T {
	if x == nil {
		return Zero[T]()
	}
	return *x
}

// FromPtrWithFallback returns the pointer value, or fallback if pointer is nil.
func FromPtrWithFallback[T any](x *T, fallback T) T {
	if x == nil {
		return fallback
	}
	return *x
}

// ToSlicePtr returns a slice of elements that are pointer copies of collection.
func ToSlicePtr[T any](collection []T) []*T {
	return Map[T, *T](collection, func(item T, index int) *T {
		return ToPtr(item)
	})
}

// FromSlicePtr returns a slice of elements that are values of collection.
func FromSlicePtr[T any](collection []*T) []T {
	return Map[*T, T](collection, func(item *T, index int) T {
		return FromPtr(item)
	})
}

// FromSlicePtrWithFallback returns a slice of elements that are values with fallback of collection.
func FromSlicePtrWithFallback[T any](collection []*T, fallback T) []T {
	return Map[*T, T](collection, func(item *T, index int) T {
		return FromPtrWithFallback(item, fallback)
	})
}

// ToAnySlice returns a slice of elements whose types are any, from collection.
func ToAnySlice[T any](collection []T) []any {
	return Map[T, any](collection, func(item T, index int) any {
		return item
	})
}

// FromAnySlice returns a slice of elements whose types are T, from collection.
func FromAnySlice[T any](collection []any) ([]T, bool) {
	ok := true
	result := Map[any, T](collection, func(item any, index int) T {
		if value, innerOk := item.(T); innerOk {
			return value
		} else {
			ok = false
			return Zero[T]()
		}
	})
	return result, ok
}

// Zero returns zero value of T.
func Zero[T any]() T {
	var zero T
	return zero
}

// IsZero returns true if x equals to zero value of T.
func IsZero[T comparable](x T) bool {
	return x == Zero[T]()
}

// IsNotZero returns true if x not equals to zero value of T.
func IsNotZero[T comparable](x T) bool {
	return !IsZero(x)
}
