package sort_struct

import (
	"io/ioutil"
	"sort"
	"strings"
)

type Field struct {
	name	string
	fType	string
	size	int
}
type Result struct{
	name string
	fType string
	size int
}
func OptimalSort(f []Field) ([]Field){
	sort.SliceStable(f, func(i, j int) bool {
		return f[i].size < f[j].size
	})
	return	f
}

func GetSizeType(s string) int {

	if strings.HasPrefix(s, "*") {
		return 8
	}
	if strings.HasPrefix(s, "map") || strings.HasPrefix(s, "func") || strings.HasPrefix(s, "chan") {
		return 8
	} else if strings.HasPrefix(s, "[]") {
		return 24
	}

	switch s {
	case "int8", "uint8", "byte", "bool":
		return 1
	case "int16", "uint16":
		return 2
	case "int32", "uint32", "float32":
		return 4
	case "int64", "uint64", "float64", "complex64", "int", "uint", "uintptr":
		return 8
	case "complex128", "string":
		return 16
	}

	return 0
}

func ParseFile(path string) []Field{

	bytes, _ := ioutil.ReadFile(path)
	data := string(bytes)

	// extracting fields
	start := strings.Index(data, "{")
	end := strings.Index(data, "}")

	contentStruct := data[start+1 : end]
	fieldsString := strings.Fields(contentStruct)

	fields := make([]Field, 0, (len(fieldsString)+1)/2)

	for i := 0; i < len(fieldsString)-1; i+=2 {

		fields = append(fields, Field{
			name:	fieldsString[i],
			fType:	fieldsString[i+1],
			size:	GetSizeType(fieldsString[i+1]),
		})
	}

	sort.SliceStable(fields, func(i, j int) bool {
		return fields[i].size < fields[j].size
	})

	return fields
}
