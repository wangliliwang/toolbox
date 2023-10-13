package toolbox

func CollectionToSet[T comparable](collection []T) map[T]struct{} {
	mp := make(map[T]struct{})
	for _, item := range collection {
		mp[item] = struct{}{}
	}
	return mp
}
