package top_words

import (
	"reflect"
	"testing"
)

func TestMaxCountString(t *testing.T) {
	testTable:=[]struct{
		strings string
		num int
		expected []string
	}{
		{
			strings: "1 2 3 4 5 6",
			num : 1,
			expected: []string{""},
		},
		{
			strings: "The the the cat cat cat The Cat Cat",
			num : 2,
			expected: []string{"the","cat"},
		},
		{
			strings: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			num: 6,
			expected: []string{"Lorem", "Ipsum", "the", "industry" ,"standard","dummy"},
		},
	}

	for _,testCase := range testTable {
		result := MaxCountString(testCase.strings,testCase.num)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result. Except %s,got %s",testCase.expected,result)

		}
	}
}
