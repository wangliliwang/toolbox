package toolbox

type Cloneable[T any] interface {
	Clone() T
}
