package toolbox

func OmitBy[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) map[K]V {

}

func OmitByKeys[K comparable, V any](in map[K]V) {

}

func OmitByValues[K comparable, V any]() {

}
