package main

import (
	"flag"
	"fmt"
	"github.com/wangliliwang/toolbox"
	"os"
	"strings"
	"text/template"
)

var (
	tupleStartIndex  = 2
	tupleEndIndex    = 17
	tuplesTmplString = `// Code generated by go run cmd/generate_tuples.go -output tuples.go
// DO NOT EDIT

package toolbox

{{ range $index, $item := . }}
type {{ $item.TupleStructName }}[{{ $item.GenericDeclaration }}] struct {
	{{ $item.TupleStructBody }}
}

// {{ $item.PackFuncName }} returns a {{ $item.TupleStructName }} instance.
func {{ $item.PackFuncName }}[{{ $item.GenericDeclaration }}]({{ $item.PackFuncParam }}) {{ $item.TupleStructName }}[{{ $item.GenericParam }}] {
	return {{ $item.TupleStructName }}[{{ $item.GenericParam }}]{ {{ $item.TupleInstanceBody }} }
}

// {{ $item.UnpackFuncName }} returns a {{ $item.TupleStructName }}'s inner value.
func {{ $item.UnpackFuncName }}[{{ $item.GenericDeclaration }}]({{ $item.UnpackFuncParam }}) ({{ $item.GenericParam }}) {
	return {{ $item.UnpackFuncReturnValue }}
}

// {{ $item.ZipFuncName }} returns a {{ $item.TupleStructName }} slice, whose length is max of input collections.
func {{ $item.ZipFuncName }}[{{ $item.GenericDeclaration }}]({{ $item.ZipFuncParam }}) []{{ $item.TupleStructName }}[{{ $item.GenericParam }}] {
	maxLength := Max({{ $item.ZipFuncMaxFuncParam }})
	result := make([]{{ $item.TupleStructName }}[{{ $item.GenericParam }}], maxLength)
	for index := 0; index < maxLength; index++ {
		{{ $item.ZipFuncBodyNth }}
		result[index] = {{ $item.TupleStructName }}[{{ $item.GenericParam }}]{ {{ $item.TupleInstanceBody }} }
	}
	return result
}

// {{ $item.UnzipFuncName }} returns {{ $item.Index }} slices, whose elements come from {{ $item.TupleStructName }}-collection.
func {{ $item.UnzipFuncName }}[{{ $item.GenericDeclaration }}](collection []{{ $item.TupleStructName }}[{{ $item.GenericParam }}]) ({{ $item.UnzipReturnParam }}) {
	length := len(collection)
	{{ $item.UnzipResultSlice }}
	for index := 0; index < length; index++ {
		{{ $item.UnzipLoopBody }}
	}
	return {{ $item.UnzipReturnValue }}
}
{{ end }}
`
)

type Item struct {
	Index              int    // 2
	GenericDeclaration string // T1 any, T2 any
	GenericParam       string // T1, T2

	TupleStructName   string // Tuple2
	TupleStructBody   string // A T1\n	B T2
	TupleInstanceBody string // A: a, B: b

	PackFuncName  string // Pack2
	PackFuncParam string // a T1, b T2

	UnpackFuncName        string // Unpack2
	UnpackFuncParam       string // t2 Tuple2[T1, T2]
	UnpackFuncReturnValue string // t2.A, t2.B

	ZipFuncName         string // Zip2
	ZipFuncParam        string // collection1 []T1, collection2 []T2
	ZipFuncMaxFuncParam string // len(collection1), len(collection2)
	ZipFuncBodyNth      string // a, _ := Nth(collection1, i)\n		b, _ := Nth(collection2, i)

	UnzipFuncName    string // Unzip2
	UnzipResultSlice string // result1 := make([]T1, len(collection))\n	result2 := make([]T2, len(collection))
	UnzipReturnParam string // []T1, []T2
	UnzipLoopBody    string // result1[i] = collection[i].A\n		result2[i] = collection[i].B
	UnzipReturnValue string // result1, result2
}

func main() {
	// parse input
	output := flag.String("output", "tuples.go", "file name to write")
	if output == nil {
		panic("nil outputFilename")
	}
	outputFilename := strings.TrimSpace(*output)
	if outputFilename == "" {
		panic("empty outputFilename")
	}

	// generate
	lowerLetters := toolbox.Map(toolbox.RangeFrom('a', 26), func(item int32, index int) string {
		return string(byte(item))
	})
	upperLetters := toolbox.Map(toolbox.RangeFrom('A', 26), func(item int32, index int) string {
		return string(byte(item))
	})
	items := make([]Item, 0)
	for i := tupleStartIndex; i <= tupleEndIndex; i++ {
		genericDeclarations := make([]string, 0)
		genericParams := make([]string, 0)
		tupleStructBodies := make([]string, 0)
		tupleInstanceBodies := make([]string, 0)
		packFuncParams := make([]string, 0)
		unpackFuncReturnValues := make([]string, 0)
		zipFuncParams := make([]string, 0)
		zipFuncMaxFuncParams := make([]string, 0)
		zipFuncBodyNths := make([]string, 0)
		unzipResultSlices := make([]string, 0)
		unzipReturnParams := make([]string, 0)
		unzipLoopBodies := make([]string, 0)
		unzipReturnValues := make([]string, 0)
		for j := 1; j <= i; j++ {
			genericDeclarations = append(genericDeclarations, fmt.Sprintf("T%d any", j))
			genericParams = append(genericParams, fmt.Sprintf("T%d", j))
			tupleStructBodies = append(tupleStructBodies, fmt.Sprintf("%s T%d", upperLetters[j-1], j))
			tupleInstanceBodies = append(tupleInstanceBodies, fmt.Sprintf("%s: %s", upperLetters[j-1], lowerLetters[j-1]))
			packFuncParams = append(packFuncParams, fmt.Sprintf("%s T%d", lowerLetters[j-1], j))
			unpackFuncReturnValues = append(unpackFuncReturnValues, fmt.Sprintf("t%d.%s", i, upperLetters[j-1]))
			zipFuncParams = append(zipFuncParams, fmt.Sprintf("collection%d []T%d", j, j))
			zipFuncMaxFuncParams = append(zipFuncMaxFuncParams, fmt.Sprintf("len(collection%d)", j))
			zipFuncBodyNths = append(zipFuncBodyNths, fmt.Sprintf("%s, _ := Nth(collection%d, index)", lowerLetters[j-1], j))
			unzipResultSlices = append(unzipResultSlices, fmt.Sprintf("result%d := make([]T%d, length)", j, j))
			unzipReturnParams = append(unzipReturnParams, fmt.Sprintf("[]T%d", j))
			unzipLoopBodies = append(unzipLoopBodies, fmt.Sprintf("result%d[index] = collection[index].%s", j, upperLetters[j-1]))
			unzipReturnValues = append(unzipReturnValues, fmt.Sprintf("result%d", j))
		}

		items = append(items, Item{
			Index:              i,
			GenericDeclaration: strings.Join(genericDeclarations, ", "), // T1 any, T2 any
			GenericParam:       strings.Join(genericParams, ", "),       // T1, T2

			TupleStructName:   fmt.Sprintf("Tuple%d", i),               // Tuple2
			TupleStructBody:   strings.Join(tupleStructBodies, "\n\t"), // A T1\n	B T2
			TupleInstanceBody: strings.Join(tupleInstanceBodies, ", "), // A: a, B: b

			PackFuncName:  fmt.Sprintf("Pack%d", i),           // Pack2
			PackFuncParam: strings.Join(packFuncParams, ", "), // a T1, b T2

			UnpackFuncName:        fmt.Sprintf("Unpack%d", i),                                              // Unpack2
			UnpackFuncParam:       fmt.Sprintf("t%d Tuple%d[%s]", i, i, strings.Join(genericParams, ", ")), // t2 Tuple2[T1, T2]
			UnpackFuncReturnValue: strings.Join(unpackFuncReturnValues, ", "),                              // t2.A, t2.B

			ZipFuncName:         fmt.Sprintf("Zip%d", i),                  // Zip2
			ZipFuncParam:        strings.Join(zipFuncParams, ", "),        // collection1 []T1, collection2 []T2
			ZipFuncMaxFuncParam: strings.Join(zipFuncMaxFuncParams, ", "), // len(collection1), len(collection2)
			ZipFuncBodyNth:      strings.Join(zipFuncBodyNths, "\n\t\t"),  // a, _ := Nth(collection1, i)\n		b, _ := Nth(collection2, i)

			UnzipFuncName:    fmt.Sprintf("Unzip%d", i),               // Unzip2
			UnzipResultSlice: strings.Join(unzipResultSlices, "\n\t"), // result1 := make([]T1, len(collection))\n	result2 := make([]T2, len(collection))
			UnzipReturnParam: strings.Join(unzipReturnParams, ", "),   // []T1, []T2
			UnzipLoopBody:    strings.Join(unzipLoopBodies, "\n\t\t"), // result1[i] = collection[i].A\n		result2[i] = collection[i].B
			UnzipReturnValue: strings.Join(unzipReturnValues, ", "),   // result1, result2
		})
	}
	tmpl, parseErr := template.New("tuples").Parse(tuplesTmplString)
	toolbox.PanicIf(parseErr)

	f, openErr := os.OpenFile(outputFilename, os.O_CREATE|os.O_WRONLY, 0644)
	toolbox.PanicIf(openErr)

	executeErr := tmpl.Execute(f, items)
	toolbox.PanicIf(executeErr)
}