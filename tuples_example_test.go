package toolbox

import "fmt"

func ExamplePack2() {
	result := Pack2(7, "beautiful number")
	fmt.Printf("%+v", result)
	// Output:
	// {A:7 B:beautiful number}
}

func ExamplePack3() {
	result := Pack3(7, "beautiful number", true)
	fmt.Printf("%+v", result)
	// Output:
	// {A:7 B:beautiful number C:true}
}

func ExamplePack4() {
	result := Pack4(7, "beautiful number", true, 3.14)
	fmt.Printf("%+v", result)
	// Output:
	// {A:7 B:beautiful number C:true D:3.14}
}

func ExamplePack5() {
	result := Pack5(7, "beautiful number", true, 3.14, "pi")
	fmt.Printf("%+v", result)
	// Output:
	// {A:7 B:beautiful number C:true D:3.14 E:pi}
}

func ExamplePack6() {
	result := Pack6(7, "beautiful number", true, 3.14, "pi", 3.14)
	fmt.Printf("%+v", result)
	// Output:
	// {A:7 B:beautiful number C:true D:3.14 E:pi F:3.14}
}

func ExamplePack7() {
	result := Pack7(7, "beautiful number", true, 3.14, "pi", 3.14, "pi")
	fmt.Printf("%+v", result)
	// Output:
	// {A:7 B:beautiful number C:true D:3.14 E:pi F:3.14 G:pi}
}

func ExamplePack8() {
	result := Pack8(7, "beautiful number", true, 3.14, "pi", 3.14, "pi", 3.14)
	fmt.Printf("%+v", result)
	// Output:
	// {A:7 B:beautiful number C:true D:3.14 E:pi F:3.14 G:pi H:3.14}
}

func ExamplePack9() {
	result := Pack9(7, "beautiful number", true, 3.14, "pi", 3.14, "pi", 3.14, "pi")
	fmt.Printf("%+v", result)
	// Output:
	// {A:7 B:beautiful number C:true D:3.14 E:pi F:3.14 G:pi H:3.14 I:pi}
}

func ExamplePack10() {
	result := Pack10(7, "beautiful number", true, 3.14, "pi", 3.14, "pi", 3.14, "pi", 3.14)
	fmt.Printf("%+v", result)
	// Output:
	// {A:7 B:beautiful number C:true D:3.14 E:pi F:3.14 G:pi H:3.14 I:pi J:3.14}
}

func ExamplePack11() {
	result := Pack11(7, "beautiful number", true, 3.14, "pi", 3.14, "pi", 3.14, "pi", 3.14, "pi")
	fmt.Printf("%+v", result)
	// Output:
	// {A:7 B:beautiful number C:true D:3.14 E:pi F:3.14 G:pi H:3.14 I:pi J:3.14 K:pi}
}

func ExampleUnpack2() {
	result1, result2 := Unpack2(Tuple2[int, string]{A: 7, B: "b"})
	fmt.Println(result1, result2)
	// Output:
	// 7 b
}

func ExampleUnpack3() {
	result1, result2, result3 := Unpack3(Tuple3[int, string, bool]{A: 7, B: "b", C: true})
	fmt.Println(result1, result2, result3)
	// Output:
	// 7 b true
}

func ExampleUnpack4() {
	result1, result2, result3, result4 := Unpack4(Tuple4[int, string, bool, float64]{A: 7, B: "b", C: true, D: 3.14})
	fmt.Println(result1, result2, result3, result4)
	// Output:
	// 7 b true 3.14
}

func ExampleUnpack5() {
	result1, result2, result3, result4, result5 := Unpack5(Tuple5[int, string, bool, float64, string]{A: 7, B: "b", C: true, D: 3.14, E: "e"})
	fmt.Println(result1, result2, result3, result4, result5)
	// Output:
	// 7 b true 3.14 e
}

func ExampleUnpack6() {
	result1, result2, result3, result4, result5, result6 := Unpack6(Tuple6[int, string, bool, float64, string, float64]{A: 7, B: "b", C: true, D: 3.14, E: "e", F: 3.14})
	fmt.Println(result1, result2, result3, result4, result5, result6)
	// Output:
	// 7 b true 3.14 e 3.14
}

func ExampleUnpack7() {
	result1, result2, result3, result4, result5, result6, result7 := Unpack7(Tuple7[int, string, bool, float64, string, float64, string]{A: 7, B: "b", C: true, D: 3.14, E: "e", F: 3.14, G: "g"})
	fmt.Println(result1, result2, result3, result4, result5, result6, result7)
	// Output:
	// 7 b true 3.14 e 3.14 g
}

func ExampleUnpack8() {
	result1, result2, result3, result4, result5, result6, result7, result8 := Unpack8(Tuple8[int, string, bool, float64, string, float64, string, float64]{A: 7, B: "b", C: true, D: 3.14, E: "e", F: 3.14, G: "g", H: 3.14})
	fmt.Println(result1, result2, result3, result4, result5, result6, result7, result8)
	// Output:
	// 7 b true 3.14 e 3.14 g 3.14
}

func ExampleUnpack9() {
	result1, result2, result3, result4, result5, result6, result7, result8, result9 := Unpack9(Tuple9[int, string, bool, float64, string, float64, string, float64, string]{A: 7, B: "b", C: true, D: 3.14, E: "e", F: 3.14, G: "g", H: 3.14, I: "i"})
	fmt.Println(result1, result2, result3, result4, result5, result6, result7, result8, result9)
	// Output:
	// 7 b true 3.14 e 3.14 g 3.14 i
}

func ExampleUnpack10() {
	result1, result2, result3, result4, result5, result6, result7, result8, result9, result10 := Unpack10(Tuple10[int, string, bool, float64, string, float64, string, float64, string, float64]{A: 7, B: "b", C: true, D: 3.14, E: "e", F: 3.14, G: "g", H: 3.14, I: "i", J: 3.14})
	fmt.Println(result1, result2, result3, result4, result5, result6, result7, result8, result9, result10)
	// Output:
	// 7 b true 3.14 e 3.14 g 3.14 i 3.14
}

func ExampleUnpack11() {
	result1, result2, result3, result4, result5, result6, result7, result8, result9, result10, result11 := Unpack11(Tuple11[int, string, bool, float64, string, float64, string, float64, string, float64, string]{A: 7, B: "b", C: true, D: 3.14, E: "e", F: 3.14, G: "g", H: 3.14, I: "i", J: 3.14, K: "k"})
	fmt.Println(result1, result2, result3, result4, result5, result6, result7, result8, result9, result10, result11)
	// Output:
	// 7 b true 3.14 e 3.14 g 3.14 i 3.14 k
}

func ExampleZip2() {
	collection1 := []int{1, 2, 3, 4}
	collection2 := []string{"a", "b", "c"}
	result := Zip2(collection1, collection2)
	fmt.Println(result)
	// Output:
	// [{1 a} {2 b} {3 c} {4 }]
}

func ExampleZip3() {
	collection1 := []int{1, 2, 3, 4}
	collection2 := []string{"a", "b", "c"}
	collection3 := []bool{true, false}
	result := Zip3(collection1, collection2, collection3)
	fmt.Println(result)
	// Output:
	// [{1 a true} {2 b false} {3 c false} {4  false}]
}

func ExampleZip4() {
	collection1 := []int{1, 2, 3, 4}
	collection2 := []string{"a", "b", "c"}
	collection3 := []bool{true, false}
	collection4 := []float64{3.14, 2.72}
	result := Zip4(collection1, collection2, collection3, collection4)
	fmt.Println(result)
	// Output:
	// [{1 a true 3.14} {2 b false 2.72} {3 c false 0} {4  false 0}]
}

func ExampleZip5() {
	collection1 := []int{1, 2, 3, 4}
	collection2 := []string{"a", "b", "c"}
	collection3 := []bool{true, false}
	collection4 := []float64{3.14, 2.72}
	collection5 := []string{"e", "f"}
	result := Zip5(collection1, collection2, collection3, collection4, collection5)
	fmt.Println(result)
	// Output:
	// [{1 a true 3.14 e} {2 b false 2.72 f} {3 c false 0 } {4  false 0 }]
}

func ExampleZip6() {
	collection1 := []int{1, 2, 3, 4}
	collection2 := []string{"a", "b", "c"}
	collection3 := []bool{true, false}
	collection4 := []float64{3.14, 2.72}
	collection5 := []string{"e", "f"}
	collection6 := []float64{3.14}
	result := Zip6(collection1, collection2, collection3, collection4, collection5, collection6)
	fmt.Println(result)
	// Output:
	// [{1 a true 3.14 e 3.14} {2 b false 2.72 f 0} {3 c false 0  0} {4  false 0  0}]
}

func ExampleZip7() {
	collection1 := []int{1, 2, 3, 4}
	collection2 := []string{"a", "b", "c"}
	collection3 := []bool{true, false}
	collection4 := []float64{3.14, 2.72}
	collection5 := []string{"e", "f"}
	collection6 := []float64{3.14}
	collection7 := []string{"g"}
	result := Zip7(collection1, collection2, collection3, collection4, collection5, collection6, collection7)
	fmt.Println(result)
	// Output:
	// [{1 a true 3.14 e 3.14 g} {2 b false 2.72 f 0 } {3 c false 0  0 } {4  false 0  0 }]
}

func ExampleZip8() {
	collection1 := []int{1, 2, 3, 4}
	collection2 := []string{"a", "b", "c"}
	collection3 := []bool{true, false}
	collection4 := []float64{3.14, 2.72}
	collection5 := []string{"e", "f"}
	collection6 := []float64{3.14}
	collection7 := []string{"g"}
	collection8 := []float64{3.14}
	result := Zip8(collection1, collection2, collection3, collection4, collection5, collection6, collection7, collection8)
	fmt.Println(result)
	// Output:
	// [{1 a true 3.14 e 3.14 g 3.14} {2 b false 2.72 f 0  0} {3 c false 0  0  0} {4  false 0  0  0}]
}

func ExampleZip9() {
	collection1 := []int{1, 2, 3, 4}
	collection2 := []string{"a", "b", "c"}
	collection3 := []bool{true, false}
	collection4 := []float64{3.14, 2.72}
	collection5 := []string{"e", "f"}
	collection6 := []float64{3.14}
	collection7 := []string{"g"}
	collection8 := []float64{3.14}
	collection9 := []string{"i"}
	result := Zip9(collection1, collection2, collection3, collection4, collection5, collection6, collection7, collection8, collection9)
	fmt.Println(result)
	// Output:
	// [{1 a true 3.14 e 3.14 g 3.14 i} {2 b false 2.72 f 0  0 } {3 c false 0  0  0 } {4  false 0  0  0 }]
}

func ExampleZip10() {
	collection1 := []int{1, 2, 3, 4}
	collection2 := []string{"a", "b", "c"}
	collection3 := []bool{true, false}
	collection4 := []float64{3.14, 2.72}
	collection5 := []string{"e", "f"}
	collection6 := []float64{3.14}
	collection7 := []string{"g"}
	collection8 := []float64{3.14}
	collection9 := []string{"i"}
	collection10 := []float64{3.14}
	result := Zip10(collection1, collection2, collection3, collection4, collection5, collection6, collection7, collection8, collection9, collection10)
	fmt.Println(result)
	// Output:
	// [{1 a true 3.14 e 3.14 g 3.14 i 3.14} {2 b false 2.72 f 0  0  0} {3 c false 0  0  0  0} {4  false 0  0  0  0}]
}

func ExampleZip11() {
	collection1 := []int{1, 2, 3, 4}
	collection2 := []string{"a", "b", "c"}
	collection3 := []bool{true, false}
	collection4 := []float64{3.14, 2.72}
	collection5 := []string{"e", "f"}
	collection6 := []float64{3.14}
	collection7 := []string{"g"}
	collection8 := []float64{3.14}
	collection9 := []string{"i"}
	collection10 := []float64{3.14}
	collection11 := []string{"j"}
	result := Zip11(collection1, collection2, collection3, collection4, collection5, collection6, collection7, collection8, collection9, collection10, collection11)
	fmt.Println(result)
	// Output:
	// [{1 a true 3.14 e 3.14 g 3.14 i 3.14 j} {2 b false 2.72 f 0  0  0 } {3 c false 0  0  0  0 } {4  false 0  0  0  0 }]
}

func ExampleUnzip2() {
	collection := []Tuple2[int, string]{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}
	result1, result2 := Unzip2(collection)
	fmt.Println(result1, result2)
	// Output:
	// [1 2 3] [a b c]
}

func ExampleUnzip3() {
	collection := []Tuple3[int, string, bool]{
		{1, "a", true},
		{2, "b", false},
		{3, "c", false},
	}
	result1, result2, result3 := Unzip3(collection)
	fmt.Println(result1, result2, result3)
	// Output:
	// [1 2 3] [a b c] [true false false]
}

func ExampleUnzip4() {
	collection := []Tuple4[int, string, bool, float64]{
		{1, "a", true, 3.14},
		{2, "b", false, 2.72},
		{3, "c", false, 0},
	}
	result1, result2, result3, result4 := Unzip4(collection)
	fmt.Println(result1, result2, result3, result4)
	// Output:
	// [1 2 3] [a b c] [true false false] [3.14 2.72 0]
}

func ExampleUnzip5() {
	collection := []Tuple5[int, string, bool, float64, string]{
		{1, "a", true, 3.14, "e"},
		{2, "b", false, 2.72, "f"},
		{3, "c", false, 0, ""},
	}
	result1, result2, result3, result4, result5 := Unzip5(collection)
	fmt.Println(result1, result2, result3, result4, result5)
	// Output:
	// [1 2 3] [a b c] [true false false] [3.14 2.72 0] [e f ]
}

func ExampleUnzip6() {
	collection := []Tuple6[int, string, bool, float64, string, float64]{
		{1, "a", true, 3.14, "e", 3.14},
		{2, "b", false, 2.72, "f", 0},
		{3, "c", false, 0, "", 0},
	}
	result1, result2, result3, result4, result5, result6 := Unzip6(collection)
	fmt.Println(result1, result2, result3, result4, result5, result6)
	// Output:
	// [1 2 3] [a b c] [true false false] [3.14 2.72 0] [e f ] [3.14 0 0]
}

func ExampleUnzip7() {
	collection := []Tuple7[int, string, bool, float64, string, float64, string]{
		{1, "a", true, 3.14, "e", 3.14, "g"},
		{2, "b", false, 2.72, "f", 0, ""},
		{3, "c", false, 0, "", 0, ""},
	}
	result1, result2, result3, result4, result5, result6, result7 := Unzip7(collection)
	fmt.Println(result1, result2, result3, result4, result5, result6, result7)
	// Output:
	// [1 2 3] [a b c] [true false false] [3.14 2.72 0] [e f ] [3.14 0 0] [g  ]
}

func ExampleUnzip8() {
	collection := []Tuple8[int, string, bool, float64, string, float64, string, float64]{
		{1, "a", true, 3.14, "e", 3.14, "g", 3.14},
		{2, "b", false, 2.72, "f", 0, "", 0},
		{3, "c", false, 0, "", 0, "", 0},
	}
	result1, result2, result3, result4, result5, result6, result7, result8 := Unzip8(collection)
	fmt.Println(result1, result2, result3, result4, result5, result6, result7, result8)
	// Output:
	// [1 2 3] [a b c] [true false false] [3.14 2.72 0] [e f ] [3.14 0 0] [g  ] [3.14 0 0]
}

func ExampleUnzip9() {
	collection := []Tuple9[int, string, bool, float64, string, float64, string, float64, string]{
		{1, "a", true, 3.14, "e", 3.14, "g", 3.14, "i"},
		{2, "b", false, 2.72, "f", 0, "", 0, ""},
		{3, "c", false, 0, "", 0, "", 0, ""},
	}
	result1, result2, result3, result4, result5, result6, result7, result8, result9 := Unzip9(collection)
	fmt.Println(result1, result2, result3, result4, result5, result6, result7, result8, result9)
	// Output:
	// [1 2 3] [a b c] [true false false] [3.14 2.72 0] [e f ] [3.14 0 0] [g  ] [3.14 0 0] [i  ]
}

func ExampleUnzip10() {
	collection := []Tuple10[int, string, bool, float64, string, float64, string, float64, string, float64]{
		{1, "a", true, 3.14, "e", 3.14, "g", 3.14, "i", 3.14},
		{2, "b", false, 2.72, "f", 0, "", 0, "", 0},
		{3, "c", false, 0, "", 0, "", 0, "", 0},
	}
	result1, result2, result3, result4, result5, result6, result7, result8, result9, result10 := Unzip10(collection)
	fmt.Println(result1, result2, result3, result4, result5, result6, result7, result8, result9, result10)
	// Output:
	// [1 2 3] [a b c] [true false false] [3.14 2.72 0] [e f ] [3.14 0 0] [g  ] [3.14 0 0] [i  ] [3.14 0 0]
}

func ExampleUnzip11() {
	collection := []Tuple11[int, string, bool, float64, string, float64, string, float64, string, float64, string]{
		{1, "a", true, 3.14, "e", 3.14, "g", 3.14, "i", 3.14, "k"},
		{2, "b", false, 2.72, "f", 0, "", 0, "", 0, ""},
		{3, "c", false, 0, "", 0, "", 0, "", 0, ""},
	}
	result1, result2, result3, result4, result5, result6, result7, result8, result9, result10, result11 := Unzip11(collection)
	fmt.Println(result1, result2, result3, result4, result5, result6, result7, result8, result9, result10, result11)
	// Output:
	// [1 2 3] [a b c] [true false false] [3.14 2.72 0] [e f ] [3.14 0 0] [g  ] [3.14 0 0] [i  ] [3.14 0 0] [k  ]
}
