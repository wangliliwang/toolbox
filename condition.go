package toolbox

// Ternary is a one line if-else statement
func Ternary[T any](condition bool, ifOutput, elseOutput T) T {
	if condition {
		return ifOutput
	}
	return elseOutput
}

// TernaryF is a one line if-else statement whose options are functions.
func TernaryF[T any](condition bool, ifFunc, elseFunc func() T) T {
	if condition {
		return ifFunc()
	}
	return elseFunc()
}

type ifElse[T any] struct {
	done   bool
	result T
}

func (i *ifElse[T]) Else(t T) T {
	if !i.done {
		return t
	}
	return i.result
}

func (i *ifElse[T]) ElseF(predicate func() T) T {
	if !i.done {
		return predicate()
	}
	return i.result
}

func (i *ifElse[T]) ElseIf(condition bool, t T) *ifElse[T] {
	if !i.done {
		if condition {
			return &ifElse[T]{
				done:   true,
				result: t,
			}
		}
		var zero T
		return &ifElse[T]{
			done:   false,
			result: zero,
		}
	}
	return i
}

func (i *ifElse[T]) ElseIfF(condition bool, predicate func() T) *ifElse[T] {
	if !i.done {
		if condition {
			return &ifElse[T]{
				done:   true,
				result: predicate(),
			}
		}
		var zero T
		return &ifElse[T]{
			done:   false,
			result: zero,
		}
	}
	return i
}

func If[T any](condition bool, t T) *ifElse[T] {
	if condition {
		return &ifElse[T]{
			done:   true,
			result: t,
		}
	}
	var zero T
	return &ifElse[T]{
		done:   false,
		result: zero,
	}
}

func IfF[T any](condition bool, predicate func() T) *ifElse[T] {
	if condition {
		return &ifElse[T]{
			done:   true,
			result: predicate(),
		}
	}
	var zero T
	return &ifElse[T]{
		done:   false,
		result: zero,
	}
}

type switchCase[T any] struct {
	predicate T
	done      bool
	result    T
}

func (s *switchCase[T]) Case(condition bool, t T) *switchCase[T] {
	if s.done {
		return s
	}
	if condition {
		return &switchCase[T]{
			predicate: s.predicate,
			done:      true,
			result:    t,
		}
	}
	var zero T
	return &switchCase[T]{
		predicate: s.predicate,
		done:      false,
		result:    zero,
	}
}

func (s *switchCase[T]) CaseF(condition bool, t func(predicate T) T) *switchCase[T] {
	if s.done {
		return s
	}
	if condition {
		return &switchCase[T]{
			predicate: s.predicate,
			done:      true,
			result:    t(s.predicate),
		}
	}
	var zero T
	return &switchCase[T]{
		predicate: s.predicate,
		done:      false,
		result:    zero,
	}
}

func (s *switchCase[T]) Default(t T) T {
	if s.done {
		return s.result
	}
	return t
}

func (s *switchCase[T]) DefaultF(t func(predicate T) T) T {
	if s.done {
		return s.result
	}
	return t(s.predicate)
}

func Switch[T any](predicate T) *switchCase[T] {
	var zero T
	return &switchCase[T]{
		predicate: predicate,
		done:      false,
		result:    zero,
	}
}
