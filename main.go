package main

import (
	"fmt"
	"github.com/Dukenbayev/onelab_homework4/sort_struct"
	"github.com/Dukenbayev/onelab_homework4/top_words"
)

func main() {
	fmt.Println(top_words.MaxCountString("",0))
	fmt.Println(sort_struct.ParseFile("text.txt"))
}