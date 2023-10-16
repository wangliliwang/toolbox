package toolbox

func Range(elementNum int) []int {
	result := make([]int, 0, elementNum)
	for i := 0; i < elementNum; i++ {
		result = append(result, i)
	}
	return result
}
