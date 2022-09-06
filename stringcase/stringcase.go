package stringcase

import (
	"regexp"
	"strings"
	"unicode"
)

const (
	kReplaceCapital = "$1 $2"
)

var capitalRegexp *regexp.Regexp
var caseReplacer *strings.Replacer

func init() {
	capitalRegexp = regexp.MustCompile("([a-z])([A-Z])")
	caseReplacer = strings.NewReplacer(".", " ", "_", " ", "-", " ")
}

func ToCamel(str string) string {
	wordArray := caseHelper(str)
	for i, word := range wordArray {
		if i == 0 {
			wordArray[i] = lcFirst(word)
			continue
		}
		wordArray[i] = ucfirst(word)
	}
	return strings.Join(wordArray, "")
}

func ToPascal(str string) string {
	wordArray := caseHelper(str)
	for i, word := range wordArray {
		wordArray[i] = ucfirst(word)
	}
	return strings.Join(wordArray, "")
}

func ToSnake(str string) string {
	str = capitalRegexp.ReplaceAllString(str, kReplaceCapital)
	wordArray := caseHelper(str)
	return strings.Join(wordArray, "_")
}

func ToKebab(str string) string {
	str = capitalRegexp.ReplaceAllString(str, kReplaceCapital)
	wordArray := caseHelper(str)
	return strings.Join(wordArray, "-")
}

func caseHelper(input string) []string {
	input = strings.Join(strings.Fields(strings.TrimSpace(input)), " ")

	input = caseReplacer.Replace(input)
	words := strings.Fields(input)
	return words
}

func lcFirst(val string) string {
	for _, v := range val {
		return string(unicode.ToLower(v)) + val[len(string(v)):]
	}
	return ""
}

func ucfirst(val string) string {
	for _, v := range val {
		return string(unicode.ToUpper(v)) + val[len(string(v)):]
	}
	return ""
}
