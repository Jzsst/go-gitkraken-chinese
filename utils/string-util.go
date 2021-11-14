package utils

import (
	"fmt"
	"sort"
	"strings"
)

const defaultStr string = "{\n    \"languageOption\": {\n      \"label\": \"English (US)\",\n      \"value\": \"en-us\"\n    },\n    \"menuStrings\": {\n\n    },\n    \"strings\": {\n\n    }\n}"
const leftSide string = "<%= "
const rightSide string = " %>"

func StringIndexNum(text string, child string) []int {
	index := strings.Index(text, child)
	if index == -1 {
		return []int{}
	} else {
		var indexes []int
		indexes = append(indexes, index)
		indexes = append(indexes, StringIndexNum(text[strings.Index(text, child)+len(child):], child)...)
		return indexes
	}
}
func StringTranslationCutting(text string) []string {
	fmt.Println(text)
	leftIndex := StringIndexNum(text, leftSide)
	rightIndex := StringIndexNum(text, rightSide)
	sort.Ints(leftIndex)
	sort.Ints(rightIndex)
	var strs []string
	for i := 0; i < len(leftIndex); i++ {
		//left
		strs = append(strs, text[:strings.Index(text, leftSide)])
		//noTran
		strs = append(strs, text[strings.Index(text, leftSide):strings.Index(text, rightSide)+len(rightSide)])
		//right
		if i+1 > len(leftIndex) {
			strs = append(strs, text[rightIndex[i]+len(rightSide):])
		} else {
			strs = append(strs, text[rightIndex[i]+len(rightSide):leftIndex[i+1]])
		}
	}
	fmt.Println(leftIndex, rightSide)
	return []string{}
}
