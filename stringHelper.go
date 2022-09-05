package main

import (
	"regexp"
	"strings"
	"unicode"
)

var replace = []string{".", " ", "_", " ", "-", " "}

const (
	SelectCapital  = "([a-z])([A-Z])"
	ReplaceCapital = "$1 $2"
)

func caseHelper(input string, isCamel bool) []string {
	if !isCamel {
		re := regexp.MustCompile(SelectCapital)
		input = re.ReplaceAllString(input, ReplaceCapital)
	}
	input = strings.Join(strings.Fields(strings.TrimSpace(input)), " ")

	replacer := strings.NewReplacer(replace...)
	input = replacer.Replace(input)
	words := strings.Fields(input)
	return words
}

func ToPascal(str string) string {
	// removing excess space
	wordArray := caseHelper(str, true)
	for i, word := range wordArray {
		wordArray[i] = ucfirst(word)
	}
	return strings.Join(wordArray, "")
}

func ToCapital(input string) string {
	re := regexp.MustCompile(SelectCapital)
	return re.ReplaceAllString(input, ReplaceCapital)
}

func ToSnake(str string) string {
	// removing excess space
	wordArray := caseHelper(str, false)
	return strings.Join(wordArray, "_")
}

func ucfirst(val string) string {
	for _, v := range val {
		return string(unicode.ToUpper(v)) + val[len(string(v)):]
	}
	return ""
}
