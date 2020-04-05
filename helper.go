package string_manipulation

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

func caseHelper(input string,isCamel bool, rule ...string) []string {
	if !isCamel {
		re := regexp.MustCompile("([a-z])([A-Z])")
		input = re.ReplaceAllString(input, "$1 $2")
	}
	input = strings.Join(strings.Fields(strings.TrimSpace(input)), " ")
	if len(rule) <= 1 {
		rule = []string{".", " ", "_", " ", "-", " "}
	}
	if len(rule) > 1 && len(rule)%2 != 0 {
		panic(errors.New("odd number rule provided please provide in even count"))
	}
	rule = append(rule, ".", " ", "_", " ", "-", " ")

	replacer := strings.NewReplacer(rule ...)
	input = replacer.Replace(input)
	words := strings.Fields(input)
	return words
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func getInput(i input) (input string) {
	if i.Result != "" {
		input = i.Result
	} else {
		input = i.Input
	}
	return
}

func replaceStr(input, search, replace, types string) string {
	lcInput := strings.ToLower(input)
	lcSearch := strings.ToLower(search)
	if input == "" || !strings.Contains(lcInput, lcSearch) {
		return input
	}
	var start int
	if types == "last" {
		start = strings.LastIndex(lcInput, lcSearch)
	} else {
		start = strings.Index(lcInput, lcSearch)
	}
	end := start + len(search)
	return input[:start] + replace + input[end:]
}

func ucfirst(val string) string {
	for i, v := range val {
		return string(unicode.ToUpper(v)) + val[i+1:]
	}
	return ""
}
