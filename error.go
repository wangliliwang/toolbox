package toolbox

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
