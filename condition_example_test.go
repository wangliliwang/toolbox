package toolbox

import "fmt"

func ExampleTernary() {
	result1 := Ternary(true, 1, 2)
	result2 := Ternary(false, 1, 2)
	fmt.Printf("%+v %+v", result1, result2)
	// Output:
	// 1 2
}

func ExampleTernaryF() {
	result1 := TernaryF(true, func() int {
		return 1
	}, func() int {
		return 2
	})
	result2 := TernaryF(false, func() int {
		return 1
	}, func() int {
		return 2
	})
	fmt.Printf("%+v %+v", result1, result2)
	// Output:
	// 1 2
}

func ExampleIf() {
	result1 := If(true, 1).
		Else(2)
	result2 := If(false, 1).
		Else(2)
	fmt.Printf("%+v %+v", result1, result2)
	// Output:
	// 1 2
}

func ExampleIfF() {
	result1 := IfF(true, func() int { return 1 }).
		ElseF(func() int { return 2 })
	result2 := IfF(false, func() int { return 1 }).
		ElseF(func() int { return 2 })
	fmt.Printf("%+v %+v", result1, result2)
	// Output:
	// 1 2
}

func ExampleElseIf() {
	result1 := If(false, 1).
		ElseIf(true, 2).
		Else(3)
	result2 := If(false, 1).
		ElseIf(false, 2).
		Else(3)
	fmt.Printf("%+v %+v", result1, result2)
	// Output:
	// 2 3
}

func ExampleElseIfF() {
	result1 := If(false, 1).
		ElseIfF(true, func() int { return 2 }).
		Else(3)
	result2 := If(false, 1).
		ElseIfF(false, func() int { return 2 }).
		Else(3)
	fmt.Printf("%+v %+v", result1, result2)
	// Output:
	// 2 3
}

func ExampleSwitch() {
	result1 := Switch(1).
		Case(false, 3).
		Case(true, 4).
		Default(0)
	result2 := Switch(1).
		Case(false, 3).
		Case(false, 4).
		Default(0)
	result3 := Switch(1).
		CaseF(false, func(predicate int) int { return 3 }).
		CaseF(true, func(predicate int) int { return 4 }).
		DefaultF(func(predicate int) int { return 0 })
	result4 := Switch(1).
		CaseF(false, func(predicate int) int { return 3 }).
		CaseF(false, func(predicate int) int { return 4 }).
		DefaultF(func(predicate int) int { return 0 })
	fmt.Printf("%+v %+v %+v %v", result1, result2, result3, result4)
	// Output:
	// 4 0 4 0
}
