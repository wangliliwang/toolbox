package toolbox

import "fmt"

func ExampleRandomString() {
	result := RandomString(10, AllCharset)
	fmt.Println(len(result))
	// Output:
	// 10
}

func ExampleSubString() {
	result1 := SubString("abcde", -2, 2)
	result2 := SubString("abcde", 0, 10)
	fmt.Println(result1, result2)
	// Output:
	// de abcde
}

func ExampleChunkString() {
	result := ChunkString("abcde", 2)
	fmt.Println(result)
	// Output:
	// [ab cd e]
}

func ExampleRuneLength() {
	result := RuneLength("王hha")
	fmt.Println(result)
	// Output:
	// 4
}

func ExampleSnakeCase() {
	result := SnakeCase("hello, world!")
	fmt.Println(result)
	// Output:
	// hello_world
}

func ExampleCamelCase() {
	result := CamelCase("hello, world!")
	fmt.Println(result)
	// Output:
	// helloWorld
}

func ExampleCapitalize() {
	result := Capitalize("hello, world!")
	fmt.Println(result)
	// Output:
	// Hello, world!
}

func ExampleRepeatString() {
	result := RepeatString("ab", 3)
	fmt.Println(result)
	// Output:
	// ababab
}

func ExampleWords() {
	result := Words("hello, world!")
	fmt.Println(result)
	// Output:
	// [hello world]
}
