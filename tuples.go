package toolbox

// TODO(@wangli) use code generator to generate these tuples.

type Tuple2[T1 any, T2 any] struct {
	A T1
	B T2
}

type Tuple3[T1 any, T2 any, T3 any] struct {
	A T1
	B T2
	C T3
}

type Tuple4[T1 any, T2 any, T3 any, T4 any] struct {
	A T1
	B T2
	C T3
	D T4
}

type Tuple5[T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	A T1
	B T2
	C T3
	D T4
	E T5
}

type Tuple6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] struct {
	A T1
	B T2
	C T3
	D T4
	E T5
	F T6
}

type Tuple7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any] struct {
	A T1
	B T2
	C T3
	D T4
	E T5
	F T6
	G T7
}

type Tuple8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any] struct {
	A T1
	B T2
	C T3
	D T4
	E T5
	F T6
	G T7
	H T8
}

type Tuple9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any] struct {
	A T1
	B T2
	C T3
	D T4
	E T5
	F T6
	G T7
	H T8
	I T9
}

type Tuple10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any] struct {
	A T1
	B T2
	C T3
	D T4
	E T5
	F T6
	G T7
	H T8
	I T9
	J T10
}

type Tuple11[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any] struct {
	A T1
	B T2
	C T3
	D T4
	E T5
	F T6
	G T7
	H T8
	I T9
	J T10
	K T11
}

// Pack2 returns a Tuple2 instance.
func Pack2[T1 any, T2 any](a T1, b T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{
		A: a,
		B: b,
	}
}

// Pack3 returns a Tuple3 instance.
func Pack3[T1 any, T2 any, T3 any](a T1, b T2, c T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{
		A: a,
		B: b,
		C: c,
	}
}

// Pack4 returns a Tuple4 instance.
func Pack4[T1 any, T2 any, T3 any, T4 any](a T1, b T2, c T3, d T4) Tuple4[T1, T2, T3, T4] {
	return Tuple4[T1, T2, T3, T4]{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

// Pack5 returns a Tuple5 instance.
func Pack5[T1 any, T2 any, T3 any, T4 any, T5 any](a T1, b T2, c T3, d T4, e T5) Tuple5[T1, T2, T3, T4, T5] {
	return Tuple5[T1, T2, T3, T4, T5]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
	}
}

// Pack6 returns a Tuple6 instance.
func Pack6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](a T1, b T2, c T3, d T4, e T5, f T6) Tuple6[T1, T2, T3, T4, T5, T6] {
	return Tuple6[T1, T2, T3, T4, T5, T6]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
	}
}

// Pack7 returns a Tuple7 instance.
func Pack7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](a T1, b T2, c T3, d T4, e T5, f T6, g T7) Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	return Tuple7[T1, T2, T3, T4, T5, T6, T7]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
	}
}

// Pack8 returns a Tuple8 instance.
func Pack8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](a T1, b T2, c T3, d T4, e T5, f T6, g T7, h T8) Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
	}
}

// Pack9 returns a Tuple9 instance.
func Pack9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](a T1, b T2, c T3, d T4, e T5, f T6, g T7, h T8, i T9) Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
		I: i,
	}
}

// Pack10 returns a Tuple10 instance.
func Pack10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](a T1, b T2, c T3, d T4, e T5, f T6, g T7, h T8, i T9, j T10) Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
		I: i,
		J: j,
	}
}

// Pack11 returns a Tuple11 instance.
func Pack11[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any](a T1, b T2, c T3, d T4, e T5, f T6, g T7, h T8, i T9, j T10, k T11) Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11] {
	return Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
		I: i,
		J: j,
		K: k,
	}
}

// Unpack2 returns a Tuple2's inner value.
func Unpack2[T1 any, T2 any](t2 Tuple2[T1, T2]) (T1, T2) {
	return t2.A, t2.B
}

// Unpack3 returns a Tuple3's inner value.
func Unpack3[T1 any, T2 any, T3 any](t3 Tuple3[T1, T2, T3]) (T1, T2, T3) {
	return t3.A, t3.B, t3.C
}

// Unpack4 returns a Tuple4's inner value.
func Unpack4[T1 any, T2 any, T3 any, T4 any](t4 Tuple4[T1, T2, T3, T4]) (T1, T2, T3, T4) {
	return t4.A, t4.B, t4.C, t4.D
}

// Unpack5 returns a Tuple5's inner value.
func Unpack5[T1 any, T2 any, T3 any, T4 any, T5 any](t5 Tuple5[T1, T2, T3, T4, T5]) (T1, T2, T3, T4, T5) {
	return t5.A, t5.B, t5.C, t5.D, t5.E
}

// Unpack6 returns a Tuple6's inner value.
func Unpack6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t6 Tuple6[T1, T2, T3, T4, T5, T6]) (T1, T2, T3, T4, T5, T6) {
	return t6.A, t6.B, t6.C, t6.D, t6.E, t6.F
}

// Unpack7 returns a Tuple7's inner value.
func Unpack7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](t7 Tuple7[T1, T2, T3, T4, T5, T6, T7]) (T1, T2, T3, T4, T5, T6, T7) {
	return t7.A, t7.B, t7.C, t7.D, t7.E, t7.F, t7.G
}

// Unpack8 returns a Tuple8's inner value.
func Unpack8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](t8 Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) (T1, T2, T3, T4, T5, T6, T7, T8) {
	return t8.A, t8.B, t8.C, t8.D, t8.E, t8.F, t8.G, t8.H
}

// Unpack9 returns a Tuple9's inner value.
func Unpack9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](t9 Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return t9.A, t9.B, t9.C, t9.D, t9.E, t9.F, t9.G, t9.H, t9.I
}

// Unpack10 returns a Tuple10's inner value.
func Unpack10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](t10 Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	return t10.A, t10.B, t10.C, t10.D, t10.E, t10.F, t10.G, t10.H, t10.I, t10.J
}

// Unpack11 returns a Tuple11's inner value.
func Unpack11[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any](t11 Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) {
	return t11.A, t11.B, t11.C, t11.D, t11.E, t11.F, t11.G, t11.H, t11.I, t11.J, t11.K
}

// Zip2 returns a Tuple2 slice, whose length is max of collection1 and collection2.
func Zip2[T1 any, T2 any](collection1 []T1, collection2 []T2) []Tuple2[T1, T2] {
	maxLength := Max(len(collection1), len(collection2))
	result := make([]Tuple2[T1, T2], maxLength)
	for i := 0; i < maxLength; i++ {
		item1, _ := Nth(collection1, i)
		item2, _ := Nth(collection2, i)
		result[i] = Tuple2[T1, T2]{A: item1, B: item2}
	}
	return result
}

// Zip3 returns a Tuple3 slice, whose length is max of collection1, collection2 and collection3.
func Zip3[T1 any, T2 any, T3 any](collection1 []T1, collection2 []T2, collection3 []T3) []Tuple3[T1, T2, T3] {
	maxLength := Max(len(collection1), len(collection2), len(collection3))
	result := make([]Tuple3[T1, T2, T3], maxLength)
	for i := 0; i < maxLength; i++ {
		item1, _ := Nth(collection1, i)
		item2, _ := Nth(collection2, i)
		item3, _ := Nth(collection3, i)
		result[i] = Tuple3[T1, T2, T3]{A: item1, B: item2, C: item3}
	}
	return result
}

// Zip4 returns a Tuple4 slice, whose length is max of collection1, collection2, collection3 and collection4.
func Zip4[T1 any, T2 any, T3 any, T4 any](collection1 []T1, collection2 []T2, collection3 []T3, collection4 []T4) []Tuple4[T1, T2, T3, T4] {
	maxLength := Max(len(collection1), len(collection2), len(collection3), len(collection4))
	result := make([]Tuple4[T1, T2, T3, T4], maxLength)
	for i := 0; i < maxLength; i++ {
		item1, _ := Nth(collection1, i)
		item2, _ := Nth(collection2, i)
		item3, _ := Nth(collection3, i)
		item4, _ := Nth(collection4, i)
		result[i] = Tuple4[T1, T2, T3, T4]{A: item1, B: item2, C: item3, D: item4}
	}
	return result
}

// Zip5 returns a Tuple5 slice, whose length is max of collection1, collection2, collection3, collection4 and collection5.
func Zip5[T1 any, T2 any, T3 any, T4 any, T5 any](collection1 []T1, collection2 []T2, collection3 []T3, collection4 []T4, collection5 []T5) []Tuple5[T1, T2, T3, T4, T5] {
	maxLength := Max(len(collection1), len(collection2), len(collection3), len(collection4), len(collection5))
	result := make([]Tuple5[T1, T2, T3, T4, T5], maxLength)
	for i := 0; i < maxLength; i++ {
		item1, _ := Nth(collection1, i)
		item2, _ := Nth(collection2, i)
		item3, _ := Nth(collection3, i)
		item4, _ := Nth(collection4, i)
		item5, _ := Nth(collection5, i)
		result[i] = Tuple5[T1, T2, T3, T4, T5]{A: item1, B: item2, C: item3, D: item4, E: item5}
	}
	return result
}

// Zip6 returns a Tuple6 slice, whose length is max of collection1, collection2, collection3, collection4, collection5 and collection6.
func Zip6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](collection1 []T1, collection2 []T2, collection3 []T3, collection4 []T4, collection5 []T5, collection6 []T6) []Tuple6[T1, T2, T3, T4, T5, T6] {
	maxLength := Max(len(collection1), len(collection2), len(collection3), len(collection4), len(collection5), len(collection6))
	result := make([]Tuple6[T1, T2, T3, T4, T5, T6], maxLength)
	for i := 0; i < maxLength; i++ {
		item1, _ := Nth(collection1, i)
		item2, _ := Nth(collection2, i)
		item3, _ := Nth(collection3, i)
		item4, _ := Nth(collection4, i)
		item5, _ := Nth(collection5, i)
		item6, _ := Nth(collection6, i)
		result[i] = Tuple6[T1, T2, T3, T4, T5, T6]{A: item1, B: item2, C: item3, D: item4, E: item5, F: item6}
	}
	return result
}

// Zip7 returns a Tuple7 slice, whose length is max of collection1, collection2, collection3, collection4, collection5, collection6 and collection7.
func Zip7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](collection1 []T1, collection2 []T2, collection3 []T3, collection4 []T4, collection5 []T5, collection6 []T6, collection7 []T7) []Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	maxLength := Max(len(collection1), len(collection2), len(collection3), len(collection4), len(collection5), len(collection6), len(collection7))
	result := make([]Tuple7[T1, T2, T3, T4, T5, T6, T7], maxLength)
	for i := 0; i < maxLength; i++ {
		item1, _ := Nth(collection1, i)
		item2, _ := Nth(collection2, i)
		item3, _ := Nth(collection3, i)
		item4, _ := Nth(collection4, i)
		item5, _ := Nth(collection5, i)
		item6, _ := Nth(collection6, i)
		item7, _ := Nth(collection7, i)
		result[i] = Tuple7[T1, T2, T3, T4, T5, T6, T7]{A: item1, B: item2, C: item3, D: item4, E: item5, F: item6, G: item7}
	}
	return result
}

// Zip8 returns a Tuple8 slice, whose length is max of collection1, collection2, collection3, collection4, collection5, collection6, collection7 and collection8.
func Zip8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](collection1 []T1, collection2 []T2, collection3 []T3, collection4 []T4, collection5 []T5, collection6 []T6, collection7 []T7, collection8 []T8) []Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	maxLength := Max(len(collection1), len(collection2), len(collection3), len(collection4), len(collection5), len(collection6), len(collection7), len(collection8))
	result := make([]Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], maxLength)
	for i := 0; i < maxLength; i++ {
		item1, _ := Nth(collection1, i)
		item2, _ := Nth(collection2, i)
		item3, _ := Nth(collection3, i)
		item4, _ := Nth(collection4, i)
		item5, _ := Nth(collection5, i)
		item6, _ := Nth(collection6, i)
		item7, _ := Nth(collection7, i)
		item8, _ := Nth(collection8, i)
		result[i] = Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{A: item1, B: item2, C: item3, D: item4, E: item5, F: item6, G: item7, H: item8}
	}
	return result
}

// Zip9 returns a Tuple9 slice, whose length is max of collection1, collection2, collection3, collection4, collection5, collection6, collection7, collection8 and collection9.
func Zip9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](collection1 []T1, collection2 []T2, collection3 []T3, collection4 []T4, collection5 []T5, collection6 []T6, collection7 []T7, collection8 []T8, collection9 []T9) []Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	maxLength := Max(len(collection1), len(collection2), len(collection3), len(collection4), len(collection5), len(collection6), len(collection7), len(collection8), len(collection9))
	result := make([]Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], maxLength)
	for i := 0; i < maxLength; i++ {
		item1, _ := Nth(collection1, i)
		item2, _ := Nth(collection2, i)
		item3, _ := Nth(collection3, i)
		item4, _ := Nth(collection4, i)
		item5, _ := Nth(collection5, i)
		item6, _ := Nth(collection6, i)
		item7, _ := Nth(collection7, i)
		item8, _ := Nth(collection8, i)
		item9, _ := Nth(collection9, i)
		result[i] = Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{A: item1, B: item2, C: item3, D: item4, E: item5, F: item6, G: item7, H: item8, I: item9}
	}
	return result
}

// Zip10 returns a Tuple10 slice, whose length is max of collection1, collection2, collection3, collection4, collection5, collection6, collection7, collection8, collection9 and collection10.
func Zip10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](collection1 []T1, collection2 []T2, collection3 []T3, collection4 []T4, collection5 []T5, collection6 []T6, collection7 []T7, collection8 []T8, collection9 []T9, collection10 []T10) []Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	maxLength := Max(len(collection1), len(collection2), len(collection3), len(collection4), len(collection5), len(collection6), len(collection7), len(collection8), len(collection9), len(collection10))
	result := make([]Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], maxLength)
	for i := 0; i < maxLength; i++ {
		item1, _ := Nth(collection1, i)
		item2, _ := Nth(collection2, i)
		item3, _ := Nth(collection3, i)
		item4, _ := Nth(collection4, i)
		item5, _ := Nth(collection5, i)
		item6, _ := Nth(collection6, i)
		item7, _ := Nth(collection7, i)
		item8, _ := Nth(collection8, i)
		item9, _ := Nth(collection9, i)
		item10, _ := Nth(collection10, i)
		result[i] = Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{A: item1, B: item2, C: item3, D: item4, E: item5, F: item6, G: item7, H: item8, I: item9, J: item10}
	}
	return result
}

// Zip11 returns a Tuple11 slice, whose length is max of collection1, collection2, collection3, collection4, collection5, collection6, collection7, collection8, collection9, collection10 and collection11.
func Zip11[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any](collection1 []T1, collection2 []T2, collection3 []T3, collection4 []T4, collection5 []T5, collection6 []T6, collection7 []T7, collection8 []T8, collection9 []T9, collection10 []T10, collection11 []T11) []Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11] {
	maxLength := Max(len(collection1), len(collection2), len(collection3), len(collection4), len(collection5), len(collection6), len(collection7), len(collection8), len(collection9), len(collection10), len(collection11))
	result := make([]Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11], maxLength)
	for i := 0; i < maxLength; i++ {
		item1, _ := Nth(collection1, i)
		item2, _ := Nth(collection2, i)
		item3, _ := Nth(collection3, i)
		item4, _ := Nth(collection4, i)
		item5, _ := Nth(collection5, i)
		item6, _ := Nth(collection6, i)
		item7, _ := Nth(collection7, i)
		item8, _ := Nth(collection8, i)
		item9, _ := Nth(collection9, i)
		item10, _ := Nth(collection10, i)
		item11, _ := Nth(collection11, i)
		result[i] = Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{A: item1, B: item2, C: item3, D: item4, E: item5, F: item6, G: item7, H: item8, I: item9, J: item10, K: item11}
	}
	return result
}

// Unzip2 returns 2 slices, whose elements come from Tuple2-collection.
func Unzip2[T1 any, T2 any](collection []Tuple2[T1, T2]) ([]T1, []T2) {
	result1 := make([]T1, len(collection))
	result2 := make([]T2, len(collection))
	for i := 0; i < len(collection); i++ {
		result1[i] = collection[i].A
		result2[i] = collection[i].B
	}
	return result1, result2
}

// Unzip3 returns 3 slices, whose elements come from Tuple3-collection.
func Unzip3[T1 any, T2 any, T3 any](collection []Tuple3[T1, T2, T3]) ([]T1, []T2, []T3) {
	result1 := make([]T1, len(collection))
	result2 := make([]T2, len(collection))
	result3 := make([]T3, len(collection))
	for i := 0; i < len(collection); i++ {
		result1[i] = collection[i].A
		result2[i] = collection[i].B
		result3[i] = collection[i].C
	}
	return result1, result2, result3
}

// Unzip4 returns 4 slices, whose elements come from Tuple4-collection.
func Unzip4[T1 any, T2 any, T3 any, T4 any](collection []Tuple4[T1, T2, T3, T4]) ([]T1, []T2, []T3, []T4) {
	result1 := make([]T1, len(collection))
	result2 := make([]T2, len(collection))
	result3 := make([]T3, len(collection))
	result4 := make([]T4, len(collection))
	for i := 0; i < len(collection); i++ {
		result1[i] = collection[i].A
		result2[i] = collection[i].B
		result3[i] = collection[i].C
		result4[i] = collection[i].D
	}
	return result1, result2, result3, result4
}

// Unzip5 returns 5 slices, whose elements come from Tuple5-collection.
func Unzip5[T1 any, T2 any, T3 any, T4 any, T5 any](collection []Tuple5[T1, T2, T3, T4, T5]) ([]T1, []T2, []T3, []T4, []T5) {
	result1 := make([]T1, len(collection))
	result2 := make([]T2, len(collection))
	result3 := make([]T3, len(collection))
	result4 := make([]T4, len(collection))
	result5 := make([]T5, len(collection))
	for i := 0; i < len(collection); i++ {
		result1[i] = collection[i].A
		result2[i] = collection[i].B
		result3[i] = collection[i].C
		result4[i] = collection[i].D
		result5[i] = collection[i].E
	}
	return result1, result2, result3, result4, result5
}

// Unzip6 returns 6 slices, whose elements come from Tuple6-collection.
func Unzip6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](collection []Tuple6[T1, T2, T3, T4, T5, T6]) ([]T1, []T2, []T3, []T4, []T5, []T6) {
	result1 := make([]T1, len(collection))
	result2 := make([]T2, len(collection))
	result3 := make([]T3, len(collection))
	result4 := make([]T4, len(collection))
	result5 := make([]T5, len(collection))
	result6 := make([]T6, len(collection))
	for i := 0; i < len(collection); i++ {
		result1[i] = collection[i].A
		result2[i] = collection[i].B
		result3[i] = collection[i].C
		result4[i] = collection[i].D
		result5[i] = collection[i].E
		result6[i] = collection[i].F
	}
	return result1, result2, result3, result4, result5, result6
}

// Unzip7 returns 7 slices, whose elements come from Tuple7-collection.
func Unzip7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](collection []Tuple7[T1, T2, T3, T4, T5, T6, T7]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7) {
	result1 := make([]T1, len(collection))
	result2 := make([]T2, len(collection))
	result3 := make([]T3, len(collection))
	result4 := make([]T4, len(collection))
	result5 := make([]T5, len(collection))
	result6 := make([]T6, len(collection))
	result7 := make([]T7, len(collection))
	for i := 0; i < len(collection); i++ {
		result1[i] = collection[i].A
		result2[i] = collection[i].B
		result3[i] = collection[i].C
		result4[i] = collection[i].D
		result5[i] = collection[i].E
		result6[i] = collection[i].F
		result7[i] = collection[i].G
	}
	return result1, result2, result3, result4, result5, result6, result7
}

// Unzip8 returns 8 slices, whose elements come from Tuple8-collection.
func Unzip8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](collection []Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8) {
	result1 := make([]T1, len(collection))
	result2 := make([]T2, len(collection))
	result3 := make([]T3, len(collection))
	result4 := make([]T4, len(collection))
	result5 := make([]T5, len(collection))
	result6 := make([]T6, len(collection))
	result7 := make([]T7, len(collection))
	result8 := make([]T8, len(collection))
	for i := 0; i < len(collection); i++ {
		result1[i] = collection[i].A
		result2[i] = collection[i].B
		result3[i] = collection[i].C
		result4[i] = collection[i].D
		result5[i] = collection[i].E
		result6[i] = collection[i].F
		result7[i] = collection[i].G
		result8[i] = collection[i].H
	}
	return result1, result2, result3, result4, result5, result6, result7, result8
}

// Unzip9 returns 9 slices, whose elements come from Tuple9-collection.
func Unzip9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](collection []Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9) {
	result1 := make([]T1, len(collection))
	result2 := make([]T2, len(collection))
	result3 := make([]T3, len(collection))
	result4 := make([]T4, len(collection))
	result5 := make([]T5, len(collection))
	result6 := make([]T6, len(collection))
	result7 := make([]T7, len(collection))
	result8 := make([]T8, len(collection))
	result9 := make([]T9, len(collection))
	for i := 0; i < len(collection); i++ {
		result1[i] = collection[i].A
		result2[i] = collection[i].B
		result3[i] = collection[i].C
		result4[i] = collection[i].D
		result5[i] = collection[i].E
		result6[i] = collection[i].F
		result7[i] = collection[i].G
		result8[i] = collection[i].H
		result9[i] = collection[i].I
	}
	return result1, result2, result3, result4, result5, result6, result7, result8, result9
}

// Unzip10 returns 10 slices, whose elements come from Tuple10-collection.
func Unzip10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](collection []Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9, []T10) {
	result1 := make([]T1, len(collection))
	result2 := make([]T2, len(collection))
	result3 := make([]T3, len(collection))
	result4 := make([]T4, len(collection))
	result5 := make([]T5, len(collection))
	result6 := make([]T6, len(collection))
	result7 := make([]T7, len(collection))
	result8 := make([]T8, len(collection))
	result9 := make([]T9, len(collection))
	result10 := make([]T10, len(collection))
	for i := 0; i < len(collection); i++ {
		result1[i] = collection[i].A
		result2[i] = collection[i].B
		result3[i] = collection[i].C
		result4[i] = collection[i].D
		result5[i] = collection[i].E
		result6[i] = collection[i].F
		result7[i] = collection[i].G
		result8[i] = collection[i].H
		result9[i] = collection[i].I
		result10[i] = collection[i].J
	}
	return result1, result2, result3, result4, result5, result6, result7, result8, result9, result10
}

// Unzip11 returns 11 slices, whose elements come from Tuple11-collection.
func Unzip11[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any](collection []Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) ([]T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8, []T9, []T10, []T11) {
	result1 := make([]T1, len(collection))
	result2 := make([]T2, len(collection))
	result3 := make([]T3, len(collection))
	result4 := make([]T4, len(collection))
	result5 := make([]T5, len(collection))
	result6 := make([]T6, len(collection))
	result7 := make([]T7, len(collection))
	result8 := make([]T8, len(collection))
	result9 := make([]T9, len(collection))
	result10 := make([]T10, len(collection))
	result11 := make([]T11, len(collection))
	for i := 0; i < len(collection); i++ {
		result1[i] = collection[i].A
		result2[i] = collection[i].B
		result3[i] = collection[i].C
		result4[i] = collection[i].D
		result5[i] = collection[i].E
		result6[i] = collection[i].F
		result7[i] = collection[i].G
		result8[i] = collection[i].H
		result9[i] = collection[i].I
		result10[i] = collection[i].J
		result11[i] = collection[i].K
	}
	return result1, result2, result3, result4, result5, result6, result7, result8, result9, result10, result11
}
