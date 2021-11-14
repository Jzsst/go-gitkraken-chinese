package cst

import "container/list"

const defaultStr string = "{\n    \"languageOption\": {\n      \"label\": \"English (US)\",\n      \"value\": \"en-us\"\n    },\n    \"menuStrings\": {\n\n    },\n    \"strings\": {\n\n    }\n}"
const length = 1

var leftSide = list.New()

var rightSide = list.New()

func GetDefaultStr() string {
	return defaultStr
}

func GetLeftSide() *list.List {
	leftSide.
	return leftSide
}

func GetRightSide() *list.List {
	return rightSide
}
