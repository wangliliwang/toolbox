package toolbox

type Clonable[T any] interface {
	Clone() T
}
